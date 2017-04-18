package parser

import (
	"strings"
)

func normalizePath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}
