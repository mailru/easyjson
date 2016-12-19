package jlexer

import "io"

type Walker interface {
	OnEnterObject()
	OnNextObjectKey(key string)
	OnExitObject()

	OnEnterArray()
	OnNextArrayElement()
	OnExitArray()
}

func (l *Lexer) WalkUpToPosition(w Walker) error {
	l1 := &Lexer{}
	l1.Data = l.Data

	type StackItem int
	const (
		Object StackItem = iota
		Array
	)

	var stack []StackItem
	haveKey := false

	for {
		l1.fetchToken()
		if l1.pos > l.pos || !l1.Ok() {
			break
		}
		switch {
		case l1.IsDelim('{'):
			l1.Skip()

			stack = append(stack, Object)
			w.OnEnterObject()
			haveKey = false

		case l1.IsDelim('}'):
			l1.Skip()

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				l1.WantComma()
			}
			w.OnExitObject()
			haveKey = false

		case l1.IsDelim('['):
			l1.Skip()

			stack = append(stack, Array)
			w.OnEnterArray()
			haveKey = false

		case l1.IsDelim(']'):
			l1.Skip()

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				l1.WantComma()
			}
			w.OnExitArray()
			haveKey = false

		case len(stack) > 0 && stack[len(stack)-1] == Object && !haveKey:
			key := l1.UnsafeString()
			w.OnNextObjectKey(key)

			l1.WantColon()
			haveKey = true

		case len(stack) > 0 && stack[len(stack)-1] == Array:
			w.OnNextArrayElement()
			l1.Skip()
			l1.WantComma()

		default:
			l1.Skip()
			l1.WantComma()
			haveKey = false
		}

	}

	if l1.Error() == io.EOF {
		return nil
	}
	return l1.Error()
}
