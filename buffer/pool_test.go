package buffer

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestAppendByte(t *testing.T) {
	var b Buffer
	var want []byte

	for i := 0; i < 1000; i++ {
		b.AppendByte(1)
		b.AppendByte(2)
		want = append(want, 1, 2)
	}

	got := b.BuildBytes()
	if !bytes.Equal(got, want) {
		t.Errorf("BuildBytes() = %v; want %v", got, want)
	}
}

func TestAppendBytes(t *testing.T) {
	var b Buffer
	var want []byte

	for i := 0; i < 1000; i++ {
		b.AppendBytes([]byte{1, 2})
		want = append(want, 1, 2)
	}

	got := b.BuildBytes()
	if !bytes.Equal(got, want) {
		t.Errorf("BuildBytes() = %v; want %v", got, want)
	}
}

func TestAppendString(t *testing.T) {
	var b Buffer
	var want []byte

	s := "test"
	for i := 0; i < 1000; i++ {
		b.AppendString(s)
		want = append(want, s...)
	}

	got := b.BuildBytes()
	if !bytes.Equal(got, want) {
		t.Errorf("BuildBytes() = %v; want %v", got, want)
	}
}

func TestDumpTo(t *testing.T) {
	var b Buffer
	var want []byte

	s := "test"
	for i := 0; i < 1000; i++ {
		b.AppendBytes([]byte(s))
		want = append(want, s...)
	}

	out := &bytes.Buffer{}
	n, err := b.DumpTo(out)
	if err != nil {
		t.Errorf("DumpTo() error: %v", err)
	}

	got := out.Bytes()
	if !bytes.Equal(got, want) {
		t.Errorf("DumpTo(): got %v; want %v", got, want)
	}

	if n != len(want) {
		t.Errorf("DumpTo() = %v; want %v", n, len(want))
	}
}

type failingWriter struct {
	out       bytes.Buffer
	remaining int
	err       error
}

func (w *failingWriter) Write(p []byte) (int, error) {
	n := len(p)
	if n > w.remaining {
		n = w.remaining
	}
	n, _ = w.out.Write(p[:n])
	w.remaining -= n
	if n < len(p) {
		return n, w.err
	}
	return n, nil
}

func TestDumpToWriteError(t *testing.T) {
	data := make([]byte, 4096)
	for i := range data {
		data[i] = byte(i % 251)
	}
	for _, limit := range []int{0, 128, 300, 1700, len(data)} {
		t.Run(fmt.Sprint(limit), func(t *testing.T) {
			var b Buffer
			b.AppendBytes(data)
			wantErr := errors.New("write failed")
			w := &failingWriter{remaining: limit, err: wantErr}
			n, err := b.DumpTo(w)
			if limit < len(data) {
				if !errors.Is(err, wantErr) {
					t.Fatalf("DumpTo() error = %v, want %v", err, wantErr)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if n != limit || !bytes.Equal(w.out.Bytes(), data[:limit]) {
				t.Fatalf("DumpTo() wrote %d bytes, want prefix of %d bytes", n, limit)
			}
			if b.Size() != 0 {
				t.Fatalf("buffer not reset after DumpTo(): %d bytes remain", b.Size())
			}
			b.AppendBytes(data)
			var out bytes.Buffer
			n, err = b.DumpTo(&out)
			if err != nil || n != len(data) || !bytes.Equal(out.Bytes(), data) {
				t.Fatalf("reused buffer: wrote %d bytes, error %v", n, err)
			}
		})
	}
}

func BenchmarkDumpToPooledChunks(b *testing.B) {
	for _, size := range []int{0, 128, 16 * 1024} {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			data := bytes.Repeat([]byte("x"), size)
			b.ReportAllocs()
			b.SetBytes(int64(size))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var buf Buffer
				buf.AppendBytes(data)
				if n, err := buf.DumpTo(io.Discard); err != nil || n != size {
					b.Fatalf("DumpTo() = %d, %v; want %d, nil", n, err, size)
				}
			}
		})
	}
}

func TestReadCloser(t *testing.T) {
	var b Buffer
	var want []byte

	s := "test"
	for i := 0; i < 1000; i++ {
		b.AppendBytes([]byte(s))
		want = append(want, s...)
	}

	out := &bytes.Buffer{}
	rc := b.ReadCloser()
	n, err := out.ReadFrom(rc)
	if err != nil {
		t.Errorf("ReadCloser() error: %v", err)
	}
	rc.Close() // Will always return nil

	got := out.Bytes()
	if !bytes.Equal(got, want) {
		t.Errorf("DumpTo(): got %v; want %v", got, want)
	}

	if n != int64(len(want)) {
		t.Errorf("DumpTo() = %v; want %v", n, len(want))
	}
}
