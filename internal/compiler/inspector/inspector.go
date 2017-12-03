package inspector

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	gotypes "go/types"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/fatih/structtag"
	"github.com/matthewmueller/joy/internal/compiler/indexer"
	"github.com/matthewmueller/joy/types"
	"golang.org/x/tools/go/loader"
)

// Inspect fn
func Inspect(program *loader.Program, index *indexer.Index) (scripts []*types.Script, err error) {
	// map[packagePath]map[variableName]dependencyPath
	compositeLits := []string{}
	maybeDependants := map[string]*types.Declaration{}
	imports := map[string]map[string]string{}
	rawfileMap := map[string][]*types.File{}

	// human-friendly pointers to the runtime declarations
	// runtimeDecls := map[string]*types.Declaration{}

	// get the runtime path
	runtimePath, err := getRuntimePath()
	if err != nil {
		return nil, err
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

		maindecl, err := index.ObjectOf(main)
		if err != nil {
			return nil, err
		}
		remaining = append(remaining, maindecl)
		mains = append(mains, maindecl)
	}

	// loop over each declaration we discover
	// until we've discovered all the declarations
	visited := map[string]bool{}
	for len(remaining) > 0 {
		declaration := remaining[0]
		remaining = remaining[1:]
		// TODO: this fixes the infinite loop,
		// but the loop should be solved down
		// inside the cases
		if visited[declaration.ID] {
			continue
		}
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
					if tag.HasOption("global") {
						declaration.Exported = false
					}
				}

			// Get "js" comment tags from the top of types
			case *ast.GenDecl:
				tag, e := getJSTag(n.Doc)
				if e != nil {
					err = e
					return false
				} else if tag != nil {
					declaration.JSTag = tag
					if tag.HasOption("global") {
						declaration.Exported = false
					}
				}

			case *ast.StructType:
				for _, field := range n.Fields.List {
					var jstag *structtag.Tag

					// parse the JS tag, if we have one
					if field.Tag != nil {
						value := field.Tag.Value
						if field.Tag.Kind == token.STRING {
							value = value[1 : len(value)-1]
						}

						tag, e := getJSTagFromString(value)
						if e != nil {
							err = e
							return false
						}
						jstag = tag
					}

					// no name
					// e.g. struct { *dep.Settings }
					if len(field.Names) == 0 {
						id, e := getIdentifier(field.Type)
						if e != nil {
							err = e
							return false
						}
						declaration.Fields = append(declaration.Fields, types.Field{
							Name: id.Name,
							Tag:  jstag,
						})
						continue
					}

					for _, id := range field.Names {
						declaration.Fields = append(declaration.Fields, types.Field{
							Name: id.Name,
							Tag:  jstag,
						})
					}
				}

			// // link each interface id with all method declarations
			// // whose receiver's implement the interface
			// case *ast.InterfaceType:

			// Include the Channel function when we find:
			// make(chan ..., capacity)
			case *ast.ChanType:
				if imports[declaration.From] == nil {
					imports[declaration.From] = map[string]string{}
				}
				imports[declaration.From]["runtime"] = runtimePath

				// if child is nil, it's a local variable or field
				channel := index.Runtime("Channel")
				if channel == nil {
					return true
				}

				// add the dependency
				declaration.Dependencies = append(
					declaration.Dependencies,
					channel,
				)

				// append to the crawl
				if !visited[channel.ID] {
					remaining = append(remaining, channel)
				}

				recv := index.Runtime("Recv")
				if recv == nil {
					return true
				}

				// add the dependency
				declaration.Dependencies = append(
					declaration.Dependencies,
					recv,
				)

				if !visited[recv.ID] {
					remaining = append(remaining, recv)
				}

				send := index.Runtime("Send")
				if send == nil {
					return true
				}

				// add the dependency
				declaration.Dependencies = append(
					declaration.Dependencies,
					send,
				)

				if !visited[send.ID] {
					remaining = append(remaining, send)
				}

				declaration.Async = true
			// Send
			case *ast.SendStmt:
				// if imports[declaration.From] == nil {
				// 	imports[declaration.From] = map[string]string{}
				// }
				// imports[declaration.From]["runtime"] = runtimePath

				// send := runtimeDecls["Send"]
				// if send == nil {
				// 	return true
				// }

				// // add the dependency
				// declaration.Dependencies = append(
				// 	declaration.Dependencies,
				// 	send,
				// )

				// if !visited[send.ID] {
				// 	remaining = append(remaining, send)
				// }

				// declaration.Async = true

			// Receive
			case *ast.UnaryExpr:
				// if n.Op != token.ARROW {
				// 	return true
				// }

				// if imports[declaration.From] == nil {
				// 	imports[declaration.From] = map[string]string{}
				// }
				// imports[declaration.From]["runtime"] = runtimePath

				// recv := runtimeDecls["Recv"]
				// if recv == nil {
				// 	return true
				// }

				// // add the dependency
				// declaration.Dependencies = append(
				// 	declaration.Dependencies,
				// 	recv,
				// )

				// if !visited[recv.ID] {
				// 	remaining = append(remaining, recv)
				// }

				// declaration.Async = true

			// dig into our identifiers to figure out where
			// they were defined and add those dependencies
			// to the crawler

			case *ast.Ident:
				// check if we depend on any other
				// declarations with this identifier
				if n.Name == "Render" {
					typ := index.TypeOf(info, n)
					log.Infof("%T", typ.Underlying())
					obj := info.ObjectOf(n)
					log.Infof("typ %s", typ)
					log.Infof("obj %s", obj)
					switch t := typ.(type) {
					case *gotypes.Interface:
						log.Infof("interface!")
					case *gotypes.Signature:
						// log.Infof("%T", t.Recv().Type().Underlying().String())
						switch t := t.Recv().Type().Underlying().(type) {
						case *gotypes.Interface:
							log.Infof("interface %s", index.ImplementedBy(t.String(), "Render"))
							// log.Infof("%T", t.Underlying())
						}
					default:
						log.Infof("type %T", t)
						// case *gotypes.Signature:
						// 	log.Infof(t.)
					}
					// log.Infof("typeof %T", typ)
					d, _ := index.DeclarationOf(info, n)
					c, _ := index.DeclarationOf(info, n)
					log.Infof("d %+v child %+v", d, c)
				}

				child, e := index.DeclarationOf(info, n)
				if e != nil {
					err = e
					return true
				} else if child == nil {
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

			case *ast.CompositeLit:
				// store composite literals since they will help
				// us determine if we should make certain interface
				// methods available in the source or not
				kind := info.TypeOf(n.Type)
				compositeLits = append(compositeLits, kind.String())

			case *ast.CallExpr:
				// look for js.RawFile(filename)
				file, e := checkJSRawFile(n, declaration.From)
				if e != nil {
					err = e
					return false
				} else if file != nil {
					// tack on the rawfile's map
					if rawfileMap[declaration.From] == nil {
						rawfileMap[declaration.From] = []*types.File{}
					}
					rawfileMap[declaration.From] = append(
						rawfileMap[declaration.From],
						file,
					)
					return true
				}

				// look for js.Rewrite(expression, variables...)
				expr, vars, e := checkJSRewrite(n)
				if e != nil {
					err = e
					return false
				} else if expr != "" {
					declaration.Rewrite = &types.Rewrite{
						Expression: expr,
						Variables:  vars,
					}
					declaration.Exported = false
					return true
				}

				switch t := n.Fun.(type) {
				case *ast.SelectorExpr:
					typeof := info.TypeOf(t.X)
					method := t.Sel.Name
					decls := index.ImplementedBy(typeof.String(), method)

					// maybe dependants if it's receivers are
					// initialized anywhere in the source code
					for _, decl := range decls {
						maybeDependants[decl.ID] = declaration
					}
					// case *ast.Ident:
					// 	log.Infof("fn: %s", t.Name)
					// default:
					// 	log.Infof("fn: %T", n.Fun)
				}
			}

			return true
		})
		if err != nil {
			return nil, err
		}
	}

	// map over the composite literals,
	// pull their methods out if they
	// have any and check those methods
	// against maybeDependants.
	for _, id := range compositeLits {
		methods := index.Methods(id)
		for _, method := range methods {
			if parent, isset := maybeDependants[method.ID]; isset {
				parent.Dependencies = append(
					parent.Dependencies,
					method,
				)
			}
		}
	}

	// organize the declarations in topological
	// order by package according to the initial
	// main()'s
	for _, main := range mains {
		packages, pkgmap := groupByPackages(main)
		rawfiles := []*types.File{}

		// tack on the dependencies along with their alias names
		for _, pkg := range packages {
			pkg.Dependencies = map[string]*types.Package{}

			// add the indexed imports
			for alias, path := range index.Imports(pkg.Path) {
				// check if we're actually using this package in
				// this main (it may be used in other mains)
				if pkgmap[path] == nil {
					continue
				}

				pkg.Dependencies[alias] = pkgmap[path]
			}
			// add any additional imports we add ourselves
			for alias, path := range imports[pkg.Path] {
				// check if we're actually using this package in
				// this main (it may be used in other mains)
				if pkgmap[path] == nil {
					continue
				}

				pkg.Dependencies[alias] = pkgmap[path]
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
func groupByPackages(d *types.Declaration) (pkgs []*types.Package, pkgmap map[string]*types.Package) {
	pkgmap = map[string]*types.Package{}
	visited := map[string]bool{}
	order := &[]string{}

	// get the list of packages from first to last referenced package
	pkgs = recurseDeclaration(d, pkgmap, visited, order)

	// reverse the list for topological order
	return reverse(pkgs), pkgmap
}

// recurse the declarations, building up an ordered
// list of packages that contain these declarations
// TODO: cleanup, this is really bad recursion
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

	// don't include joy/js in the build.
	// this package contains only compile-time stubs
	if strings.HasSuffix(d.From, "joy/js") {
		return pkgs
	}

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

	if d.Exported {
		// append the declaration
		pkgmap[d.From].Exports = append(
			pkgmap[d.From].Exports,
			d,
		)
	}

	// loop over each of the dependencies
	for _, decl := range d.Dependencies {
		recurseDeclaration(decl, pkgmap, visited, order)

		// async bubbles up, so if a child dependency is async,
		// then the parent dependency also becomes async
		if !d.Async && decl.Async {
			d.Async = decl.Async
		}
	}

	// arrange according to first seen
	for _, o := range *order {
		pkgs = append(pkgs, pkgmap[o])
	}

	return pkgs
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

		jstag, err := getJSTagFromString(comment.Text[2:])
		if err != nil {
			return nil, err
		}

		return jstag, nil
	}

	return nil, nil
}

func getJSTagFromString(tag string) (*structtag.Tag, error) {
	tags, err := structtag.Parse(tag)
	if err != nil {
		return nil, err
	}

	jstag, err := tags.Get("js")
	if err != nil && err.Error() != "tag does not exist" {
		return nil, err
	}

	return jstag, nil
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

func checkJSRewrite(cx *ast.CallExpr) (string, []string, error) {
	sel, ok := cx.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", nil, nil
	}

	x, ok := sel.X.(*ast.Ident)
	if !ok {
		return "", nil, nil
	}

	if x.Name != "js" || sel.Sel.Name != "Rewrite" {
		return "", nil, nil
	}

	if len(cx.Args) == 0 {
		return "", nil, nil
	}

	lit, ok := cx.Args[0].(*ast.BasicLit)
	if !ok {
		return "", nil, nil
	}

	var args []string
	for _, arg := range cx.Args[1:] {
		switch t := arg.(type) {
		case *ast.Ident:
			args = append(args, t.Name)
		default:
			return "", nil, fmt.Errorf("unhandled checkJSRewrite(): %T", arg)
		}
	}

	// trim off quotes if we've got a string
	expr := lit.Value
	if lit.Kind == token.STRING {
		x, e := strconv.Unquote(expr)
		if e != nil {
			return "", nil, e
		}
		expr = x
	}

	return expr, args, nil
}

func getIdentifier(n ast.Expr) (*ast.Ident, error) {
	switch t := n.(type) {
	case *ast.Ident:
		return t, nil
	case *ast.StarExpr:
		return getIdentifier(t.X)
	case *ast.SelectorExpr:
		return t.Sel, nil
	default:
		return nil, fmt.Errorf("unhandled getIdentifier: %T", n)
	}
}
