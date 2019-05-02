package parser

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const structComment = "easyjson:json"

type Parser struct {
	PkgPath     string
	PkgName     string
	StructNames []string
	AllStructs  bool

	workingDirectory string
}

type visitor struct {
	*Parser

	name     string
	explicit bool
}

func (p *Parser) needType(comments string) bool {
	for _, v := range strings.Split(comments, "\n") {
		if strings.HasPrefix(v, structComment) {
			return true
		}
	}
	return false
}

func (v *visitor) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		return v
	case *ast.File:
		v.PkgName = n.Name.String()
		return v

	case *ast.GenDecl:
		v.explicit = v.needType(n.Doc.Text())

		if !v.explicit && !v.AllStructs {
			return nil
		}
		return v
	case *ast.TypeSpec:
		v.name = n.Name.String()

		// Allow to specify non-structs explicitly independent of '-all' flag.
		if v.explicit {
			v.StructNames = append(v.StructNames, v.name)
			return nil
		}
		return v
	case *ast.StructType:
		v.StructNames = append(v.StructNames, v.name)
		return nil
	}
	return nil
}

func (p *Parser) Parse(fname string, isDir bool) error {
	var err error
	if p.PkgPath, err = getPkgPath(fname, isDir); err != nil {
		return err
	}
	if isDir {
		p.workingDirectory = filepath.Clean(fname)
	}

	fset := token.NewFileSet()
	if isDir {
		packages, err := parser.ParseDir(fset, fname, p.isFileAllowed, parser.ParseComments)
		if err != nil {
			return err
		}

		for _, pckg := range packages {
			ast.Walk(&visitor{Parser: p}, pckg)
		}
	} else {
		f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		ast.Walk(&visitor{Parser: p}, f)
	}
	return nil
}

func getDefaultGoPath() (string, error) {
	output, err := exec.Command("go", "env", "GOPATH").Output()
	return string(bytes.TrimSpace(output)), err
}

var prefixBuildTagIgnore = []byte(`// +build ignore`)
var prefixPackage = []byte(`package `)

func (p *Parser) isFileAllowed(fi os.FileInfo) bool {
	if strings.HasSuffix(fi.Name(), "_test.go") {
		return false
	}

	f, err := os.Open(filepath.Join(p.workingDirectory, fi.Name()))
	if err != nil {
		return false
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if bytes.HasPrefix(s.Bytes(), prefixBuildTagIgnore) {
			return false
		}
		if bytes.HasPrefix(s.Bytes(), prefixPackage) {
			return true // no need to process the whole file
		}
	}
	return true
}
