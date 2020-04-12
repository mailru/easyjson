package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
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
		explicit := v.needType(n.Doc.Text())
		if !explicit {
			return v
		}

		for _, nc := range n.Specs {
			switch nct := nc.(type) {
			case *ast.TypeSpec:
				nct.Doc = n.Doc
			}
		}

		return v
	case *ast.TypeSpec:
		explicit := v.needType(n.Doc.Text())
		if !explicit && !v.AllStructs {
			return nil
		}

		v.name = n.Name.String()

		// Allow to specify non-structs explicitly independent of '-all' flag.
		if explicit {
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

	fset := token.NewFileSet()
	if isDir {
		packages, err := parser.ParseDir(fset, fname, excludeTestFiles, parser.ParseComments)
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

func excludeTestFiles(fi os.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go")
}
