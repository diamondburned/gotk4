// Copyright (C) 2021, 2023  diamondburned
// Copyright (C) 2023-2024  Luke T. Shumaker

// Package pkgconfig provides a wrapper around the pkg-config binary.
package pkgconfig

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// A ValueSource is a function that takes a list of package names and
// returns a map of package-name to some value.
type ValueSource func(pkgs ...string) (map[string]string, error)

// VarValues returns a map of package-name to variable-value of the
// given variable name for each requested package.  It is not an error
// for a variable to be unset or empty; ret[pkgname] is an empty
// string if that package does not have the variable set.
func VarValues(varname string, pkgs ...string) (map[string]string, error) {
	if len(pkgs) == 0 {
		return nil, nil
	}

	cmdline := append([]string{"pkg-config",
		// Don't be opaque when we fail.
		"--print-errors",
		// On Nix, there are "normal" and "dev" split
		// packages, with different prefixes.  The "prefix"
		// variable is set to the normal package's prefix,
		// while the .pc, .h and .gir files are installed
		// relative to the dev package's prefix.  So tell
		// pkg-config to override the prefix variable with the
		// location of the .pc file.  An alternative way of
		// getting the dev package's suffix would be to
		// inspect the "includedirs" variable and trim off the
		// "/include(/.*)?$" suffix.
		"--define-prefix",
		// main
		"--variable=" + varname, "--"}, pkgs...)
	var stdout strings.Builder
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			return nil, fmt.Errorf("pkg-config failed: %w", err)
		}
		return nil, fmt.Errorf("pkg-config failed with status %d:\n%s",
			exitErr.ExitCode(), exitErr.Stderr)
	}
	return parseValues(pkgs, cmdline, stdout.String())
}

// parseValues parses the output of `pkg-config`.  It is a separate
// function from [Values] for unit-testing purposes.
func parseValues(pkgs []string, cmdline []string, stdout string) (map[string]string, error) {
	ret := make(map[string]string, len(pkgs))
	stdoutStr := strings.TrimRight(stdout, "\n")
	if stdoutStr == "" {
		for i := range pkgs {
			ret[pkgs[i]] = ""
		}
	} else {
		vals := strings.Split(stdoutStr, " ")
		if len(vals) < len(pkgs) {
			// FDO `pkg-config` omits the separating
			// spaces for any leading empty values.  This
			// is likely a bug where the author assumed
			// that `if (str->len > 0)` would be true
			// after the first iteration, but it isn't
			// when the first iteration's `var->len==0`.
			//
			// https://gitlab.freedesktop.org/pkg-config/pkg-config/-/blob/pkg-config-0.29.2/pkg.c?ref_type=tags#L1061-L1062
			partialVals := vals
			vals = make([]string, len(pkgs))
			copy(vals[len(vals)-len(partialVals):], partialVals)
		}
		if len(vals) > len(pkgs) {
			return nil, fmt.Errorf("%v returned %d values, but only expected %d",
				cmdline, len(vals), len(pkgs))
		}

		for i := range pkgs {
			ret[pkgs[i]] = vals[i]
		}
	}

	return ret, nil
}

// VarValuesOrElse is like [VarValues], but if a package has an
// empty/unset value, then that empty value is replaced with the value
// that is returned from the elseFn function.
func VarValuesOrElse(varname string, elseFn ValueSource, pkgs ...string) (map[string]string, error) {
	ret, err := VarValues(varname, pkgs...)
	if err != nil {
		return nil, err
	}

	var badPkgs []string
	for _, pkg := range pkgs {
		if ret[pkg] == "" {
			badPkgs = append(badPkgs, pkg)
		}
	}
	if len(badPkgs) > 0 {
		aug, err := elseFn(badPkgs...)
		if err != nil {
			return nil, err
		}
		for pkg, val := range aug {
			ret[pkg] = val
		}
	}

	return ret, nil
}

// AddPathSuffix takes a [ValueSource] that returns a map of
// package-name to directory, and wraps it so that each `dir` is set
// to `filepath.Join(dir, ...suffix)`.
func AddPathSuffix(inner ValueSource, suffix ...string) ValueSource {
	return func(pkgs ...string) (map[string]string, error) {
		ret, err := inner(pkgs...)
		if err != nil {
			return nil, err
		}
		for pkg, dir := range ret {
			ret[pkg] = filepath.Join(append([]string{dir}, suffix...)...)
		}
		return ret, nil
	}
}

// Prefixes returns a map of package-name to the install prefix for
// each requested package, or an error if this cannot be determined
// for any of the packages.  Common values are "/", "/usr", or
// "/usr/local".
//
// Prefixes is a [ValueSource].
func Prefixes(pkgs ...string) (map[string]string, error) {
	return VarValuesOrElse("prefix", func(pkgs ...string) (map[string]string, error) {
		return nil, fmt.Errorf("could not resolve install prefix for packages: %v", pkgs)
	}, pkgs...)
}

// DataRootDirs returns a map of package-name to the directory for
// read-only architecture-independent data files for each requested
// package, or an error if this cannot be determined for any of the
// packages.  The usual value is "${prefix}/share", i.e. "/usr/share"
// or "/usr/local/share".
//
// DataRootDirs is a [ValueSource].
func DataRootDirs(pkgs ...string) (map[string]string, error) {
	return VarValuesOrElse("datarootdir", AddPathSuffix(Prefixes, "share"), pkgs...)
}

// DataDirs returns a map of package-name to the base directory for
// package-idiosyncratic read-only architecture-independent data files
// for each requested package, or an error if this cannot be
// determined for any of the packages.  The usual value is
// "${datarootdir}", i.e. "/usr/share" or "/usr/local/share"; this is
// *not* a per-package directory, packages usually install their data
// to "${datadir}/${package_name}/".
//
// DataDirs is a [ValueSource].
func DataDirs(pkgs ...string) (map[string]string, error) {
	return VarValuesOrElse("datadir", DataRootDirs, pkgs...)
}

// GIRDirs returns a map of package-name to the directory for GObject
// Introspection Repository XML files for each requested package, or
// an error if this cannot be determined for any of the packages.  The
// usual value is "${datadir}/gir-1.0", i.e. "/usr/share/gir-1.0" or
// "/usr/local/share/gir-1.0".
//
// GIRDirs is a [ValueSource].
func GIRDirs(pkgs ...string) (map[string]string, error) {
	return VarValuesOrElse("girdir", AddPathSuffix(DataDirs, "gir-1.0"), pkgs...)
}
