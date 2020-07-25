package parser

import "flag"

var allStructs = flag.Bool("all", false, "generate marshaler/unmarshalers for all structs in a file")

func Fuzz(data []byte) int {
	p := Parser{AllStructs: *allStructs}
	err := p.Parse(string(data), false)
	if err != nil {
		return 0
	}
	return 1
}
