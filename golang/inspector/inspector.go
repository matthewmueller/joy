package inspector

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"github.com/apex/log"
	"github.com/fatih/structtag"
	"golang.org/x/tools/go/loader"
)

// Script struct
type Script struct {
	name     string
	packages []*Package
}

// Package struct
type Package struct {
	path         string
	declarations []*Declaration
}

// Declaration struct
type Declaration struct {
	exported bool
	from     string
	id       string
	node     ast.Decl

	// after analysis
	jstag        *structtag.Tag
	dependencies []*Declaration
	async        bool
}

// Inspect fn
func Inspect(program *loader.Program) (scripts []*Script, err error) {
	// build a map of all the declarations in all the packages
	// this will be used as a lookup to map object declarations
	// to actual AST nodes. object.String() is a unique identifier
	// that points to a declaration in a go package (e.g. main())
	//
	// map[object.String()] => ast
	index := map[string]*Declaration{}
	for _, info := range program.AllPackages {
		for _, file := range info.Files {
			for _, decl := range file.Decls {
				switch t := decl.(type) {
				case *ast.FuncDecl:
					obj := info.ObjectOf(t.Name)
					packagePath := obj.Pkg().Path()
					id := obj.String()
					index[id] = &Declaration{
						exported: obj.Exported(),
						from:     packagePath,
						id:       id,
						node:     decl,
					}
				case *ast.GenDecl:
					for _, spec := range t.Specs {
						switch y := spec.(type) {
						case *ast.ValueSpec:
							for _, name := range y.Names {
								obj := info.ObjectOf(name)
								packagePath := obj.Pkg().Path()
								id := obj.String()
								index[id] = &Declaration{
									exported: obj.Exported(),
									from:     packagePath,
									id:       id,
									node:     decl,
								}
							}
						case *ast.TypeSpec:
							obj := info.ObjectOf(y.Name)
							packagePath := obj.Pkg().Path()
							id := obj.String()
							index[id] = &Declaration{
								exported: obj.Exported(),
								from:     packagePath,
								id:       id,
								node:     decl,
							}
						case *ast.ImportSpec:
							// do nothing
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

	// get the mains for our initial packages
	remaining := []*Declaration{}
	mains := []*Declaration{}
	for _, info := range program.InitialPackages() {
		main := info.Pkg.Scope().Lookup("main")
		if main == nil {
			return nil, fmt.Errorf("no main() found in: %s", info.Pkg.Path())
		}

		maindecl := index[main.String()]
		remaining = append(remaining, maindecl)
		mains = append(mains, maindecl)
	}

	// loop over each declaration we discover
	// until we've discovered all the declarations
	visited := map[string]bool{}
	for len(remaining) > 0 {
		declaration := remaining[0]
		remaining = remaining[1:]
		visited[declaration.id] = true

		// get the package information
		info := program.Package(declaration.from)

		// walk through the AST of each declaration
		// the goal of this inspection is to:
		//
		// - build the dependency tree for dead-code elimination
		// - figure out which functions are async (from goroutines)
		// - does this declaration contain "js" tags
		var err error
		ast.Inspect(declaration.node, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.FuncDecl:
				tag, e := getJSTag(n.Doc)
				if e != nil {
					err = e
					return false
				} else if tag != nil {
					declaration.jstag = tag
				}
			case *ast.TypeSpec:
				tag, e := getJSTag(n.Doc)
				if e != nil {
					err = e
					return false
				} else if tag != nil {
					declaration.jstag = tag
				}
			case *ast.GoStmt:
				declaration.async = true
			case *ast.Ident:
				// check if we depend on any other
				// declarations with this identifier
				obj := info.ObjectOf(n)
				dep := getDependency(obj)

				if dep != "" {
					// if child is nil, it's a local variable or field
					child := index[dep]
					if child == nil {
						return true
					}

					// add the dependency
					declaration.dependencies = append(
						declaration.dependencies,
						child,
					)

					// append to crawl
					if !visited[child.id] {
						remaining = append(remaining, child)
					}
				}
			}

			return true
		})
		if err != nil {
			return nil, err
		}
	}

	// organize the declarations in topological
	// order by package according to the initial
	// main()'s
	for _, main := range mains {
		scripts = append(scripts, &Script{
			packages: groupByPackages(main),
			name:     main.from,
		})
	}

	return scripts, nil
}

func groupByPackages(d *Declaration) (pkgs []*Package) {
	pkgmap := map[string]*Package{}
	visited := map[string]bool{}
	order := &[]string{}

	pkgs = recursePackage(d, pkgmap, visited, order)

	// reverse the packages for topological order
	return reverse(pkgs)
}

func recursePackage(
	d *Declaration,
	pkgmap map[string]*Package,
	visited map[string]bool,
	order *[]string,
) (pkgs []*Package) {
	if visited[d.id] {
		return pkgs
	}
	visited[d.id] = true

	// if we haven't seen this package yet,
	// then create it
	if pkgmap[d.from] == nil {
		*order = append(*order, d.from)
		pkgmap[d.from] = &Package{
			path: d.from,
		}
	}
	// append the declaration
	pkgmap[d.from].declarations = append(
		pkgmap[d.from].declarations,
		d,
	)

	// loop over each of the dependencies
	for _, decl := range d.dependencies {
		recursePackage(decl, pkgmap, visited, order)
	}

	// arrange occurding to first seen
	for _, o := range *order {
		pkgs = append(pkgs, pkgmap[o])
	}

	return pkgs
}

func getDependency(obj types.Object) string {
	if obj == nil {
		return ""
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return ""
	}

	switch t := obj.(type) {
	case *types.Var:
		return t.String()
	case *types.Func:
		return t.String()
	case *types.Const:
		return t.String()
	}

	return ""
}

func getJSTag(n *ast.CommentGroup) (*structtag.Tag, error) {
	if n == nil {
		return nil, nil
	}

	comments := n.List
	for _, comment := range comments {
		if !strings.Contains(comment.Text, "js:") {
			continue
		}

		tags, err := structtag.Parse(comment.Text[2:])
		if err != nil {
			return nil, err
		}

		jstag, err := tags.Get("js")
		if err != nil {
			return nil, err
		}

		return jstag, nil
	}

	return nil, nil
}

func reverse(s []*Package) []*Package {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
