package assembler

import (
	"github.com/apex/log"
	"github.com/matthewmueller/golly/compiler/graph"
	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/jsast"
)

type file struct {
	name   string
	path   string
	source string
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Path() string {
	return f.path
}

func (f *file) Source() string {
	return f.source
}

// Assemble the graph into files
func Assemble(g *graph.Graph) (files []types.File, err error) {
	defer log.Trace("assemble").Stop(&err)

	// get the roots of the graph, these will be our files
	for _, decl := range g.Roots() {
		program, err := translate(decl.Path(), group(g.Sort(decl)))
		if err != nil {
			return files, err
		}

		files = append(files, &file{
			name:   decl.Name(),
			path:   decl.Path(),
			source: program.String(),
		})
	}

	return files, nil
}

// group declarations into modules
func group(decls []types.Declaration) []*Module {
	moduleMap := map[string]*Module{}
	var order []string

	for _, decl := range decls {
		from := decl.Path()
		if moduleMap[from] == nil {
			moduleMap[from] = newModule(from)
			order = append(order, from)
		}
		moduleMap[from].Declaration(decl)
	}

	var modules []*Module
	for _, from := range order {
		modules = append(modules, moduleMap[from])
	}

	return modules
}

func translate(rootPath string, modules []*Module) (p jsast.Program, e error) {
	var body []interface{}

	// `var pkg = {}`: initialize pkg variable
	body = append(body, jsast.CreateVariableDeclaration(
		"var",
		jsast.CreateVariableDeclarator(
			jsast.CreateIdentifier("pkg"),
			jsast.CreateObjectExpression([]jsast.Property{}),
		),
	))

	// translate each of the modules
	for _, module := range modules {
		m, e := module.Translate()
		if e != nil {
			return p, e
		} else if m == nil {
			continue
		}
		body = append(body, m)
	}

	// `pkg[$main].main()`: create final call expression
	body = append(body, jsast.CreateReturnStatement(
		jsast.CreateCallExpression(
			jsast.CreateMemberExpression(
				jsast.CreateMemberExpression(
					jsast.CreateIdentifier("pkg"),
					jsast.CreateString(rootPath),
					true,
				),
				jsast.CreateIdentifier("main"),
				false,
			),
			[]jsast.IExpression{},
		),
	))

	// ;(function () { ...body... })()
	// wrap it up in a program
	return jsast.CreateProgram(
		jsast.CreateEmptyStatement(),
		jsast.CreateExpressionStatement(
			jsast.CreateCallExpression(
				jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
					jsast.CreateFunctionBody(body...),
				),
				[]jsast.IExpression{},
			),
		),
	), nil
}

// func createScript(name string, pkgs []*types.Package) (program jsast.Program, err error) {

// 	// `var pkg = {}`: initialize pkg variable
// 	body = append(body, jsast.CreateVariableDeclaration(
// 		"var",
// 		jsast.CreateVariableDeclarator(
// 			jsast.CreateIdentifier("pkg"),
// 			jsast.CreateObjectExpression([]jsast.Property{}),
// 		),
// 	))

// 	for _, pkg := range pkgs {
// 		p, e := createPackage(pkg)
// 		if e != nil {
// 			return program, e
// 		}
// 		body = append(body, p)
// 	}

// 	// `pkg[$main].main()`: create final call expression
// 	body = append(body, jsast.CreateReturnStatement(
// 		jsast.CreateCallExpression(
// 			jsast.CreateMemberExpression(
// 				jsast.CreateMemberExpression(
// 					jsast.CreateIdentifier("pkg"),
// 					jsast.CreateString(name),
// 					true,
// 				),
// 				jsast.CreateIdentifier("main"),
// 				false,
// 			),
// 			[]jsast.IExpression{},
// 		),
// 	))

// 	// ;(function () { ...body... })()
// 	// wrap it up in a program
// 	return jsast.CreateProgram(
// 		jsast.CreateEmptyStatement(),
// 		jsast.CreateExpressionStatement(
// 			jsast.CreateCallExpression(
// 				jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
// 					jsast.CreateFunctionBody(body...),
// 				),
// 				[]jsast.IExpression{},
// 			),
// 		),
// 	), nil

// 	// imports := map[string]string{}
// 	// exports := []string{}

// 	// for _, decl := range decls {
// 	// 	stmt, err := decl.Translate()
// 	// 	if err != nil {
// 	// 		return program, err
// 	// 	}
// 	// 	body = append(body, stmt)

// 	// 	// add to exported
// 	// 	if decl.Exported() {
// 	// 		exports = append(exports, decl.Name())
// 	// 	}
// 	// }

// 	// // don't export a package that has no exports
// 	// if len(exports) == 0 {
// 	// 	return program, nil
// 	// }

// 	// log.Infof("decls %+v", decls)
// 	// return program, nil
// }

// func createPackage(pkg *types.Package) (jsast.IStatement, error) {
// 	var body []interface{}

// 	// create imports linking of the pkgs between packages
// 	// e.g. var runner = pkg["github.com/gorunner/runner"]
// 	// var imports []jsast.VariableDeclarator
// 	// for name, dep := range pkg.Dependencies {
// 	// 	// don't include a dependencies that doesn't have
// 	// 	// any exports as that package will be eliminated
// 	// 	// from the build
// 	// 	if len(dep.Exports) == 0 {
// 	// 		continue
// 	// 	}

// 	// 	imports = append(imports, jsast.CreateVariableDeclarator(
// 	// 		jsast.CreateIdentifier(name),
// 	// 		jsast.CreateMemberExpression(
// 	// 			jsast.CreateIdentifier("pkg"),
// 	// 			jsast.CreateString(dep.Path),
// 	// 			true,
// 	// 		),
// 	// 	))
// 	// }
// 	// if len(imports) > 0 {
// 	// 	body = append(body, jsast.CreateVariableDeclaration(
// 	// 		"var",
// 	// 		imports...,
// 	// 	))
// 	// }

// 	return nil, nil
// }
