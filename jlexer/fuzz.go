package jlexer

func Fuzz(data []byte) int {
	l := Lexer{Data: data}
	_ = l.UnsafeString()
	return 1
}
