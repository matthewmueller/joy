package golang

import (
	"fmt"
	"go/ast"
	"go/parser"
	"reflect"
	"sort"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/translator"
	"github.com/matthewmueller/golly/jsast"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// File interface
type File interface {
	Name() string
	Source() string
}

type file struct {
	name   string
	source string
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Source() string {
	return f.source
}

// Declaration struct
type Declaration struct {
	node ast.Decl
	info *loader.PackageInfo
}

// Compiler struct
type Compiler struct {
	index        map[string]ast.Node
	declarations []Declaration
}

// New Compiler
func New() *Compiler {
	return &Compiler{}
}

// Compile a series of packages
func (c *Compiler) Compile(packages ...string) (files []File, err error) {
	var conf loader.Config

	// add all the packages as imports
	for _, pkg := range packages {
		conf.Import(pkg)
	}

	// add comments to the AST
	conf.ParserMode = parser.ParseComments

	// load all the packages
	pkgs, err := conf.Load()
	if err != nil {
		return files, errors.Wrap(err, "unable to load the go package")
	}

	// build a map of all the declarations in all the packages
	// this will be used as a lookup to map object declarations
	// to actual AST nodes
	//
	// map[object.String()] => ast
	index := map[string]*Declaration{}
	for _, info := range pkgs.AllPackages {
		for _, file := range info.Files {
			for _, decl := range file.Decls {
				switch t := decl.(type) {
				case *ast.FuncDecl:
					id := info.ObjectOf(t.Name).String()
					index[id] = &Declaration{decl, info}
				case *ast.GenDecl:
					for _, spec := range t.Specs {
						switch y := spec.(type) {
						case *ast.ValueSpec:
							for _, name := range y.Names {
								index[info.ObjectOf(name).String()] = &Declaration{decl, info}
							}
						case *ast.TypeSpec:
							index[info.ObjectOf(y.Name).String()] = &Declaration{decl, info}
						case *ast.ImportSpec:
						default:
							log.Fatalf("unknown type %s", reflect.TypeOf(y))
						}
					}
				default:
					log.Fatalf("unknown type %s", reflect.TypeOf(t))
				}
			}
		}
	}

	deps := map[string][]string{}
	remaining := []string{}
	initials := []string{}

	// get the mains for our initial packages
	for _, info := range pkgs.InitialPackages() {
		main := info.Pkg.Scope().Lookup("main")
		if main == nil {
			log.Fatalf("no main() found in: %s", info.Pkg.Path())
		}

		id := main.String()
		initials = append(initials, id)
		remaining = append(remaining, id)
	}

	// loop through the stack adding more
	// packages as we translate
	asts := map[string]jsast.INode{}
	for len(remaining) > 0 {
		id := remaining[0]
		remaining = remaining[1:]
		decl := index[id]
		if decl == nil {
			// log.Warnf("%s not found in index", id)
			continue
		}

		// translate our package from a Go AST to a Javascript AST
		result, err := translator.Translate(decl.info, decl.node)
		if err != nil {
			return nil, err
		}
		asts[id] = result.Node

		// build a dependency graph using the
		// dependencies that the translator
		// finds
		if deps[id] == nil {
			deps[id] = []string{}
		}
		deps[id] = append(deps[id], result.Dependencies...)

		// tack on this node's dependencies
		remaining = append(remaining, result.Dependencies...)
	}

	// now sort topologically, group the packages
	// and add to their appropriate files
	for _, initial := range initials {
		pkgnodes := map[string][]interface{}{}
		sorted := toposort(deps, initial)
		order := []string{}

		if index[initial] == nil {
			return nil, fmt.Errorf("%s expected to have an index entry", initial)
		}
		initialPath := index[initial].info.Pkg.Path()

		// groups by packages while maintaining topological order
		for _, dep := range sorted {
			idx := index[dep]
			ast := asts[dep]
			if idx == nil {
				// log.Warnf("nil idx: %s", dep)
				continue
			} else if ast == nil {
				// log.Warnf("nil ast: %s", dep)
				continue
			}

			// ensure proper ordering
			pkgpath := idx.info.Pkg.Path()
			if pkgnodes[pkgpath] == nil {
				pkgnodes[pkgpath] = []interface{}{}
				order = append(order, pkgpath)
			}

			// convert nodes to source code
			pkgnodes[pkgpath] = append(pkgnodes[pkgpath], ast)
		}

		// loop over each ordered package and wrap it.
		// all the packages will be the body of the file.
		var allpkgs []interface{}
		for _, pkgpath := range order {
			pkgbody := jsast.CreateCallExpression(
				jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
					jsast.CreateFunctionBody(pkgnodes[pkgpath]...),
				),
				[]jsast.IExpression{},
			)

			// create: pkg["$dep"] = (function () { $yourPkg })()
			wrap := jsast.CreateExpressionStatement(
				jsast.CreateAssignmentExpression(
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(pkgpath),
						true,
					),
					jsast.AssignmentOperator("="),
					pkgbody,
				),
			)

			allpkgs = append(allpkgs, wrap)
		}

		// create initial varaible declaration: `var pkg = {}`
		init := jsast.CreateVariableDeclaration(
			"var",
			jsast.CreateVariableDeclarator(
				jsast.CreateIdentifier("pkg"),
				jsast.CreateObjectExpression([]jsast.Property{}),
			),
		)

		// create final call expression: `pkg[$main].main()`
		callmain := jsast.CreateExpressionStatement(
			jsast.CreateCallExpression(
				jsast.CreateMemberExpression(
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(initialPath),
						true,
					),
					jsast.CreateIdentifier("main"),
					false,
				),
				[]jsast.IExpression{},
			),
		)

		// create the program body
		var filebody []interface{}
		filebody = append(filebody, init)
		filebody = append(filebody, allpkgs...)
		filebody = append(filebody, callmain)

		// put everything together into a JS program
		prog := jsast.CreateProgram(
			jsast.CreateExpressionStatement(
				jsast.CreateCallExpression(
					jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
						jsast.CreateFunctionBody(filebody...),
					),
					[]jsast.IExpression{},
				),
			),
		)

		files = append(files, &file{
			name:   initialPath,
			source: prog.String(),
		})
	}

	return files, nil
}

// simple topological sort
func toposort(m map[string][]string, initial string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := []string{initial}
	seen[initial] = true
	for _, key := range m[initial] {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return append(order, initial)
}
