package defs

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/internal/compiler/def"
	"github.com/matthewmueller/golly/internal/compiler/index"
	"github.com/matthewmueller/golly/internal/compiler/util"

	"golang.org/x/tools/go/loader"
)

// Valuer interface
type Valuer interface {
	def.Definition

	Params() []string
	Rewrite() def.Rewrite
	IsVariadic() bool
	Node() *ast.ValueSpec
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
	rewrite   def.Rewrite
	deps      []def.Definition
	imports   map[string]string
	async     bool
	omit      bool
	tag       util.JSTag
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
	d.deps = state.deps
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
	if d.tag.Rename != "" {
		return d.tag.Rename
	}
	return d.name
}

func (d *values) OriginalName() string {
	return d.name
}

func (d *values) Kind() string {
	return "VALUE"
}

func (d *values) Path() string {
	return d.path
}

func (d *values) Dependencies() (deps []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	e := d.process()
	if e != nil {
		return deps, e
	}

	return d.deps, nil
}

func (d *values) Exported() bool {
	return d.exported
}

func (d *values) Omitted() bool {
	if d.tag.Omit {
		return true
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
func (d *values) Rewrite() def.Rewrite {
	return d.rewrite
}

// TODO: these are only needed to be able
// to rewrite
func (d *values) IsVariadic() bool {
	return false
}

// TODO: these are only needed to be able
// to rewrite
func (d *values) Params() (params []string) {
	return params
}
