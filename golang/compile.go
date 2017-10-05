package golang

import (
	"go/ast"
	"go/build"
	"go/parser"
	"sort"

	"github.com/matthewmueller/golly/golang/inspector"
	"github.com/matthewmueller/golly/golang/translator"
	"github.com/matthewmueller/golly/jsast"
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
func (c *Compiler) Compile(packages ...string) (files []*types.File, err error) {
	var conf loader.Config

	// add all the packages as imports
	for _, pkg := range packages {
		conf.Import(pkg)
	}

	// add comments to the AST
	conf.ParserMode = parser.ParseComments

	// mock out "js"
	conf.FindPackage = func(ctx *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
		if importPath == "github.com/matthewmueller/golly/js" {
			return ctx.Import(importPath+"mock", fromDir, mode)
		}
		return ctx.Import(importPath, fromDir, mode)
	}

	// load all the packages
	program, err := conf.Load()
	if err != nil {
		return files, errors.Wrap(err, "unable to load the go package")
	}

	scripts, err := inspector.Inspect(program)
	if err != nil {
		return nil, err
	}

	// TODO: split into functions
	for _, script := range scripts {

		var allpkgs []interface{}
		for _, pkg := range script.Packages {
			var decls []interface{}
			var exports []string

			// create imports linking of the pkgs between packages
			// e.g. var runner = pkg["github.com/gorunner/runner"]
			var imports []jsast.VariableDeclarator
			for name, from := range pkg.Dependencies {
				imports = append(imports, jsast.CreateVariableDeclarator(
					jsast.CreateIdentifier(name),
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(from),
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
				ast, err := translator.Translate(decl)
				if err != nil {
					return nil, err
				}
				decls = append(decls, ast)

				if decl.Exported {
					exports = append(exports, decl.Name)
				}
			}

			// create a return statement for the exported fields
			// e.g. return { $export: $export }
			var exported interface{}
			if len(exports) > 0 {
				var props []jsast.Property
				for _, export := range exports {
					props = append(props, jsast.CreateProperty(
						jsast.CreateIdentifier(export),
						jsast.CreateIdentifier(export),
						"init",
					))
				}
				exported = jsast.CreateReturnStatement(
					jsast.CreateObjectExpression(props),
				)

				decls = append(decls, exported)
			}

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
		callmain := jsast.CreateExpressionStatement(
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
