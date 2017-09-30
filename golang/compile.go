package golang

import (
	"go/build"
	"go/parser"
	"sort"

	"github.com/matthewmueller/golly/js"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// CompilePackage compiles a package by it's path
func CompilePackage(packagePath string) (string, error) {
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
		return "", errors.Wrap(err, "unable to load the go package")
	}

	// get a deterministic toposort of the imports
	// for _, dep := range deps {
	// 	log.Infof("dep: %s", dep)
	// }

	// get the topological sort of the dependencies
	var deps []string
	for _, o := range order {
		lvl := imports[o]
		sort.Strings(lvl)
		deps = append(deps, lvl...)
	}
	deps = reverse(deps)

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
			return "", errors.Wrapf(err, "error translating %s into a JS package", info.Pkg.Name())
		}

		// create: pkg["$dep"] = (function () { $yourPkg })()
		wrap := js.CreateExpressionStatement(
			js.CreateAssignmentExpression(
				js.CreateMemberExpression(
					js.CreateIdentifier("pkg"),
					js.CreateString(dep),
					true,
				),
				js.AssignmentOperator("="),
				pkgfn,
			),
		)

		spkgs = append(spkgs, wrap)
	}

	// create: `var pkg = {}`
	init := js.CreateVariableDeclaration(
		"var",
		js.CreateVariableDeclarator(
			js.CreateIdentifier("pkg"),
			js.CreateObjectExpression([]js.Property{}),
		),
	)

	// create `pkg[$main].main()`
	callmain := js.CreateExpressionStatement(
		js.CreateCallExpression(
			js.CreateMemberExpression(
				js.CreateMemberExpression(
					js.CreateIdentifier("pkg"),
					js.CreateString(packagePath),
					true,
				),
				js.CreateIdentifier("main"),
				false,
			),
			[]js.IExpression{},
		),
	)

	// create the program body
	var pbody []interface{}
	pbody = append(pbody, init)
	pbody = append(pbody, spkgs...)
	pbody = append(pbody, callmain)

	// put everything together into a program
	prog := js.CreateProgram(
		js.CreateExpressionStatement(
			js.CreateCallExpression(
				js.CreateFunctionExpression(nil, []js.IPattern{},
					js.CreateFunctionBody(pbody...),
				),
				[]js.IExpression{},
			),
		),
	)

	// assemble that program
	return prog.String(), nil
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
