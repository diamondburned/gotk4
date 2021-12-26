// Package pkgconfig provides a wrapper around the pkg-config binary.
package pkgconfig

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var (
	path     string
	pathErr  error
	pathOnce sync.Once
)

func ensure() error {
	pathOnce.Do(func() {
		path, pathErr = exec.LookPath("pkg-config")
		pathErr = fmt.Errorf("failed to find pkg-config: %w", pathErr)
	})
	return pathErr
}

// IncludeDir finds the include directories for the given packages.
func IncludeDirs(pkgs ...string) ([]string, error) {
	if err := ensure(); err != nil {
		return nil, err
	}

	args := []string{"--variable=includedir"}
	args = append(args, pkgs...)

	out, err := exec.Command(path, args...).Output()
	if err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			return nil, fmt.Errorf("pkg-config failed: %w", err)
		}

		return nil, fmt.Errorf(
			"pkg-config failed with status %d:\n%s",
			exitErr.ExitCode(), exitErr.Stderr,
		)
	}

	return strings.Fields(string(out)), nil
}

// FindGIRFiles finds gir files from the given list of pkgs.
func FindGIRFiles(pkgs ...string) ([]string, error) {
	includeDirs, err := IncludeDirs(pkgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to find include dirs: %w", err)
	}

	var girFiles []string

	for _, includeDir := range includeDirs {
		baseDir, name := filepath.Split(includeDir)
		if name != "include" {
			return nil, fmt.Errorf("includeDir has unexpected name %q not 'include'", name)
		}

		err := fs.WalkDir(os.DirFS(baseDir), ".",
			func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if filepath.Ext(d.Name()) == ".gir" {
					girFiles = append(girFiles, filepath.Join(baseDir, path))
				}

				return nil
			},
		)

		if err != nil {
			return nil, fmt.Errorf("failed to walk %q: %w", baseDir, err)
		}
	}

	return girFiles, nil
}
