package gioutil

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"testing"
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
