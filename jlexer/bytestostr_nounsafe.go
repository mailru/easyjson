//+build easyjson_nounsafe appengine

package jlexer

// bytesToStr creates a string normally from []byte
//
// Note that this method is roughly 1.5x slower than using the 'unsafe' method.
func bytesToStr(data []byte) string {
	return string(data)
}
