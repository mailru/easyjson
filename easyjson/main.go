package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mailru/easyjson/bootstrap"
	"github.com/mailru/easyjson/parser"
)

// Flags
var (
	buildTags                = flag.String("build_tags", "", "build tags to add to generated file")
	genBuildFlags            = flag.String("gen_build_flags", "", "build flags when running the generator while bootstrapping")
	snakeCase                = flag.Bool("snake_case", false, "use snake_case names instead of CamelCase by default")
	lowerCamelCase           = flag.Bool("lower_camel_case", false, "use lowerCamelCase names instead of CamelCase by default")
	noStdMarshalers          = flag.Bool("no_std_marshalers", false, "don't generate MarshalJSON/UnmarshalJSON funcs")
	omitEmpty                = flag.Bool("omit_empty", false, "omit empty fields by default")
	allStructs               = flag.Bool("all", false, "generate marshaler/unmarshalers for all structs in a file")
	simpleBytes              = flag.Bool("byte", false, "use simple bytes instead of Base64Bytes for slice of bytes")
	leaveTemps               = flag.Bool("leave_temps", false, "do not delete temporary files")
	stubs                    = flag.Bool("stubs", false, "only generate stubs for marshaler/unmarshaler funcs")
	noformat                 = flag.Bool("noformat", false, "do not run 'gofmt -w' on output file")
	specifiedName            = flag.String("output_filename", "", "specify the filename of the output")
	processPkg               = flag.Bool("pkg", false, "process the whole package instead of just the given file")
	disallowUnknownFields    = flag.Bool("disallow_unknown_fields", false, "return error if any unknown field in json appeared")
	skipMemberNameUnescaping = flag.Bool("disable_members_unescape", false, "don't perform unescaping of member names to improve performance")
)

// generate executes code generation.
func generate(filename string) error {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	parser := parser.Parser{AllStructs: *allStructs}
	if err := parser.Parse(filename, fileInfo.IsDir()); err != nil {
		return fmt.Errorf("error parsing %v: %v", filename, err)
	}

	outName := determineOutputFileName(filename, parser)

	generator := bootstrap.Generator{
		BuildTags:                trimStringFlag(*buildTags),
		GenBuildFlags:            trimStringFlag(*genBuildFlags),
		PkgPath:                  parser.PkgPath,
		PkgName:                  parser.PkgName,
		Types:                    parser.StructNames,
		SnakeCase:                *snakeCase,
		LowerCamelCase:           *lowerCamelCase,
		NoStdMarshalers:          *noStdMarshalers,
		DisallowUnknownFields:    *disallowUnknownFields,
		SkipMemberNameUnescaping: *skipMemberNameUnescaping,
		OmitEmpty:                *omitEmpty,
		LeaveTemps:               *leaveTemps,
		OutName:                  outName,
		StubsOnly:                *stubs,
		NoFormat:                 *noformat,
		SimpleBytes:              *simpleBytes,
	}

	if err := generator.Run(); err != nil {
		return fmt.Errorf("bootstrap failed: %v", err)
	}
	return nil
}

// determineOutputFileName determines the output filename based on the input filename.
func determineOutputFileName(filename string, parser parser.Parser) string {
	if fileInfo, err := os.Stat(filename); err == nil && fileInfo.IsDir() {
		return filepath.Join(filename, parser.PkgName+"_easyjson.go")
	}

	if !strings.HasSuffix(filename, ".go") {
		return strings.TrimSuffix(filename, filepath.Ext(filename)) + "_easyjson.go"
	}

	return strings.TrimSuffix(filename, ".go") + "_easyjson.go"
}

// trimStringFlag removes whitespace from a string flag.
func trimStringFlag(flagValue string) string {
	return strings.TrimSpace(flagValue)
}

func main() {
	flag.Parse()

	files := flag.Args()

	if *processPkg {
		gofile := os.Getenv("GOFILE")
		if gofile != "" {
			files = []string{filepath.Dir(gofile)}
		}
	}

	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, filename := range files {
		if err := generate(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}