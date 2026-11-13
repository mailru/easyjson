package bootstrap

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratorIgnoresDependencyInitOutput(t *testing.T) {
	t.Setenv("GOWORK", "off")
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	repo := filepath.Dir(wd)
	for _, test := range []struct {
		name     string
		noFormat bool
	}{
		{name: "formatted"},
		{name: "unformatted", noFormat: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			dir := filepath.Join(t.TempDir(), "module with spaces")
			files := map[string]string{
				"go.mod": fmt.Sprintf("module example.com/bootstrap\n\ngo 1.20\n\nrequire github.com/mailru/easyjson v0.0.0\n\nreplace github.com/mailru/easyjson => %q\n", filepath.ToSlash(repo)),
				"noisy/noisy.go": `package noisy

import "fmt"

func init() {
	fmt.Println("dependency initialization output")
}
`,
				"model/model.go": `package model

import _ "example.com/bootstrap/noisy"

type Record struct {
	Value string ` + "`json:\"value\"`" + `
}
`,
				"model/model_test.go": `package model

import (
	"encoding/json"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	want := Record{Value: "example"}
	data, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "{\"value\":\"example\"}" {
		t.Fatalf("unexpected JSON: %s", data)
	}
	var got Record
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("round trip = %+v, want %+v", got, want)
	}
}
`,
			}
			for name, content := range files {
				path := filepath.Join(dir, name)
				if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					t.Fatal(err)
				}
			}

			outName := filepath.Join(dir, "model", "model_easyjson.go")
			g := Generator{
				PkgPath:       "example.com/bootstrap/model",
				PkgName:       "model",
				Types:         []string{"Record"},
				OutName:       outName,
				GenBuildFlags: "-mod=mod",
				NoFormat:      test.noFormat,
			}
			if err := g.Run(); err != nil {
				t.Fatalf("generate with dependency init output: %v", err)
			}
			data, err := os.ReadFile(outName)
			if err != nil {
				t.Fatal(err)
			}
			if strings.Contains(string(data), "dependency initialization output") {
				t.Fatal("dependency output was included in generated code")
			}
			if _, err := parser.ParseFile(token.NewFileSet(), outName, data, parser.AllErrors); err != nil {
				t.Fatalf("parse generated code: %v", err)
			}
			cmd := exec.Command("go", "test", "./model")
			cmd.Dir = dir
			if output, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("test generated code: %v\n%s", err, output)
			}
		})
	}
}
