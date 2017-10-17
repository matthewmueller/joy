package golang

import (
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/script"
	"github.com/matthewmueller/golly/golang/translator"
	"github.com/matthewmueller/golly/jsast"
	"github.com/pkg/errors"

	"github.com/matthewmueller/golly/golang/db"
	"github.com/matthewmueller/golly/golang/graph"
	"github.com/matthewmueller/golly/golang/loader"
)

// Declaration struct
// type Declaration struct {
// 	node ast.Decl
// 	info *loader.PackageInfo
// }

// module struct
type module struct {
	path    string
	defs    []def.Definition
	exports []string
	imports map[string]string
}

// file struct
type file struct {
	path    string
	source  string
	modules []*module
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
func (c *Compiler) Compile(packages ...string) (scripts []*script.Script, err error) {
	program, err := loader.Load(packages...)
	if err != nil {
		return scripts, err
	}

	index, err := db.New(program)
	if err != nil {
		return scripts, nil
	}

	// initial packages
	visited := map[string]bool{}
	queue := []def.Definition{}
	mains := index.Mains()
	g := graph.New()

	// queue the initial packages
	queue = append(queue, mains...)

	// build our dependencies graph
	for len(queue) > 0 {
		// dequeue
		def := queue[0]
		queue = queue[1:]

		// get the dependencies
		deps, err := def.Dependencies()
		if err != nil {
			return scripts, errors.Wrap(err, "error getting dependencies")
		}

		// add the dependencies to the graph
		g.AddDependency(def, deps...)

		// add any dependency that hasn't
		// already been visited
		for _, dep := range deps {
			if !visited[dep.ID()] {
				queue = append(queue, dep)
				visited[dep.ID()] = true
			}
		}
	}

	tr := translator.New(index)

	// assemble into scripts
	var files []*file
	for _, main := range mains {
		modules := group(g.Sort(main))
		files = append(files, &file{
			path:    main.Path(),
			modules: modules,
		})
	}

	for _, file := range files {
		var body []interface{}

		// create initial variable declaration: `var pkg = {}`
		body = append(body, jsast.CreateVariableDeclaration(
			"var",
			jsast.CreateVariableDeclarator(
				jsast.CreateIdentifier("pkg"),
				jsast.CreateObjectExpression([]jsast.Property{}),
			),
		))

		for _, module := range file.modules {
			var modbody []interface{}

			if len(module.exports) == 0 {
				continue
			}

			// TODO: imports

			// create the definition body
			for _, def := range module.defs {
				ast, err := tr.Translate(def)
				if err != nil {
					return scripts, err
				}
				modbody = append(modbody, ast)
			}

			// create a return statement for the exported fields
			// e.g. return { $export: $export }
			var props []jsast.Property
			for _, export := range module.exports {
				props = append(props, jsast.CreateProperty(
					jsast.CreateIdentifier(export),
					jsast.CreateIdentifier(export),
					"init",
				))
			}
			modbody = append(modbody, jsast.CreateReturnStatement(
				jsast.CreateObjectExpression(props),
			))

			// create: pkg["$dep"] = (function () { $yourPkg })()
			body = append(body, jsast.CreateExpressionStatement(
				jsast.CreateAssignmentExpression(
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(module.path),
						true,
					),
					jsast.AssignmentOperator("="),
					jsast.CreateCallExpression(
						jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
							jsast.CreateFunctionBody(modbody...),
						),
						[]jsast.IExpression{},
					),
				),
			))
		}

		// create final call expression: `pkg[$main].main()`
		body = append(body, jsast.CreateReturnStatement(
			jsast.CreateCallExpression(
				jsast.CreateMemberExpression(
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(file.path),
						true,
					),
					jsast.CreateIdentifier("main"),
					false,
				),
				[]jsast.IExpression{},
			),
		))

		// put everything together into a JS program
		prog := jsast.CreateProgram(
			jsast.CreateEmptyStatement(),
			jsast.CreateExpressionStatement(
				jsast.CreateCallExpression(
					jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
						jsast.CreateFunctionBody(body...),
					),
					[]jsast.IExpression{},
				),
			),
		)

		scripts = append(scripts, script.New(
			file.path, // TODO
			file.path,
			prog.String(),
		))
	}

	return scripts, nil
}

// group declarations into modules
func group(defs []def.Definition) []*module {
	moduleMap := map[string]*module{}
	order := []string{}
	for _, def := range defs {
		from := def.Path()

		if moduleMap[from] == nil {
			moduleMap[from] = &module{path: from}
			order = append(order, from)
		}

		moduleMap[from].defs = append(
			moduleMap[from].defs,
			def,
		)

		if def.Exported() {
			moduleMap[from].exports = append(
				moduleMap[from].exports,
				def.Name(),
			)
		}
	}

	var modules []*module
	for _, from := range order {
		modules = append(modules, moduleMap[from])
	}
	return modules
}
