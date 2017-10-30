package golang

import (
	"sort"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/defs"
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
	buckets := map[string][]def.Definition{}
	visited := map[string]bool{}
	queue := []def.Definition{}
	mains := index.Mains()
	g := graph.New()

	// queue the initial packages
	queue = append(queue, mains...)
	// mark mains as visited immediately
	for _, main := range mains {
		visited[main.ID()] = true
	}

	// build our dependencies graph
	for len(queue) > 0 {
		// dequeue
		d := queue[0]
		queue = queue[1:]

		// path => definition buckets
		// NOTE: may contain duplicate definitions
		if buckets[d.Path()] == nil {
			buckets[d.Path()] = []def.Definition{}
		}
		buckets[d.Path()] = append(buckets[d.Path()], d)

		// get the dependencies
		edges, err := d.Dependencies()
		if err != nil {
			return scripts, errors.Wrap(err, "error getting dependencies")
		}

		// add the dependencies to the graph
		for _, edge := range edges {
			g.Edge(d.Path(), edge.Definition().Path(), edge.Type())
		}

		// add any dependency that hasn't
		// already been visited
		for _, edge := range edges {
			if !visited[edge.Definition().ID()] {
				queue = append(queue, edge.Definition())
				visited[d.ID()] = true
			}
		}
	}
	tr := translator.New(index)

	// assemble into scripts
	var files []*file
	for _, main := range mains {
		log.Debugf("main: %s", main.Path())

		sorted := g.Toposort(main.Path())

		// paths back to definitions
		defs := []def.Definition{}
		for _, path := range sorted {
			// Sort to have a consistent definition order (alphabetical)
			// TODO: investigate if we can maintain the order
			// the code was written in. It seems difficult
			// because we don't have easy access to raw source files,
			// just package paths, which may put definitions across
			// files in any order.
			ds := sortByID(buckets[path])
			defs = append(defs, ds...)
		}
		// uniquify all the duplicates per main file
		defs = unique(defs)

		for _, d := range defs {
			log.Debugf("sorted=%s", d.ID())
		}

		// prune the definitions for any obvious
		// errors that the graph wouldn't catch
		pruned, e := prune(defs)
		for _, d := range pruned {
			log.Debugf("pruned=%s", d.ID())
		}

		// group into modules
		// TODO: can we skip some steps here?
		// seems like we're going paths => defs => paths
		modules, e := group(pruned)
		if e != nil {
			return scripts, e
		}

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

		// build a module map to lookup
		// imported modules quickly
		moduleMap := map[string]*module{}
		for _, module := range file.modules {
			moduleMap[module.path] = module
		}

		for _, module := range file.modules {
			var modbody []interface{}

			// create imports linking of the pkgs between packages
			// e.g. var runner = pkg["github.com/gorunner/runner"]
			imports := map[string]jsast.VariableDeclarator{}
			order := []string{}
			for alias, path := range module.imports {
				// don't implicitly include imports that don't
				// have any exports. Note that jsfiles don't
				// have any dependencies but they will still be
				// included because of an explicit include:
				// e.g. js.RawFile => pkg[]
				if len(moduleMap[path].exports) == 0 {
					continue
				}

				imports[path] = jsast.CreateVariableDeclarator(
					jsast.CreateIdentifier(alias),
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("pkg"),
						jsast.CreateString(path),
						true,
					),
				)
				order = append(order, path)
			}

			// sort so we're deterministic
			sort.Strings(order)

			for _, path := range order {
				modbody = append(modbody, jsast.CreateVariableDeclaration(
					"var",
					imports[path],
				))
			}

			// create the definition body
			var defbody []interface{}
			for _, def := range module.defs {
				if def.Omitted() {
					continue
				}

				ast, err := tr.Translate(def)
				if err != nil {
					return scripts, err
				} else if ast == nil {
					return scripts, errors.New("translate shouldn't return nil")
				}
				defbody = append(defbody, ast)
			}
			if len(defbody) == 0 {
				continue
			}

			// handle raw files differently
			if len(module.defs) == 1 &&
				module.defs[0].Kind() == "FILE" {
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
								jsast.CreateFunctionBody(defbody...),
							),
							[]jsast.IExpression{},
						),
					),
				))
				continue
			}

			// add the body to the module body
			modbody = append(modbody, defbody...)

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

// Prunes the definitions. Right now this means:
// - remove any method that doesn't have a receiver present
func prune(inp []def.Definition) (out []def.Definition, err error) {
	defmap := map[string]def.Definition{}
	for _, def := range inp {
		defmap[def.ID()] = def
	}

	for _, def := range inp {
		// ignore any method that doesn't have a receiver present
		if m, ok := def.(defs.Methoder); ok {
			if defmap[m.Recv().ID()] == nil {
				continue
			}
		}
		out = append(out, def)
	}

	return out, nil
}

// group declarations into modules
func group(defs []def.Definition) (modules []*module, err error) {
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

		// log.Debugf("%s: exported=%t omitted=%t", def.ID(), def.Exported(), def.Omitted())
		if def.Exported() && !def.Omitted() {
			moduleMap[from].exports = append(
				moduleMap[from].exports,
				def.Name(),
			)
		}

		moduleMap[from].imports = map[string]string{}
		for alias, path := range def.Imports() {
			// only include modules whos path is in our
			// topologically sorted map
			if moduleMap[path] != nil {
				moduleMap[from].imports[alias] = path
			}
		}
	}

	for _, from := range order {
		modules = append(modules, moduleMap[from])
	}

	return modules, nil
}

func unique(defs []def.Definition) []def.Definition {
	u := make([]def.Definition, 0, len(defs))
	seen := make(map[string]bool)

	for _, def := range defs {
		if _, ok := seen[def.ID()]; !ok {
			seen[def.ID()] = true
			u = append(u, def)
		}
	}

	return u
}

func sortByID(defs []def.Definition) (sorted []def.Definition) {
	defmap := map[string]def.Definition{}
	order := []string{}

	for _, def := range defs {
		id := def.ID()
		defmap[def.ID()] = def
		order = append(order, id)
	}

	sort.Strings(order)

	for _, o := range order {
		sorted = append(sorted, defmap[o])
	}

	return sorted
}
