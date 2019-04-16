// Package compiler is the entrypoint to Joy's compiler
// It follows the steps outlined in: https://mat.tm/joy/#how
package compiler

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/bindata"
	"github.com/matthewmueller/joy/internal/compiler/def"
	"github.com/matthewmueller/joy/internal/compiler/defs"
	"github.com/matthewmueller/joy/internal/compiler/index"
	"github.com/matthewmueller/joy/internal/compiler/indexer"
	"github.com/matthewmueller/joy/internal/compiler/script"
	"github.com/matthewmueller/joy/internal/compiler/translator"
	"github.com/matthewmueller/joy/internal/jsast"
	"github.com/pkg/errors"
	govalidator "gopkg.in/asaskevich/govalidator.v4"

	"github.com/matthewmueller/joy/internal/compiler/graph"
	"github.com/matthewmueller/joy/internal/compiler/loader"
)

// module struct
type module struct {
	path           string
	defs           []def.Definition
	exports        []string
	imports        map[string]string
	isImplicitFile bool
}

// file struct
type file struct {
	path    string
	source  string
	modules []*module
}

// Config for the compiler
type Config struct {
	// build a development or production bundle
	Development bool

	// joy's root path
	JoyPath string `valid:"required"`

	// packages to compile
	Packages []string
}

// Compile our Go application into a Javascript application
func Compile(cfg *Config) (scripts []*script.Script, err error) {
	if valid, err := govalidator.ValidateStruct(cfg); !valid || err != nil {
		return scripts, errors.Wrapf(err, "invalid loader config")
	}

	// Setup the local state we need for the runtime, stdlib
	// and macro dependencies
	if err := Setup(cfg); err != nil {
		return scripts, errors.Wrapf(err, "compiler setup error")
	}

	// Parse: loads, indexes and graphs the program
	index, graph, err := Parse(cfg)
	if err != nil {
		return scripts, errors.Wrap(err, "compiler parse error")
	}

	// Assemble: translates and generates the javascript
	scripts, err = Assemble(index, graph)
	if err != nil {
		return scripts, errors.Wrap(err, "compiler assemble error")
	}

	return scripts, nil
}

// Setup the compiler's stdlib, runtime and macro source code
// we're going to keep this stupid simple right now, and just
// overwrite everytime. Overhead right now is ~3ms on the compiler.
//
// TODO: optimize
func Setup(cfg *Config) (err error) {
	for _, asset := range bindata.AssetNames() {
		fullpath := path.Join(cfg.JoyPath, asset)
		if err := os.MkdirAll(path.Dir(fullpath), 0755); err != nil {
			return errors.Wrapf(err, "error mkdirall")
		}

		buf, err := bindata.Asset(asset)
		if err != nil {
			return errors.Wrapf(err, "unable to get asset")
		}

		if err := ioutil.WriteFile(fullpath, buf, 0644); err != nil {
			return errors.Wrapf(err, "error writefile")
		}
	}
	return nil
}

// Parse loads, indexes and graphs the program
//
// This is breaking down our program into all the pieces
// we need to build it back up as a Javascript program.
//
// Eventually, we'll be able use this stage separately
// from assembling to incrementally compile packages
func Parse(cfg *Config) (idx *index.Index, g *graph.Graph, err error) {
	// defer log.Trace("parse").Stop(&err)

	program, err := loader.Load(&loader.Config{
		JoyPath:  cfg.JoyPath,
		Packages: cfg.Packages,
	})
	if err != nil {
		return idx, g, errors.Wrapf(err, "load error")
	}

	idx, err = indexer.New(program)
	if err != nil {
		return idx, g, errors.Wrapf(err, "indexing error")
	}

	// init() functions are not yet supported, so let's disable
	if len(idx.Inits()) > 0 {
		return idx, g, fmt.Errorf("init() functions are not supported yet")
	}

	// initial packages
	visited := map[string]bool{}
	queue := []def.Definition{}
	mains := idx.Mains()
	g = graph.New()

	// queue the initial packages
	queue = append(queue, mains...)

	// mark mains as visited immediately
	for _, q := range queue {
		visited[q.ID()] = true
	}

	// build our dependencies graph
	// hotidx := idx.Copy()
	for len(queue) > 0 {
		// dequeue
		d := queue[0]
		queue = queue[1:]

		// mark as visited
		visited[d.ID()] = true

		log.Debugf("visiting=%s kind=%s", d.ID(), d.Kind())

		// get the dependencies
		deps, err := d.Dependencies()
		if err != nil {
			return idx, g, errors.Wrap(err, "error getting dependencies")
		}

		// add the dependencies to the graph
		for _, dep := range deps {
			log.Debugf("edge: %s->%s", d.Path(), dep.Path())
			g.Edge(d, dep)

			// queue up any dependency that hasn't
			// already been visited
			if !visited[dep.ID()] {
				queue = append(queue, dep)
			}
		}
	}

	return idx, g, nil
}

// Assemble the code
func Assemble(idx *index.Index, g *graph.Graph) (scripts []*script.Script, err error) {
	// defer log.Trace("assemble").Stop(&err)

	tr := translator.New(idx)

	// assemble into scripts
	var files []*file
	for _, main := range idx.Mains() {
		log.Debugf("main: %s", main.ID())

		sorted := g.Toposort(main)
		for _, path := range sorted {
			log.Debugf("path=%s", path)
		}

		// paths back to definitions
		defs := []def.Definition{}
		for _, id := range sorted {
			def := idx.Get(id)
			if def == nil {
				return scripts, fmt.Errorf("compiler/assemble: couldn't find def: %s", id)
			}
			defs = append(defs, def)
		}

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
				// included because of an explicit include
				// in the code itself:
				//
				// e.g. preact := js.File("./preact.js") translates
				// to   var preact = pkg["./preact.js"]
				if len(moduleMap[path].exports) == 0 && !moduleMap[path].isImplicitFile {
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
			jsast.CreateExpressionStatement(
				jsast.CreateCallExpression(
					jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
						jsast.CreateFunctionBody(body...),
					),
					[]jsast.IExpression{},
				),
			),
		)

		code, err := jsast.Assemble(prog)
		if err != nil {
			return scripts, errors.Wrapf(err, "unable to assemble AST for %s", file.path)
		}

		scripts = append(scripts, script.New(
			file.path, // TODO
			file.path,
			code,
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

// rearrange fixes the order for some common issues the
// graph wouldn't know about. Right now this:
// 1. moves methods below constructors
// TODO: 2. move prototypes below constructor functions
func rearrange(defs []def.Definition) []def.Definition {
	return defs
}

// group declarations into modules
func group(definitions []def.Definition) (modules []*module, err error) {
	moduleMap := map[string]*module{}
	order := []string{}
	for _, def := range definitions {
		from := def.Path()

		if moduleMap[from] == nil {
			moduleMap[from] = &module{path: from}
			order = append(order, from)
		}

		if file, ok := def.(defs.Filer); ok {
			moduleMap[from].isImplicitFile = file.Implicit()
		}

		moduleMap[from].defs = append(
			moduleMap[from].defs,
			def,
		)

		log.Debugf("%s: exported=%t omitted=%t", def.ID(), def.Exported(), def.Omitted())
		if def.Exported() && !def.Omitted() {
			moduleMap[from].exports = append(
				moduleMap[from].exports,
				def.Name(),
			)
		}

		moduleMap[from].imports = map[string]string{}
		for alias, path := range def.Imports() {
			// only include modules whose path is in our
			// topologically sorted map
			if moduleMap[path] != nil {
				log.Debugf("adding path %s => %s", alias, path)
				moduleMap[from].imports[alias] = path
			}
		}
	}

	for _, from := range order {
		modules = append(modules, moduleMap[from])
	}

	return modules, nil
}
