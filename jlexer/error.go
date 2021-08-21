package jlexer

import "strconv"

// LexerError implements the error interface and represents all possible errors that can be
// generated during parsing the JSON data.
type LexerError struct {
	Reason string
	Offset int
	Data   string
}

func (l *LexerError) Error() string {
	msg := "parse error: " + l.Reason + " near offset " + strconv.Itoa(l.Offset) + "of '" + l.Data + "'"
	return msg
}

// This is a (temporary?) helper to use in place of errors.New
func NewError(msg string) error {
	return myError{msg: msg}
}

type myError struct {
	msg string
}

func (m myError) Error() string {
	return m.msg
}
