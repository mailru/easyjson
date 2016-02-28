package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

const structComment = "easyjson:json"

type Parser struct {
	PkgPath     string
	PkgName     string
	StructNames []string
	AllStructs  bool
}

type visitor struct {
	*Parser

	name string
}

func (p *Parser) needStruct(comments string) bool {
	if p.AllStructs {
		return true
	}
	for _, v := range strings.Split(comments, "\n") {
		if strings.HasPrefix(v, structComment) {
			return true
		}
	}
	return false
}

func (v *visitor) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.File:
		v.PkgName = n.Name.String()
		return v

	case *ast.GenDecl:
		if !v.needStruct(n.Doc.Text()) {
			return nil
		}
		return v
	case *ast.TypeSpec:
		v.name = n.Name.String()
		return v
	case *ast.StructType:
		v.StructNames = append(v.StructNames, v.name)
		return nil
	}
	return nil
}

func (p *Parser) Parse(fname string) error {
	var err error
	if p.PkgPath, err = getPkgPath(fname); err != nil {
		return err
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Walk(&visitor{Parser: p}, f)
	return nil
}
