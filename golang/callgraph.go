package golang

import (
	"bytes"
	"fmt"
	"go/build"
	"go/parser"
	"go/types"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/rta"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// https://github.com/chenzhongtao/go-test/blob/28a756e5cbf0e2338c55a6e6d69886b404554639/src/golang.org/x/tools/go/callgraph/rta/rta_test.go

// CallGraph fn
func CallGraph(packagePath string) error {
	var conf loader.Config
	conf.Import(packagePath)

	// TODO: clean up, this is ugly
	order := []string{}
	imports := map[string][]string{}

	// load each of the imports
	// NOTE: we can override import here
	conf.FindPackage = func(ctx *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
		if imports[fromDir] == nil {
			order = append(order, fromDir)
		}
		imports[fromDir] = append(imports[fromDir], importPath)
		return ctx.Import(importPath, fromDir, mode)
	}

	// add comments to the AST
	conf.ParserMode = parser.ParseComments

	// load the package
	pkgs, err := conf.Load()
	if err != nil {
		return errors.Wrap(err, "unable to load the go package")
	}

	prog := ssautil.CreateProgram(pkgs, 0)
	for _, pkg := range prog.AllPackages() {
		pkg.Build()
		main := pkg.Func("main")
		if main != nil {
			res := rta.Analyze([]*ssa.Function{
				main,
			}, true)
			println(printResult(res, pkg.Pkg))
		}
	}

	return nil
}

func printResult(res *rta.Result, from *types.Package) string {
	var buf bytes.Buffer

	writeSorted := func(ss []string) {
		sort.Strings(ss)
		for _, s := range ss {
			fmt.Fprintf(&buf, "  %s\n", s)
		}
	}

	buf.WriteString("Dynamic calls\n")
	var edges []string
	callgraph.GraphVisitEdges(res.CallGraph, func(e *callgraph.Edge) error {
		if strings.Contains(e.Description(), "dynamic") {
			edges = append(edges, fmt.Sprintf("%s --> %s",
				e.Caller.Func.RelString(from),
				e.Callee.Func.RelString(from)))
		}
		return nil
	})
	writeSorted(edges)

	buf.WriteString("Reachable functions\n")
	var reachable []string
	for f := range res.Reachable {
		reachable = append(reachable, f.RelString(from))
	}
	writeSorted(reachable)

	buf.WriteString("Reflect types\n")
	var rtypes []string
	res.RuntimeTypes.Iterate(func(key types.Type, value interface{}) {
		if value == false { // accessible to reflection
			rtypes = append(rtypes, types.TypeString(key, types.RelativeTo(from)))
		}
	})
	writeSorted(rtypes)

	return strings.TrimSpace(buf.String())
}
