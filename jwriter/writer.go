// Package jwriter contains a JSON writer.
package jwriter

import (
	"io"
	"strconv"
	"unicode/utf8"

	"github.com/mailru/easyjson/buffer"
)

// Flags describe various encoding options. The behavior may be actually implemented in the encoder, but
// Flags field in Writer is used to set and pass them around.
type Flags int

const (
	NilMapAsEmpty   Flags = 1 << iota // Encode nil map as '{}' rather than 'null'.
	NilSliceAsEmpty                   // Encode nil slice as '[]' rather than 'null'.
)

type Writer interface {
	Flags() Flags

	// ReadCloser returns an io.ReadCloser that can be used to read the data.
	// ReadCloser also resets the buffer.
	ReadCloser() (io.ReadCloser, error)

	// RawByte appends raw binary data to the buffer.
	RawByte(c byte) error

	// RawByte appends raw binary data to the buffer.
	RawBytes(data []byte) error

	RawBytesWithErr(data []byte, err error) error

	// RawByte appends raw binary data to the buffer.
	RawString(s string) error

	// Base64Bytes appends data to the buffer after base64 encoding it
	Base64Bytes(data []byte) error

	Uint8(n uint8) error

	Uint16(n uint16) error

	Uint32(n uint32) error

	Uint(n uint) error

	Uint64(n uint64) error

	Int8(n int8) error

	Int16(n int16) error

	Int32(n int32) error

	Int(n int) error

	Int64(n int64) error

	Uint8Str(n uint8) error

	Uint16Str(n uint16) error

	Uint32Str(n uint32) error

	UintStr(n uint) error

	Uint64Str(n uint64) error

	UintptrStr(n uintptr) error

	Int8Str(n int8) error

	Int16Str(n int16) error

	Int32Str(n int32) error

	IntStr(n int) error

	Int64Str(n int64) error

	Float32(n float32) error

	Float32Str(n float32) error

	Float64(n float64) error

	Float64Str(n float64) error

	Bool(v bool) error

	String(s string) error
}

// Writer is a JSON writer.
type BufWriter struct {
	flags Flags

	Buffer       buffer.Buffer
	NoEscapeHTML bool
}

// Size returns the size of the data that was written out.
func (w *BufWriter) Flags() Flags {
	return w.flags
}

// Size returns the size of the data that was written out.
func (w *BufWriter) Size() int {
	return w.Buffer.Size()
}

// DumpTo outputs the data to given io.Writer, resetting the buffer.
func (w *BufWriter) DumpTo(out io.Writer) (written int, err error) {
	return w.Buffer.DumpTo(out)
}

// BuildBytes returns writer data as a single byte slice. You can optionally provide one byte slice
// as argument that it will try to reuse.
func (w *BufWriter) BuildBytes(reuse ...[]byte) ([]byte, error) {
	return w.Buffer.BuildBytes(reuse...), nil
}

// ReadCloser returns an io.ReadCloser that can be used to read the data.
// ReadCloser also resets the buffer.
func (w *BufWriter) ReadCloser() (io.ReadCloser, error) {
	return w.Buffer.ReadCloser(), nil
}

// RawByte appends raw binary data to the buffer.
func (w *BufWriter) RawByte(c byte) error {
	w.Buffer.AppendByte(c)
	return nil
}

// RawByte appends raw binary data to the buffer.
func (w *BufWriter) RawBytes(data []byte) error {
	w.Buffer.AppendBytes(data)
	return nil
}

func (w *BufWriter) RawBytesWithErr(data []byte, err error) error {
	if err != nil {
		return err
	}
	w.RawBytes(data)
	return nil
}

// RawByte appends raw binary data to the buffer.
func (w *BufWriter) RawString(s string) error {
	w.Buffer.AppendString(s)
	return nil
}

// Base64Bytes appends data to the buffer after base64 encoding it
func (w *BufWriter) Base64Bytes(data []byte) error {
	if data == nil {
		w.Buffer.AppendString("null")
		return nil
	}
	w.Buffer.AppendByte('"')
	base64(&w.Buffer, data)
	w.Buffer.AppendByte('"')
	return nil
}

func (w *BufWriter) Uint8(n uint8) error {
	w.Buffer.EnsureSpace(3)
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	return nil
}

func (w *BufWriter) Uint16(n uint16) error {
	w.Buffer.EnsureSpace(5)
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	return nil
}

func (w *BufWriter) Uint32(n uint32) error {
	w.Buffer.EnsureSpace(10)
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	return nil
}

func (w *BufWriter) Uint(n uint) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	return nil
}

func (w *BufWriter) Uint64(n uint64) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, n, 10)
	return nil
}

func (w *BufWriter) Int8(n int8) error {
	w.Buffer.EnsureSpace(4)
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	return nil
}

func (w *BufWriter) Int16(n int16) error {
	w.Buffer.EnsureSpace(6)
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	return nil
}

func (w *BufWriter) Int32(n int32) error {
	w.Buffer.EnsureSpace(11)
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	return nil
}

func (w *BufWriter) Int(n int) error {
	w.Buffer.EnsureSpace(21)
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	return nil
}

func (w *BufWriter) Int64(n int64) error {
	w.Buffer.EnsureSpace(21)
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, n, 10)
	return nil
}

func (w *BufWriter) Uint8Str(n uint8) error {
	w.Buffer.EnsureSpace(3)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Uint16Str(n uint16) error {
	w.Buffer.EnsureSpace(5)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Uint32Str(n uint32) error {
	w.Buffer.EnsureSpace(10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) UintStr(n uint) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Uint64Str(n uint64) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, n, 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) UintptrStr(n uintptr) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendUint(w.Buffer.Buf, uint64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Int8Str(n int8) error {
	w.Buffer.EnsureSpace(4)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Int16Str(n int16) error {
	w.Buffer.EnsureSpace(6)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Int32Str(n int32) error {
	w.Buffer.EnsureSpace(11)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) IntStr(n int) error {
	w.Buffer.EnsureSpace(21)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, int64(n), 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Int64Str(n int64) error {
	w.Buffer.EnsureSpace(21)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendInt(w.Buffer.Buf, n, 10)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Float32(n float32) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = strconv.AppendFloat(w.Buffer.Buf, float64(n), 'g', -1, 32)
	return nil
}

func (w *BufWriter) Float32Str(n float32) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendFloat(w.Buffer.Buf, float64(n), 'g', -1, 32)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Float64(n float64) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = strconv.AppendFloat(w.Buffer.Buf, n, 'g', -1, 64)
	return nil
}

func (w *BufWriter) Float64Str(n float64) error {
	w.Buffer.EnsureSpace(20)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	w.Buffer.Buf = strconv.AppendFloat(w.Buffer.Buf, float64(n), 'g', -1, 64)
	w.Buffer.Buf = append(w.Buffer.Buf, '"')
	return nil
}

func (w *BufWriter) Bool(v bool) error {
	w.Buffer.EnsureSpace(5)
	if v {
		w.Buffer.Buf = append(w.Buffer.Buf, "true"...)
	} else {
		w.Buffer.Buf = append(w.Buffer.Buf, "false"...)
	}
	return nil
}

const chars = "0123456789abcdef"

func getTable(falseValues ...int) [128]bool {
	table := [128]bool{}

	for i := 0; i < 128; i++ {
		table[i] = true
	}

	for _, v := range falseValues {
		table[v] = false
	}

	return table
}

var (
	htmlEscapeTable   = getTable(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, '"', '&', '<', '>', '\\')
	htmlNoEscapeTable = getTable(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, '"', '\\')
)

func (w *BufWriter) String(s string) error {
	w.Buffer.AppendByte('"')

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

			w.Buffer.AppendString(s[p:i])
			switch c {
			case '\t':
				w.Buffer.AppendString(`\t`)
			case '\r':
				w.Buffer.AppendString(`\r`)
			case '\n':
				w.Buffer.AppendString(`\n`)
			case '\\':
				w.Buffer.AppendString(`\\`)
			case '"':
				w.Buffer.AppendString(`\"`)
			default:
				w.Buffer.AppendString(`\u00`)
				w.Buffer.AppendByte(chars[c>>4])
				w.Buffer.AppendByte(chars[c&0xf])
			}

			i++
			p = i
			continue
		}

		// broken utf
		runeValue, runeWidth := utf8.DecodeRuneInString(s[i:])
		if runeValue == utf8.RuneError && runeWidth == 1 {
			w.Buffer.AppendString(s[p:i])
			w.Buffer.AppendString(`\ufffd`)
			i++
			p = i
			continue
		}

		// jsonp stuff - tab separator and line separator
		if runeValue == '\u2028' || runeValue == '\u2029' {
			w.Buffer.AppendString(s[p:i])
			w.Buffer.AppendString(`\u202`)
			w.Buffer.AppendByte(chars[runeValue&0xf])
			i += runeWidth
			p = i
			continue
		}
		i += runeWidth
	}
	w.Buffer.AppendString(s[p:])
	w.Buffer.AppendByte('"')
	return nil
}

const encode = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const padChar = '='

func base64(b *buffer.Buffer, in []byte) {

	if len(in) == 0 {
		return
	}

	b.EnsureSpace(((len(in)-1)/3 + 1) * 4)

	si := 0
	n := (len(in) / 3) * 3

	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(in[si+0])<<16 | uint(in[si+1])<<8 | uint(in[si+2])

		b.Buf = append(b.Buf, encode[val>>18&0x3F], encode[val>>12&0x3F], encode[val>>6&0x3F], encode[val&0x3F])

		si += 3
	}

	remain := len(in) - si
	if remain == 0 {
		return
	}

	// Add the remaining small block
	val := uint(in[si+0]) << 16
	if remain == 2 {
		val |= uint(in[si+1]) << 8
	}

	b.Buf = append(b.Buf, encode[val>>18&0x3F], encode[val>>12&0x3F])

	switch remain {
	case 2:
		b.Buf = append(b.Buf, encode[val>>6&0x3F], byte(padChar))
	case 1:
		b.Buf = append(b.Buf, byte(padChar), byte(padChar))
	}
}
