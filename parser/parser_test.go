package parser

import "testing"

func TestGetPkgPath(t *testing.T) {
	cases := []struct {
		fname string
		isDir bool
	}{
		{"parser.go", false},
		{".", true},
	}
	exp := "github.com/mailru/easyjson/parser"

	for _, tc := range cases {

		pkg, err := getPkgPath(tc.fname, tc.isDir)
		if err != nil {
			t.Error(err)
		}
		if pkg != exp {
			t.Errorf("in: \"%s\" isDir: %v want: %s got: %s", tc.fname, tc.isDir, exp, pkg)
		}
	}

}
