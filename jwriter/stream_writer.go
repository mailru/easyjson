package jwriter

// This file defines the easyjson-based implementation of the low-level JSON writer, which is used instead
// of token_writer_default.go if the launchdarkly_easyjson build tag is enabled.
//
// For the contract governing the behavior of the exported methods in this type, see the comments on the
// corresponding methods in token_writer_default.go.

import (
	"io"
	"strconv"
	"unicode/utf8"

	"github.com/mailru/easyjson/buffer"
)

var nullToken = []byte("null")

type tokenWriter struct {
	flags            Flags
	buffer           buffer.Buffer
	targetIOWriter   io.Writer
	targetBufferSize int
	NoEscapeHTML     bool
}

func newStreamingTokenWriter(dest io.Writer, bufferSize int) tokenWriter {
	tw := tokenWriter{
		targetIOWriter:   dest,
		targetBufferSize: bufferSize,
	}

	tw.buffer.EnsureSpace(bufferSize)
	return tw
}

func (w *tokenWriter) Flags() Flags {
	return w.flags
}

func (tw *tokenWriter) Flush() error {
	if tw.targetIOWriter == nil {
		return nil
	}
	_, err := tw.buffer.DumpTo(tw.targetIOWriter)
	return err
}

func (tw *tokenWriter) maybeFlush() error {
	if tw.targetIOWriter == nil || tw.buffer.Size() < tw.targetBufferSize {
		return nil
	}
	return tw.Flush()
}

// -------

// ReadCloser returns an io.ReadCloser that can be used to read the data.
// ReadCloser also resets the buffer.
func (w *tokenWriter) ReadCloser() (io.ReadCloser, error) {
	return w.buffer.ReadCloser(), nil
}

// RawByte appends raw binary data to the buffer.
func (w *tokenWriter) RawByte(c byte) error {
	w.buffer.AppendByte(c)
	return w.maybeFlush()
}

// RawByte appends raw binary data to the buffer.
func (w *tokenWriter) RawBytes(data []byte) error {
	w.buffer.AppendBytes(data)
	return w.maybeFlush()
}

func (w *tokenWriter) RawBytesWithErr(data []byte, err error) error {
	if err != nil {
		return err
	}
	w.RawBytes(data)
	return nil
}

// RawByte appends raw binary data to the buffer.
func (w *tokenWriter) RawString(s string) error {
	w.buffer.AppendString(s)
	return w.maybeFlush()
}

// Base64Bytes appends data to the buffer after base64 encoding it
func (w *tokenWriter) Base64Bytes(data []byte) error {
	if data == nil {
		w.buffer.AppendString("null")
		return w.maybeFlush()
	}
	w.buffer.AppendByte('"')
	base64(&w.buffer, data)
	w.buffer.AppendByte('"')
	return w.maybeFlush()
}

func (w *tokenWriter) Uint8(n uint8) error {
	w.buffer.EnsureSpace(3)
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Uint16(n uint16) error {
	w.buffer.EnsureSpace(5)
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Uint32(n uint32) error {
	w.buffer.EnsureSpace(10)
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Uint(n uint) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Uint64(n uint64) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, n, 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Int8(n int8) error {
	w.buffer.EnsureSpace(4)
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Int16(n int16) error {
	w.buffer.EnsureSpace(6)
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Int32(n int32) error {
	w.buffer.EnsureSpace(11)
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Int(n int) error {
	w.buffer.EnsureSpace(21)
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Int64(n int64) error {
	w.buffer.EnsureSpace(21)
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, n, 10)
	return w.maybeFlush()
}

func (w *tokenWriter) Uint8Str(n uint8) error {
	w.buffer.EnsureSpace(3)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Uint16Str(n uint16) error {
	w.buffer.EnsureSpace(5)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Uint32Str(n uint32) error {
	w.buffer.EnsureSpace(10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) UintStr(n uint) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Uint64Str(n uint64) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, n, 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) UintptrStr(n uintptr) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendUint(w.buffer.Buf, uint64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Int8Str(n int8) error {
	w.buffer.EnsureSpace(4)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Int16Str(n int16) error {
	w.buffer.EnsureSpace(6)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Int32Str(n int32) error {
	w.buffer.EnsureSpace(11)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) IntStr(n int) error {
	w.buffer.EnsureSpace(21)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, int64(n), 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Int64Str(n int64) error {
	w.buffer.EnsureSpace(21)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendInt(w.buffer.Buf, n, 10)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Float32(n float32) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = strconv.AppendFloat(w.buffer.Buf, float64(n), 'g', -1, 32)
	return w.maybeFlush()
}

func (w *tokenWriter) Float32Str(n float32) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendFloat(w.buffer.Buf, float64(n), 'g', -1, 32)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Float64(n float64) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = strconv.AppendFloat(w.buffer.Buf, n, 'g', -1, 64)
	return w.maybeFlush()
}

func (w *tokenWriter) Float64Str(n float64) error {
	w.buffer.EnsureSpace(20)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	w.buffer.Buf = strconv.AppendFloat(w.buffer.Buf, float64(n), 'g', -1, 64)
	w.buffer.Buf = append(w.buffer.Buf, '"')
	return w.maybeFlush()
}

func (w *tokenWriter) Bool(v bool) error {
	w.buffer.EnsureSpace(5)
	if v {
		w.buffer.Buf = append(w.buffer.Buf, "true"...)
	} else {
		w.buffer.Buf = append(w.buffer.Buf, "false"...)
	}
	return w.maybeFlush()
}

func (w *tokenWriter) String(s string) error {
	w.buffer.AppendByte('"')

	// Portions of the string that contain no escapes are appended as
	// byte slices.

	p := 0 // last non-escape symbol

	escapeTable := &htmlEscapeTable
	if w.NoEscapeHTML {
		escapeTable = &htmlNoEscapeTable
	}

	for i := 0; i < len(s); {
		c := s[i]

		if c < utf8.RuneSelf {
			if escapeTable[c] {
				// single-width character, no escaping is required
				i++
				continue
			}

			w.buffer.AppendString(s[p:i])
			switch c {
			case '\t':
				w.buffer.AppendString(`\t`)
			case '\r':
				w.buffer.AppendString(`\r`)
			case '\n':
				w.buffer.AppendString(`\n`)
			case '\\':
				w.buffer.AppendString(`\\`)
			case '"':
				w.buffer.AppendString(`\"`)
			default:
				w.buffer.AppendString(`\u00`)
				w.buffer.AppendByte(chars[c>>4])
				w.buffer.AppendByte(chars[c&0xf])
			}

			i++
			p = i
			continue
		}

		// broken utf
		runeValue, runeWidth := utf8.DecodeRuneInString(s[i:])
		if runeValue == utf8.RuneError && runeWidth == 1 {
			w.buffer.AppendString(s[p:i])
			w.buffer.AppendString(`\ufffd`)
			i++
			p = i
			continue
		}

		// jsonp stuff - tab separator and line separator
		if runeValue == '\u2028' || runeValue == '\u2029' {
			w.buffer.AppendString(s[p:i])
			w.buffer.AppendString(`\u202`)
			w.buffer.AppendByte(chars[runeValue&0xf])
			i += runeWidth
			p = i
			continue
		}
		i += runeWidth
	}
	w.buffer.AppendString(s[p:])
	w.buffer.AppendByte('"')
	return w.maybeFlush()
}
