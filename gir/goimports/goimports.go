package goimports

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
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

// File runs goimports on a temporary file while reading from the given
// filename, then overrides the file with the output.
func File(file string, args ...string) error {
	// Make a new temp file in the same directory as the output to allow mv.
	tmp, err := os.CreateTemp(filepath.Dir(file), ".gir-goimports-")
	if err != nil {
		return errors.Wrap(err, "failed to mktemp")
	}

	defer os.Remove(tmp.Name())
	defer tmp.Close()

	args = append([]string{file}, args...)

	if err := run(tmp, nil, args...); err != nil {
		return err
	}

	if err := tmp.Close(); err != nil {
		return errors.Wrap(err, "failed to close tmp")
	}

	if err := os.Rename(tmp.Name(), file); err != nil {
		return errors.Wrap(err, "failed to swap output")
	}

	return nil
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
		if !errors.As(err, &exitErr) {
			return errors.Wrap(err, "goimports failed")
		}

		return fmt.Errorf(
			"goimports failed with status %d:\n%s",
			exitErr.ExitCode(), exitErr.Stderr,
		)
	}

	return nil
}
