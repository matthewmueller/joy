package defs

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/loader"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"
)

// Typer interface
type Typer interface {
	def.Definition
	Gen() *ast.GenDecl
	Node() *ast.TypeSpec
	Transform(n ast.Expr) (ast.Expr, error)
	Methods() []def.Definition
	Implements() ([]def.Definition, error)
}

var _ Typer = (*typedef)(nil)

type typedef struct {
	exported  bool
	path      string
	name      string
	id        string
	index     *index.Index
	gen       *ast.GenDecl
	node      *ast.TypeSpec
	kind      types.Type
	tag       *structtag.Tag
	fields    []*field
	deps      []def.Definition
	processed bool
	omit      bool
}

// Type fn
func Type(index *index.Index, info *loader.PackageInfo, gn *ast.GenDecl, n *ast.TypeSpec) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}
	id := strings.Join(idParts, " ")

	// log.Infof("info.TypeOf(n.Name) %s", info.TypeOf(n.Name))

	return &typedef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     name,
		id:       id,
		index:    index,
		node:     n,
		gen:      gn,
		kind:     info.TypeOf(n.Name),
	}, nil
}

func (d *typedef) process() (err error) {
	// state, e := process(d.index, d, d.gen)
	// if e != nil {
	// 	return e
	// }

	// // copy state into function
	// d.processed = true
	// d.deps = state.deps
	// d.omit = state.omit
	// d.tag = state.tag
	// d.fields = state.fields

	return nil
}

func (d *typedef) Dependencies() (deps []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	err = d.process()
	if err != nil {
		return deps, err
	}
	return d.deps, nil
}

func (d *typedef) ID() string {
	return d.id
}

func (d *typedef) Name() string {
	if d.tag != nil {
		return d.tag.Name
	}
	return d.name
}

func (d *typedef) Kind() string {
	return "TYPE"
}

func (d *typedef) OriginalName() string {
	return d.name
}

func (d *typedef) Path() string {
	return d.path
}

func (d *typedef) Exported() bool {
	return d.exported
}

func (d *typedef) Omitted() bool {
	return true
}

func (d *typedef) Gen() *ast.GenDecl {
	return d.gen
}

func (d *typedef) Node() *ast.TypeSpec {
	return d.node
}

func (d *typedef) Type() types.Type {
	return d.kind
}

func (d *typedef) Imports() map[string]string {
	return d.index.GetImports(d.path)
}

func (d *typedef) FromRuntime() bool {
	return false
}

// Methods functions
func (d *typedef) Methods() (methods []def.Definition) {
	for _, def := range d.index.All() {
		method, ok := def.(Methoder)
		if !ok {
			continue
		}
		if d.ID() == method.Recv().ID() {
			methods = append(methods, method)
		}
	}
	return methods
}

// Implements function
func (d *typedef) Implements() (defs []def.Definition, err error) {
	return defs, nil
}

func (d *typedef) Transform(n ast.Expr) (ast.Expr, error) {
	if len(d.Methods()) > 0 {
		return nil, fmt.Errorf("defs/Type: doesn't support attaching methods to the '%s' type yet", d.name)
	}

	switch t := n.(type) {
	case *ast.CallExpr:
		return d.callExpr(t)
	case *ast.CompositeLit:
		return d.compositeLit(t)
	case *ast.Ident:
		return d.node.Type, nil
	default:
		return nil, fmt.Errorf("defs/Type/Transform: unhandled expression %T", n)
	}
}

func (d *typedef) callExpr(n *ast.CallExpr) (ast.Expr, error) {
	if len(n.Args) != 1 {
		return nil, fmt.Errorf("defs/Type: expected n.Args in typedef.CallExpr(n) to have 1 argument but received %d args", len(n.Args))
	}

	switch d.node.Type.(type) {
	case *ast.Ident, *ast.SelectorExpr:
		return n.Args[0], nil
	default:
		return nil, fmt.Errorf("defs/Type/CallExpr: unhandled call expression %T", d.node.Type)
	}
}

func (d *typedef) compositeLit(n *ast.CompositeLit) (ast.Expr, error) {
	return d.node.Type, nil
}
