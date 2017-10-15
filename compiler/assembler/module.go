package assembler

import (
	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/jsast"
)

// Module is a go package that
// has many declarations
type Module struct {
	path         string
	declarations []types.Declaration
	imports      map[string]string
	exports      []string
}

// Module creates a Module
func newModule(path string) *Module {
	return &Module{path: path}
}

// Declaration add a declaration
func (m *Module) Declaration(d types.Declaration) {
	// TODO: imports

	if d.Exported() {
		m.exports = append(m.exports, d.Name())
	}

	m.declarations = append(m.declarations, d)
}

// Translate fn
func (m *Module) Translate() (jsast.IStatement, error) {
	var decls []interface{}

	// skip over any packages that don't have any exports
	if len(m.exports) == 0 {
		return nil, nil
	}

	// // create imports linking of the pkgs between packages
	// // e.g. var runner = pkg["github.com/gorunner/runner"]
	// var imports []jsast.VariableDeclarator
	// for name, dep := range pkg.Dependencies {
	// 	// don't include a dependencies that doesn't have
	// 	// any exports as that package will be eliminated
	// 	// from the build
	// 	if len(dep.Exports) == 0 {
	// 		continue
	// 	}

	// 	imports = append(imports, jsast.CreateVariableDeclarator(
	// 		jsast.CreateIdentifier(name),
	// 		jsast.CreateMemberExpression(
	// 			jsast.CreateIdentifier("pkg"),
	// 			jsast.CreateString(dep.Path),
	// 			true,
	// 		),
	// 	))
	// }
	// if len(imports) > 0 {
	// 	decls = append(decls, jsast.CreateVariableDeclaration(
	// 		"var",
	// 		imports...,
	// 	))
	// }

	// create the declaration body
	for _, decl := range m.declarations {
		stmt, err := decl.Translate()
		if err != nil {
			return nil, err
		}
		decls = append(decls, stmt)
	}

	// create a return statement for the exported fields
	// e.g. return { $export: $export }
	var props []jsast.Property
	for _, name := range m.exports {
		props = append(props, jsast.CreateProperty(
			jsast.CreateIdentifier(name),
			jsast.CreateIdentifier(name),
			"init",
		))
	}
	decls = append(decls, jsast.CreateReturnStatement(
		jsast.CreateObjectExpression(props),
	))

	// create: pkg["$dep"] = (function () { $yourPkg })()
	return jsast.CreateExpressionStatement(
		jsast.CreateAssignmentExpression(
			jsast.CreateMemberExpression(
				jsast.CreateIdentifier("pkg"),
				jsast.CreateString(m.path),
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
	), nil
}
