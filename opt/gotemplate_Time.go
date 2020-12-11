// generated by gotemplate

package opt

import (
	"fmt"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"time"
)

// template type Optional(A)

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Time struct {
	V       time.Time
	Defined bool
}

// Creates an optional type with a given value.
func OTime(v time.Time) Time {
	return Time{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Time) Get(deflt time.Time) time.Time {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Time) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Time(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Time) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Time{}
	} else {
		v.V = l.Time()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Time) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Time) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Time) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Time) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}
