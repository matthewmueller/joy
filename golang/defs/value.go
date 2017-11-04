package defs

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"

	"golang.org/x/tools/go/loader"
)

// Valuer interface
type Valuer interface {
	def.Definition
	Node() *ast.ValueSpec
	Rewrite(caller string, arguments ...string) (string, error)
}

var _ Valuer = (*values)(nil)

type values struct {
	exported  bool
	path      string
	name      string
	id        string
	index     *index.Index
	gen       *ast.GenDecl
	node      *ast.ValueSpec
	kind      types.Type
	processed bool
	rewrite   *rewrite
	edges     []def.Edge
	imports   map[string]string
	async     bool
	omit      bool
	tag       *structtag.Tag
}

// Value fn
func Value(index *index.Index, info *loader.PackageInfo, gn *ast.GenDecl, n *ast.ValueSpec) (def.Definition, error) {
	packagePath := info.Pkg.Path()
	names := []string{}
	exported := false

	for _, ident := range n.Names {
		obj := info.ObjectOf(ident)
		if obj.Exported() {
			exported = true
		}
		names = append(names, ident.Name)
	}

	name := strings.Join(names, ",")
	idParts := []string{packagePath, name}
	id := strings.Join(idParts, " ")

	return &values{
		exported: exported,
		path:     packagePath,
		name:     name,
		id:       id,
		index:    index,
		node:     n,
		gen:      gn,
		imports:  map[string]string{},
	}, nil
}

func (d *values) process() (err error) {
	state, e := process(d.index, d, d.gen)
	if e != nil {
		return e
	}

	// copy state into function
	d.processed = true
	d.async = state.async
	d.edges = state.edges.Edges()
	d.imports = state.imports
	d.rewrite = state.rewrite
	d.omit = state.omit
	d.tag = state.tag

	return nil
}

func (d *values) ID() string {
	return d.id
}

func (d *values) Name() string {
	if d.tag != nil {
		return d.tag.Name
	}
	return d.name
}

func (d *values) Kind() string {
	return "VALUE"
}

func (d *values) Path() string {
	return d.path
}

func (d *values) Dependencies() (edges []def.Edge, err error) {
	if d.processed {
		return d.edges, nil
	}
	e := d.process()
	if e != nil {
		return edges, e
	}

	return d.edges, nil
}

func (d *values) Exported() bool {
	return d.exported
}

func (d *values) Omitted() bool {
	if d.tag != nil {
		return d.tag.HasOption("omit")
	}
	return d.omit
}

func (d *values) Node() *ast.ValueSpec {
	return d.node
}

func (d *values) Type() types.Type {
	return d.kind
}

func (d *values) Imports() map[string]string {
	return d.index.GetImports(d.path)
}

func (d *values) FromRuntime() bool {
	return false
}

// Rewrite fn
func (d *values) Rewrite(caller string, arguments ...string) (string, error) {
	if !d.processed {
		if e := d.process(); e != nil {
			return "", e
		}
	}

	if d.rewrite == nil {
		return "", nil
	}

	return d.rewrite.Rewrite(caller, arguments)
}
