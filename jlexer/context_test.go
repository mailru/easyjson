package jlexer

import (
	"fmt"
	"testing"
)

type TestWalker struct {
	ErrorPrefix string
	T           *testing.T
	Items       []string
}

func (w *TestWalker) item(s string) {
	if len(w.Items) == 0 {
		w.T.Errorf("%sTestWalker(): no items left; want %q", w.ErrorPrefix, s)
	} else if w.Items[0] != s {
		w.T.Errorf("%sTestWalker(): got %q; want %q", w.ErrorPrefix, s, w.Items[0])
		w.Items = w.Items[1:]
	} else {
		w.Items = w.Items[1:]
	}
}

func (w *TestWalker) OnEnterObject()             { w.item("{") }
func (w *TestWalker) OnExitObject()              { w.item("}") }
func (w *TestWalker) OnNextObjectKey(key string) { w.item("e:" + key) }
func (w *TestWalker) OnEnterArray()              { w.item("[") }
func (w *TestWalker) OnNextArrayElement()        { w.item("e") }
func (w *TestWalker) OnExitArray()               { w.item("]") }

func TestWalkUpToPosition(t *testing.T) {
	for i, test := range []struct {
		JSON       string
		Items      []string
		Start, End int
	}{
		{
			JSON:  ``,
			Items: []string{},
			End:   -1,
		}, {
			JSON:  `{"aaa": 5, "qqq": 10}`,
			Items: []string{"{", "e:aaa", "e:qqq", "}"},
			End:   -1,
		}, {
			JSON:  `{"aaa": 5, "qqq": {"\t\t": null}}`,
			Items: []string{"{", "e:aaa", "e:qqq", "{", "e:\t\t", "}", "}"},
			End:   -1,
		}, {
			JSON:  `{"aaa": 5, "qqq": 10}`,
			End:   len(`{"aaa": 5, "qqq": `),
			Items: []string{"{", "e:aaa", "e:qqq"},
		}, {
			JSON:  `{"aaa": 5, "qqq": 10}`,
			End:   len(`{"aaa": 5, `),
			Items: []string{"{", "e:aaa"},
		}, {
			JSON:  `[null, false, {"aaa": 5}]`,
			Items: []string{"[", "e", "e", "{", "e:aaa", "}", "]"},
			End:   -1,
		}, {
			JSON:  `[null, "aaa"]`,
			Items: []string{"[", "e", "e", "]"},
			End:   -1,
		},
	} {
		l := &Lexer{Data: []byte(test.JSON)}

		if test.End != -1 {
			l.pos = test.End
		} else {
			l.pos = len(test.JSON)
		}
		w := &TestWalker{T: t, Items: test.Items, ErrorPrefix: fmt.Sprintf("[%d,%q] ", i, test.JSON)}

		if err := l.WalkUpToPosition(w); err != nil {
			t.Errorf("[%d,%q] WalkUpToPosition() error: %v", i, test.JSON, err)
		}

		if len(w.Items) > 0 {
			t.Errorf("[%d,%q] WalkUpToPosition: items %q left", i, test.JSON, w.Items)
		}
	}
}
