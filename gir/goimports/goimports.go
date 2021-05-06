package goimports

import (
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/pkg/errors"
)

var (
	path     []string
	pathOnce sync.Once
)

func lookPath() error {
	pathOnce.Do(func() {
		goimports, err := exec.LookPath("goimports")
		if err == nil {
			path = []string{goimports}
			return
		}

		goExec, err := exec.LookPath("go")
		if err == nil {
			path = []string{goExec, "run", "golang.org/x/tools/cmd/goimports"}
			return
		}
	})

	if path == nil {
		return errors.New("missing goimports")
	}

	return nil
}

// Pipe pipes the given reader through goimports and into the writer.
func Pipe(w io.Writer, r io.Reader, args ...string) error {
	return run(w, r, args...)
}

// Dir runs goimports recursively on a directory.
func Dir(dir string, args ...string) error {
	return File(dir, args...)
}

// File runs goimports on the given file.
func File(file string, args ...string) error {
	argv := []string{"-w", file}
	argv = append(argv, args...)

	return run(nil, nil, argv...)
}

// run runs goimports.
func run(dst io.Writer, r io.Reader, args ...string) error {
	if err := lookPath(); err != nil {
		return err
	}

	argv := make([]string, 0, len(path)+len(args))
	argv = append(argv, path[1:]...)
	argv = append(argv, args...)

	cmd := exec.Command(path[0], argv...)
	cmd.Stdout = dst
	cmd.Stdin = r

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) || exitErr.Stderr == nil {
			return errors.Wrap(err, "goimports failed")
		}

		return fmt.Errorf(
			"goimports failed with status %d:\n%s",
			exitErr.ExitCode(), exitErr.Stderr,
		)
	}

	return nil
}
