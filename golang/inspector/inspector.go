package inspector

import (
	"errors"
	"fmt"
	"go/ast"
	gotypes "go/types"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/apex/log"
	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/types"
	"golang.org/x/tools/go/loader"
)

// Inspect fn
func Inspect(program *loader.Program) (scripts []*types.Script, err error) {
	// map[packagePath]map[variableName]dependencyPath
	imports := map[string]map[string]string{}
	rawfileMap := map[string][]*types.File{}

	// human-friendly pointers to the runtime declarations
	runtimeDecls := map[string]*types.Declaration{}

	// get the runtime path
	runtimePath, err := getRuntimePath()
	if err != nil {
		return nil, err
	}

	// build a map of all the declarations in all the packages
	// this will be used as a lookup to map object declarations
	// to actual AST nodes. object.String() is a unique identifier
	// that points to a declaration in a go package (e.g. main())
	//
	// map[object.String()] => ast
	//
	// TODO: split into separate functions
	// and perhaps a separate package
	// (e.g. indexer)
	index := map[string]*types.Declaration{}
	for _, info := range program.AllPackages {
		packagePath := info.Pkg.Path()
		for _, file := range info.Files {
			for _, decl := range file.Decls {
				switch t := decl.(type) {
				case *ast.FuncDecl:
					obj := info.ObjectOf(t.Name)
					name := t.Name.Name
					id := obj.String()

					// if it's a method don't export,
					// if it's the main() function
					// export either way
					exported := obj.Exported()
					if t.Recv != nil {
						exported = false
					} else if name == "main" {
						exported = true
					}

					index[id] = &types.Declaration{
						Exported: exported,
						From:     packagePath,
						Name:     name,
						ID:       id,
						Node:     decl,
					}

					// point human-friendly names to the declaration
					if runtimePath == packagePath {
						runtimeDecls[name] = index[id]
					}

				case *ast.GenDecl:
					for _, spec := range t.Specs {
						switch y := spec.(type) {
						case *ast.ValueSpec:
							for _, name := range y.Names {
								obj := info.ObjectOf(name)
								// packagePath := obj.Pkg().Path()
								id := obj.String()
								index[id] = &types.Declaration{
									Exported: obj.Exported(),
									From:     packagePath,
									Name:     name.Name,
									ID:       id,
									Node:     decl,
								}
							}
						case *ast.TypeSpec:
							obj := info.ObjectOf(y.Name)
							// packagePath := obj.Pkg().Path()
							id := obj.String()
							index[id] = &types.Declaration{
								Exported: obj.Exported(),
								From:     packagePath,
								Name:     y.Name.Name,
								ID:       id,
								Node:     decl,
							}
						case *ast.ImportSpec:
							if imports[packagePath] == nil {
								imports[packagePath] = map[string]string{}
							}

							// trim the "" of package path's
							depPath := strings.Trim(y.Path.Value, `"`)

							// TODO: can y.Path be nil?
							var name string
							if y.Name != nil {
								name = y.Name.Name
							} else {
								name = path.Base(depPath)
							}

							imports[packagePath][name] = depPath
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
	remaining := []*types.Declaration{}
	mains := []*types.Declaration{}
	for _, info := range program.InitialPackages() {
		pkgPath := info.Pkg.Path()
		main := info.Pkg.Scope().Lookup("main")

		if main == nil {
			// ignore the runtime
			// TODO: more robust
			if path.Base(pkgPath) == "runtime" {
				continue
			}
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
		visited[declaration.ID] = true

		// get the package information
		info := program.Package(declaration.From)

		// walk through the AST of each declaration
		// the goal of this inspection is to:
		//
		// - build the dependency tree for dead-code elimination
		// - figure out which functions are async (from goroutines)
		// - does this declaration contain "js" tags
		var err error
		ast.Inspect(declaration.Node, func(node ast.Node) bool {
			switch n := node.(type) {

			// Get "js" comment tags from the top of functions
			case *ast.FuncDecl:
				tag, e := getJSTag(n.Doc)
				if e != nil {
					err = e
					return false
				} else if tag != nil {
					declaration.JSTag = tag
				}

			// case *ast.BlockStmt:

			// 	n.List

			// Get "js" comment tags from the top of types
			case *ast.TypeSpec:
				tag, e := getJSTag(n.Doc)
				if e != nil {
					err = e
					return false
				} else if tag != nil {
					declaration.JSTag = tag
				}

			// if we encounter and Go statements or channels,
			// we'll make the whole declaration asynchronous
			// TODO: goroutines on their own don't cause a function
			// to be async, it's the channels sending and receiving
			// Note: we'll also need to add Channel when make(chan)
			// is called.
			case *ast.GoStmt:
				if imports[declaration.From] == nil {
					imports[declaration.From] = map[string]string{}
				}
				imports[declaration.From]["runtime"] = runtimePath

				// if child is nil, it's a local variable or field
				channel := runtimeDecls["Channel"]
				if channel == nil {
					return true
				}

				send := runtimeDecls["Send"]
				if send == nil {
					return true
				}

				recv := runtimeDecls["Recv"]
				if recv == nil {
					return true
				}

				// add the dependency
				declaration.Dependencies = append(
					declaration.Dependencies,
					channel,
					send,
					recv,
				)

				// append to the crawl
				if !visited[channel.ID] {
					remaining = append(remaining, channel)
				}

				if !visited[send.ID] {
					remaining = append(remaining, send)
				}

				if !visited[recv.ID] {
					remaining = append(remaining, recv)
				}

				// ast, e := jsast.Parse(string(bindata.MustAsset("bindata/channel.js")))
				// if e != nil {
				// 	log.Infof("e %s", e)
				// 	err = e
				// 	return false
				// }

				// _ = ast
				// log.Infof("got ast!")
				// log.Infof("")
				// declaration.Includes = append(declaration.Includes, &types.File{
				// 	Name:   "channel.js",
				// 	Source: string(bindata.MustAsset("bindata/channel.js")),
				// })

				declaration.Async = true
			// dig into our identifiers to figure out where
			// they were defined and add those dependencies
			// to the crawler
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
					declaration.Dependencies = append(
						declaration.Dependencies,
						child,
					)

					// append to the crawl
					if !visited[child.ID] {
						remaining = append(remaining, child)
					}
				}

			// look for js.RawFile(filename)
			case *ast.CallExpr:
				file, e := checkJSRawFile(n, declaration.From)
				if e != nil {
					err = e
					return false
				}
				if file == nil {
					return true
				}

				// tack on the rawfile's map
				if rawfileMap[declaration.From] == nil {
					rawfileMap[declaration.From] = []*types.File{}
				}
				rawfileMap[declaration.From] = append(
					rawfileMap[declaration.From],
					file,
				)
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
		packages := groupByPackages(main)
		rawfiles := []*types.File{}

		// tack on the dependencies along with their alias names
		for _, pkg := range packages {
			pkg.Dependencies = map[string]string{}
			for alias, path := range imports[pkg.Path] {
				pkg.Dependencies[alias] = path
			}

			for _, file := range rawfileMap[pkg.Path] {
				rawfiles = append(rawfiles, file)
			}
		}

		scripts = append(scripts, &types.Script{
			Name:     main.From,
			Packages: packages,
			RawFiles: rawfiles,
		})
	}

	return scripts, nil
}

// groupByPackages takes a dependency tree of declarations
// and organizes them into a topologically order list of
// packages
func groupByPackages(d *types.Declaration) (pkgs []*types.Package) {
	pkgmap := map[string]*types.Package{}
	visited := map[string]bool{}
	order := &[]string{}

	// get the list of packages from first to last referenced package
	pkgs = recurseDeclaration(d, pkgmap, visited, order)

	// reverse the list for topological order
	return reverse(pkgs)
}

// recurse the declarations, building up an ordered
// list of packages that contain these declarations
func recurseDeclaration(
	d *types.Declaration,
	pkgmap map[string]*types.Package,
	visited map[string]bool,
	order *[]string,
) (pkgs []*types.Package) {
	if visited[d.ID] {
		return pkgs
	}
	visited[d.ID] = true

	// if we haven't seen this package yet,
	// then create it
	if pkgmap[d.From] == nil {
		*order = append(*order, d.From)
		pkgmap[d.From] = &types.Package{
			Path: d.From,
		}
	}
	// append the declaration
	pkgmap[d.From].Declarations = append(
		pkgmap[d.From].Declarations,
		d,
	)

	// loop over each of the dependencies
	for _, decl := range d.Dependencies {
		recurseDeclaration(decl, pkgmap, visited, order)
	}

	// arrange according to first seen
	for _, o := range *order {
		pkgs = append(pkgs, pkgmap[o])
	}

	return pkgs
}

func getDependency(obj gotypes.Object) string {
	if obj == nil {
		return ""
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return ""
	}

	switch t := obj.(type) {
	case *gotypes.Var:
		return t.String()
	case *gotypes.Func:
		return t.String()
	case *gotypes.Const:
		return t.String()
	case *gotypes.TypeName:
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

func reverse(s []*types.Package) []*types.Package {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// unique returns a unique set of packages
func unique(input []*types.Package) []*types.Package {
	u := make([]*types.Package, 0, len(input))
	m := make(map[string]bool)

	for _, pkg := range input {
		if _, ok := m[pkg.Path]; !ok {
			m[pkg.Path] = true
			u = append(u, pkg)
		}
	}

	return u
}

func getRuntimePath() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the filepath")
	}
	runtimePkg, err := filepath.Rel(path.Join(os.Getenv("GOPATH"), "src"), path.Join(path.Dir(path.Dir(path.Dir(file))), "runtime"))
	if err != nil {
		return "", err
	}

	return runtimePkg, nil
}

func checkJSRawFile(cx *ast.CallExpr, from string) (*types.File, error) {
	sel, ok := cx.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}

	x, ok := sel.X.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	if x.Name != "js" || sel.Sel.Name != "RawFile" {
		return nil, nil
	}

	if len(cx.Args) == 0 {
		return nil, nil
	}

	lit, ok := cx.Args[0].(*ast.BasicLit)
	if !ok {
		return nil, nil
	}

	filepath := lit.Value[1 : len(lit.Value)-1]
	importPath := path.Join(from, filepath)
	filepath = path.Join(os.Getenv("GOPATH"), "src", importPath)

	// modify the value to be the fullpath
	// TODO: this is super convenient, but
	// it may not be clear that this function
	// does make modifications to the AST
	lit.Value = `"` + importPath + `"`

	src, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return &types.File{
		Name:   importPath,
		Source: string(src),
	}, nil
}
