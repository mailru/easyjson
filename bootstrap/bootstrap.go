// Package bootstrap implements the bootstrapping logic: generation of a .go file to
// launch the actual generator and launching the generator itself.
//
// The package may be preferred to a command-line utility if generating the serializers
// from golang code is required.
package bootstrap

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const genPackage = "github.com/mailru/easyjson/gen"
const pkgWriter = "github.com/mailru/easyjson/jwriter"
const pkgLexer = "github.com/mailru/easyjson/jlexer"

type Generator struct {
	PkgPath, PkgName string
	Types            []string

	NoStdMarshalers bool
	SnakeCase       bool
	OmitEmpty       bool

	OutName   string
	BuildTags string

	StubsOnly  bool
	LeaveTemps bool
}

// writeStub outputs an initial stubs for marshalers/unmarshalers so that the package
// using marshalers/unmarshales compiles correctly for boostrapping code.
func (g *Generator) writeStub() error {
	f, err := os.Create(g.OutName)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, "package ", g.PkgName)
	fmt.Fprintln(f)

	for _, t := range g.Types {
		fmt.Fprintln(f, "func (*", t, ") MarshalJSON() ([]byte, error) { return nil, nil }")
		fmt.Fprintln(f, "func (*", t, ") UnmarshalJSON([]byte) error { return nil }")
	}
	return nil
}

// writeMain creates a .go file that launches the generator if 'go run'.
func (g *Generator) writeMain() (path string, err error) {
	f, err := ioutil.TempFile("", "easyjson")
	if err != nil {
		return "", err
	}

	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f, "import (")
	fmt.Fprintln(f, `  "fmt"`)
	fmt.Fprintln(f, `  "os"`)
	fmt.Fprintln(f)
	fmt.Fprintf(f, "  %q\n", genPackage)
	fmt.Fprintln(f)
	fmt.Fprintf(f, "  pkg %q\n", g.PkgPath)
	fmt.Fprintln(f, ")")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "func main() {")
	fmt.Fprintln(f, "  g := gen.NewGenerator()")
	fmt.Fprintf(f, "  g.SetPkg(%q, %q)\n", g.PkgName, g.PkgPath)
	if g.BuildTags != "" {
		fmt.Fprintf(f, "  g.SetBuildTags(%q)\n", g.BuildTags)
	}
	if g.SnakeCase {
		fmt.Fprintln(f, "  g.UseSnakeCase()")
	}
	if g.OmitEmpty {
		fmt.Fprintln(f, "  g.OmitEmpty()")
	}
	if g.NoStdMarshalers {
		fmt.Fprintln(f, "  g.NoStdMarshalers()")
	}
	for _, v := range g.Types {
		fmt.Fprintln(f, "  g.Add(pkg."+v+"{})")
	}

	fmt.Fprintln(f, "  if err := g.Run(os.Stdout); err != nil {")
	fmt.Fprintln(f, "    fmt.Fprintln(os.Stderr, err)")
	fmt.Fprintln(f, "    os.Exit(1)")
	fmt.Fprintln(f, "  }")
	fmt.Fprintln(f, "}")

	p := f.Name()
	os.Rename(p, p+".go")
	return p + ".go", f.Close()
}

func (g *Generator) Run() error {
	if err := g.writeStub(); err != nil {
		return err
	}
	if g.StubsOnly {
		return nil
	}

	path, err := g.writeMain()
	if err != nil {
		return err
	}
	if !g.LeaveTemps {
		defer os.Remove(path)
	}

	f, err := ioutil.TempFile("", "easyjson-out")
	if err != nil {
		return err
	}
	if !g.LeaveTemps {
		defer os.Remove(f.Name()) // will not remove after rename
	}

	cmd := exec.Command("go", "run", path)
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return err
	}

	return os.Rename(f.Name(), g.OutName)
}
