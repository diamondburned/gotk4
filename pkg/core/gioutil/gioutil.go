// Package gioutil provides wrappers around certain GIO classes to be more Go
// idiomatic.
package gioutil

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// NewInputReader wraps around an io.Reader to form a gio.InputStreamer. There
// currently isn't a good way for errors from the reader to make it through to
// the InputStreamer, but it will be closed if an error occurs.
//
// The user should not assume the implementation of gio.InputStreamer, as it may
// be changed in the future.
func NewInputReader(r io.Reader) (gio.InputStreamer, error) {
	rp, wp, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	// Allow GIO to close the pipe using the syscall.
	istream := gio.NewUnixInputStream(int(rp.Fd()), true)

	go func() {
		io.Copy(wp, r)
		wp.Close()
		rp.Close()
	}()

	return istream, nil
}

// StreamReader wraps around a gio.InputStreamer.
type StreamReader struct {
	s   *gio.InputStream
	ctx context.Context
}

// Reader wraps a gio.InputStreamer to provide an io.ReadCloser. The given
// context allows the caller to cancel all ongoing operations done on the new
// ReadCloser.
func Reader(ctx context.Context, s gio.InputStreamer) *StreamReader {
	return &StreamReader{s.BaseInputStream(), ctx}
}

// Read implements io.Reader.
func (r *StreamReader) Read(b []byte) (int, error) {
	n, err := r.s.Read(r.ctx, b)
	if err != nil {
		return 0, err
	}
	if n == 0 {
		return 0, io.EOF
	}
	return n, nil
}

// Close implements io.Closer.
func (r *StreamReader) Close() error {
	return r.s.Close(r.ctx)
}

// StreamWriter wraps around a gio.OutputStreamer.
type StreamWriter struct {
	s   *gio.OutputStream
	ctx context.Context
}

// Writer wraps a gio.OutputStreamer to provide an io.WriteCloser with flushing
// capability.
func Writer(ctx context.Context, s gio.OutputStreamer) *StreamWriter {
	return &StreamWriter{s.BaseOutputStream(), ctx}
}

// Write implements io.Writer.
func (w *StreamWriter) Write(b []byte) (int, error) {
	return w.s.Write(w.ctx, b)
}

// ReadFrom implements io.ReaderFrom. It has a fast path for gio.InputStreamers
// wrapped using gioutil.Reader.
func (w *StreamWriter) ReadFrom(r io.Reader) (int64, error) {
	streamer, ok := r.(*StreamReader)
	if ok {
		n, err := w.s.Splice(w.ctx, streamer.s, 0)
		return int64(n), err
	}

	buf := make([]byte, 32*1024)

	var written int64
	var err error

	// Code taken from io.Copy to avoid an infinite recursion.
	for {
		nr, er := r.Read(buf)
		if nr > 0 {
			nw, ew := w.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errors.New("invalid write return")
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	return written, err
}

// Close implements io.Closer.
func (w *StreamWriter) Close() error {
	return w.s.Close(w.ctx)
}

// Flush flushes the writer. See gio.OutputStreamer.Flush.
func (w *StreamWriter) Flush() error {
	return w.s.Flush(w.ctx)
}

type seeker struct {
	s   gio.Seekabler
	ctx context.Context
}

// Seeker wraps around a gio.Seekable.
func Seeker(ctx context.Context, s gio.Seekabler) io.Seeker {
	return seeker{s, ctx}
}

func (s seeker) Seek(offset int64, whence int) (int64, error) {
	var typ glib.SeekType

	switch whence {
	case io.SeekStart:
		typ = glib.SeekSet
	case io.SeekCurrent:
		typ = glib.SeekCur
	case io.SeekEnd:
		typ = glib.SeekEnd
	default:
		return 0, fmt.Errorf("unknown whence %d", whence)
	}

	if err := s.s.Seek(s.ctx, offset, typ); err != nil {
		return 0, err
	}

	return s.s.Tell(), nil
}
