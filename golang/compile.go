package golang

import (
	"go/ast"
	"go/build"
	"go/parser"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/matthewmueller/golly/golang/indexer"
	"github.com/matthewmueller/golly/golang/inspector"
	"github.com/matthewmueller/golly/golang/translator"
	"github.com/matthewmueller/golly/jsast"
	"github.com/matthewmueller/golly/stdlib"
	"github.com/matthewmueller/golly/types"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// Declaration struct
type Declaration struct {
	node ast.Decl
	info *loader.PackageInfo
}

// Package struct
type Package struct {
	info  *loader.PackageInfo
	nodes []ast.Decl
}

// Compiler struct
type Compiler struct {
	// index        map[string]ast.Node
	// declarations []Declaration
}

// New Compiler
func New() *Compiler {
	return &Compiler{}
}

// Compile a series of packages
//
// packages should be fullpaths to the files or packages
func (c *Compiler) Compile(packages ...string) (files []*types.File, scripts []*types.Script, err error) {
	var conf loader.Config

	// TODO: will this work everytime?
	gosrc := path.Join(os.Getenv("GOPATH"), "src")

	// add all the packages as imports
	for _, pkgpath := range packages {
		if filepath.HasPrefix(pkgpath, gosrc) {
			rel, e := filepath.Rel(gosrc, pkgpath)
			if e != nil {
				return nil, nil, e
			}
			pkgpath = rel
		}

		// support both files and packages
		if path.Ext(pkgpath) == ".go" {
			conf.CreateFromFilenames(path.Dir(pkgpath), pkgpath)
		} else {
			conf.Import(pkgpath)
		}
	}

	// import the runtime by default
	// TODO: there's gotta be a better way to do this
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, nil, errors.New("unable to get the filepath")
	}
	runtimePkg, err := filepath.Rel(path.Join(os.Getenv("GOPATH"), "src"), path.Join(path.Dir(path.Dir(file)), "runtime"))
	if err != nil {
		return nil, nil, err
	}
	conf.Import(runtimePkg)

	// add comments to the AST
	conf.ParserMode = parser.ParseComments

	// replace stdlib packages with our own
	conf.FindPackage = func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
		alias, e := stdlib.Supports(importPath)
		if e != nil {
			return nil, e
		}

		// not part of the stdlib
		if alias == "" {
			return ctxt.Import(importPath, fromDir, mode)
		}

		// use the alias
		return ctxt.Import(alias, fromDir, mode)
	}

	// load all the packages
	program, err := conf.Load()
	if err != nil {
		return files, nil, errors.Wrap(err, "unable to load the go package")
	}

	// index the entire program
	index, err := indexer.New(program)
	if err != nil {
		return nil, nil, err
	}

	scripts, err = inspector.Inspect(program, index)
	if err != nil {
		return nil, nil, err
	}

	// TODO: split into functions
	for _, script := range scripts {
		var allpkgs []interface{}

		// append any raw file
		for _, rawfile := range script.RawFiles {
			// create: pkg["$dep"] = (function () { $yourPkg })()
			wrap := jsast.CreateExpressionStatement(
				jsast.CreateAssignmentExpression(
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(rawfile.Name),
						true,
					),
					jsast.AssignmentOperator("="),
					jsast.CreateCallExpression(
						jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
							jsast.CreateFunctionBody(
								jsast.CreateRaw(rawfile.Source),
							),
						),
						[]jsast.IExpression{},
					),
				),
			)

			allpkgs = append(allpkgs, wrap)
		}

		for _, pkg := range script.Packages {
			var decls []interface{}
			var exports []string

			// skip over any packages that don't have any exports
			if len(pkg.Exports) == 0 {
				continue
			}

			// create imports linking of the pkgs between packages
			// e.g. var runner = pkg["github.com/gorunner/runner"]
			var imports []jsast.VariableDeclarator
			for name, dep := range pkg.Dependencies {
				// don't include a dependencies that doesn't have
				// any exports as that package will be eliminated
				// from the build
				if len(dep.Exports) == 0 {
					continue
				}

				imports = append(imports, jsast.CreateVariableDeclarator(
					jsast.CreateIdentifier(name),
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(dep.Path),
						true,
					),
				))
			}
			if len(imports) > 0 {
				decls = append(decls, jsast.CreateVariableDeclaration(
					"var",
					imports...,
				))
			}

			// create the declaration body
			for _, decl := range pkg.Declarations {
				ast, err := translator.Translate(index, program.Package(pkg.Path), decl)
				if err != nil {
					return nil, nil, err
				}
				decls = append(decls, ast)

				// append if it's exported and not global
				// TODO: should decl.Exported encompass the global option?
				if decl.Exported {
					exports = append(exports, decl.Name)
				}
			}

			// create a return statement for the exported fields
			// e.g. return { $export: $export }
			var props []jsast.Property
			for _, decl := range pkg.Exports {
				props = append(props, jsast.CreateProperty(
					jsast.CreateIdentifier(decl.Name),
					jsast.CreateIdentifier(decl.Name),
					"init",
				))
			}
			decls = append(decls, jsast.CreateReturnStatement(
				jsast.CreateObjectExpression(props),
			))

			// create: pkg["$dep"] = (function () { $yourPkg })()
			wrap := jsast.CreateExpressionStatement(
				jsast.CreateAssignmentExpression(
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(pkg.Path),
						true,
					),
					jsast.AssignmentOperator("="),
					jsast.CreateCallExpression(
						jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
							jsast.CreateFunctionBody(decls...),
						),
						[]jsast.IExpression{},
					),
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
		callmain := jsast.CreateReturnStatement(
			jsast.CreateCallExpression(
				jsast.CreateMemberExpression(
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(script.Name),
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
			jsast.CreateEmptyStatement(),
			jsast.CreateExpressionStatement(
				jsast.CreateCallExpression(
					jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
						jsast.CreateFunctionBody(filebody...),
					),
					[]jsast.IExpression{},
				),
			),
		)

		files = append(files, &types.File{
			Name:   script.Name,
			Source: prog.String(),
		})
	}

	return files, scripts, nil
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
