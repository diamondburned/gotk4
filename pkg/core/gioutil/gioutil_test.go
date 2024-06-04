package gioutil

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"testing"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// newJunk returns a byte slice of new random junk.
func newJunk(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func assertReader(t *testing.T, r io.Reader, eq []byte) {
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal("cannot read:", err)
	}

	if !bytes.Equal(b, eq) {
		t.Logf("got    %q", b)
		t.Logf("expect %q", eq)
		t.Fatal("bytes mismatch from io.Reader")
	}
}

func TestInputStream(t *testing.T) {
	b := newJunk(128)

	istream := NewInputStream(bytes.NewReader(b))
	greader := Reader(context.Background(), istream)

	assertReader(t, greader, b)
}

func TestOutputStream(t *testing.T) {
	b := newJunk(128)
	var out bytes.Buffer

	ostream := NewOutputStream(&out)
	gwriter := Writer(context.Background(), ostream)

	_, err := gwriter.Write(b)
	if err != nil {
		t.Fatal("cannot write into gio.OutputStream:", err)
	}

	assertReader(t, &out, b)
}

var (
	_ io.Reader = (*bytes.Reader)(nil)
	_ io.Seeker = (*bytes.Reader)(nil)
)

func TestInputStreamSeeker(t *testing.T) {
	r := bytes.NewReader(newJunk(128))
	istream := NewInputStream(r)

	sstream := &gio.Seekable{Object: istream.Object}
	if !sstream.CanSeek() {
		t.Fatal("bytes.Reader is not seekable over gio.InputStream")
	}

	if err := sstream.Seek(context.Background(), 2, glib.SeekSet); err != nil {
		t.Fatal("gio.InputStream.Seek(SeekSet) failed:", err)
	}
	assertCursorPos(t, r, sstream, 2)

	if err := sstream.Seek(context.Background(), 1, glib.SeekCur); err != nil {
		t.Fatal("gio.InputStream.Seek(SeekCur) failed:", err)
	}
	assertCursorPos(t, r, sstream, 3)

	if err := sstream.Seek(context.Background(), 0, glib.SeekEnd); err != nil {
		t.Fatal("gio.InputStream.Seek(SeekEnd) failed:", err)
	}
	assertCursorPos(t, r, sstream, r.Size())
}

func assertCursorPos(t *testing.T, s io.Seeker, gs *gio.Seekable, want int64) {
	t.Helper()

	n, err := s.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatal("Seek() failed:", err)
	}

	if n != want {
		t.Fatalf("did not seek to the correct position (got %d, want %d)", n, want)
	}

	if tell := gs.Tell(); n != tell {
		t.Fatalf("gio.Seekable.Tell() returned an incorrect value (got %d, want %d)", tell, n)
	}
}

const (
	benchBytes = 512 << 10 // 512KB
	sinkBytes  = 512
)

func sinker() func(io.Reader) {
	sinkhole := make([]byte, sinkBytes)
	return func(r io.Reader) { r.Read(sinkhole) }
}

func BenchmarkInputStream(b *testing.B) {
	big := newJunk(benchBytes)
	sink := sinker()
	b.SetBytes(benchBytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := NewInputStream(bytes.NewReader(big))
		r := Reader(context.Background(), s)
		sink(r)
	}
}

func BenchmarkReader(b *testing.B) {
	big := newJunk(benchBytes)
	sink := sinker()
	b.SetBytes(benchBytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sink(bytes.NewReader(big))
	}
}
