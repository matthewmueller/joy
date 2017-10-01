package golang

import (
	"go/build"
	"go/parser"
	"sort"

	"github.com/apex/log"
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

// Compiler struct
type Compiler struct {
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

	// load all the packages
	pkgs, err := conf.Load()
	if err != nil {
		return files, errors.Wrap(err, "unable to load the go package")
	}

	// get the topological sort of the dependencies
	var deps []string
	for _, o := range order {
		lvl := imports[o]
		sort.Strings(lvl)
		deps = append(deps, lvl...)
	}
	deps = reverse(deps)

	// get a deterministic toposort of the imports
	for _, dep := range deps {
		log.Infof("dep: %s", dep)
	}

	// translate each file to their Javascript counterpart
	// this is done in topological order to ensure that
	// we have all the type information by the time we
	// need to use it and also so JS initializes in the
	// right order
	var spkgs []interface{}
	for _, dep := range deps {
		info := pkgs.Package(dep)

		// translate the package returning a self-executing
		// function that returns the public (capitalized)
		// values.
		//
		// Example:
		//
		//  (function () {
		//    ...body...
		//    return { main: main, Another: Another }
		//  })()
		//
		pkgfn, err := translatePackage(info)
		if err != nil {
			return files, errors.Wrapf(err, "error translating %s into a JS package", info.Pkg.Name())
		}

		// create: pkg["$dep"] = (function () { $yourPkg })()
		wrap := jsast.CreateExpressionStatement(
			jsast.CreateAssignmentExpression(
				jsast.CreateMemberExpression(
					jsast.CreateIdentifier("pkg"),
					jsast.CreateString(dep),
					true,
				),
				jsast.AssignmentOperator("="),
				pkgfn,
			),
		)

		spkgs = append(spkgs, wrap)
	}

	// create: `var pkg = {}`
	init := jsast.CreateVariableDeclaration(
		"var",
		jsast.CreateVariableDeclarator(
			jsast.CreateIdentifier("pkg"),
			jsast.CreateObjectExpression([]jsast.Property{}),
		),
	)

	// create `pkg[$main].main()`
	callmain := jsast.CreateExpressionStatement(
		jsast.CreateCallExpression(
			jsast.CreateMemberExpression(
				jsast.CreateMemberExpression(
					jsast.CreateIdentifier("pkg"),
					jsast.CreateString(packages[0]),
					true,
				),
				jsast.CreateIdentifier("main"),
				false,
			),
			[]jsast.IExpression{},
		),
	)

	// create the program body
	var pbody []interface{}
	pbody = append(pbody, init)
	pbody = append(pbody, spkgs...)
	pbody = append(pbody, callmain)

	// put everything together into a program
	prog := jsast.CreateProgram(
		jsast.CreateExpressionStatement(
			jsast.CreateCallExpression(
				jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
					jsast.CreateFunctionBody(pbody...),
				),
				[]jsast.IExpression{},
			),
		),
	)

	// assemble that program
	files = append(files, &file{
		name:   packages[0],
		source: prog.String(),
	})

	return files, nil
}

func (c *Compiler) compilePackage(info *loader.PackageInfo) error {
	return nil
}

// 	// translate each file to their Javascript counterpart
// 	// this is done in topological order to ensure that
// 	// we have all the type information by the time we
// 	// need to use it and also so JS initializes in the
// 	// right order
// 	var spkgs []interface{}
// 	for _, dep := range deps {
// 		info := pkgs.Package(dep)

// 		// translate the package returning a self-executing
// 		// function that returns the public (capitalized)
// 		// values.
// 		//
// 		// Example:
// 		//
// 		//  (function () {
// 		//    ...body...
// 		//    return { main: main, Another: Another }
// 		//  })()
// 		//
// 		pkgfn, err := translatePackage(info)
// 		if err != nil {
// 			return files, errors.Wrapf(err, "error translating %s into a JS package", info.Pkg.Name())
// 		}

// 		// create: pkg["$dep"] = (function () { $yourPkg })()
// 		wrap := jsast.CreateExpressionStatement(
// 			jsast.CreateAssignmentExpression(
// 				jsast.CreateMemberExpression(
// 					jsast.CreateIdentifier("pkg"),
// 					jsast.CreateString(dep),
// 					true,
// 				),
// 				jsast.AssignmentOperator("="),
// 				pkgfn,
// 			),
// 		)

// 		spkgs = append(spkgs, wrap)
// 	}

// 	// create: `var pkg = {}`
// 	init := jsast.CreateVariableDeclaration(
// 		"var",
// 		jsast.CreateVariableDeclarator(
// 			jsast.CreateIdentifier("pkg"),
// 			jsast.CreateObjectExpression([]jsast.Property{}),
// 		),
// 	)

// 	// create `pkg[$main].main()`
// 	callmain := jsast.CreateExpressionStatement(
// 		jsast.CreateCallExpression(
// 			jsast.CreateMemberExpression(
// 				jsast.CreateMemberExpression(
// 					jsast.CreateIdentifier("pkg"),
// 					jsast.CreateString(packagePath),
// 					true,
// 				),
// 				jsast.CreateIdentifier("main"),
// 				false,
// 			),
// 			[]jsast.IExpression{},
// 		),
// 	)

// 	// create the program body
// 	var pbody []interface{}
// 	pbody = append(pbody, init)
// 	pbody = append(pbody, spkgs...)
// 	pbody = append(pbody, callmain)

// 	// put everything together into a program
// 	prog := jsast.CreateProgram(
// 		jsast.CreateExpressionStatement(
// 			jsast.CreateCallExpression(
// 				jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
// 					jsast.CreateFunctionBody(pbody...),
// 				),
// 				[]jsast.IExpression{},
// 			),
// 		),
// 	)

// 	// assemble that program
// 	return prog.String(), nil

// 	return []File{}, nil
// }

// CompilePackage compiles a package by it's path
// func CompilePackage(packagePath string) (string, error) {
// 	var conf loader.Config
// 	conf.Import(packagePath)

// 	// TODO: clean up, this is ugly
// 	order := []string{}
// 	imports := map[string][]string{}

// 	// load each of the imports
// 	// NOTE: we can override import here
// 	conf.FindPackage = func(ctx *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
// 		if imports[fromDir] == nil {
// 			order = append(order, fromDir)
// 		}
// 		imports[fromDir] = append(imports[fromDir], importPath)
// 		return ctx.Import(importPath, fromDir, mode)
// 	}

// }

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
