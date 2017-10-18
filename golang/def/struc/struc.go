package struc

import (
	"fmt"
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/util"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"

	"golang.org/x/tools/go/loader"
)

// Struct interface
type Struct interface {
	def.Definition
	Node() *ast.TypeSpec
	Fields() []Field
	Field(name string) Field
	OriginalName() string
}

// Field interface
type Field interface {
	Name() string
	Type() ast.Expr
}

var _ Struct = (*structdef)(nil)
var _ Field = (*field)(nil)

type structdef struct {
	exported  bool
	path      string
	name      string
	id        string
	index     *index.Index
	node      *ast.TypeSpec
	kind      types.Type
	tag       *structtag.Tag
	fields    []*field
	deps      []def.Definition
	processed bool
}

// field of a struct
type field struct {
	name string
	tag  *structtag.Tag
	kind ast.Expr
}

func (f *field) Name() string {
	if f.tag != nil {
		return f.tag.Name
	}
	return f.name
}

func (f *field) Type() ast.Expr {
	return f.kind
}

// NewStruct fn
func NewStruct(index *index.Index, info *loader.PackageInfo, n *ast.TypeSpec) (def.Definition, error) {
	// typeof := info.TypeOf(n.Name)
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	idParts := []string{packagePath, n.Name.Name}
	id := strings.Join(idParts, " ")

	tag, e := util.JSTag(n.Doc)
	if e != nil {
		return nil, e
	}

	st, ok := n.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("NewStruct: expected a *ast.StructType, but got a %T", n.Type)
	}

	fields := []*field{}
	for _, f := range st.Fields.List {
		fld := &field{kind: f.Type}

		// add a tag if we have one
		if f.Tag != nil {
			v, e := strconv.Unquote(f.Tag.Value)
			if e != nil {
				return nil, e
			}
			tag, e := util.JSTagFromString(v)
			if e != nil {
				return nil, e
			}
			fld.tag = tag
		}

		// handle User{*dep.Settings} w/o name
		if len(f.Names) == 0 {
			id, e := util.GetIdentifier(f.Type)
			if e != nil {
				return nil, e
			}
			fld.name = id.Name
			fields = append(fields, fld)
			continue
		}

		for _, id := range f.Names {
			// log.Infof("names=%s", id.Name)
			fields = append(fields, &field{
				name: id.Name,
				kind: fld.kind,
				tag:  fld.tag,
			})
		}
	}

	return &structdef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     n.Name.Name,
		id:       id,
		index:    index,
		node:     n,
		kind:     info.TypeOf(n.Name),
		tag:      tag,
		fields:   fields,
	}, nil
}

func (d *structdef) process() (err error) {
	seen := map[string]bool{}
	_ = seen

	ast.Inspect(d.node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.SelectorExpr:
			def, e := d.index.DefinitionOf(d.Path(), t)
			if e != nil {
				err = e
				return false
			} else if def != nil && !seen[def.ID()] {
				d.deps = append(d.deps, def)
				seen[def.ID()] = true
			}

		case *ast.Ident:
			def, e := d.index.DefinitionOf(d.Path(), t)
			if e != nil {
				err = e
				return false
			} else if def != nil && !seen[def.ID()] {
				d.deps = append(d.deps, def)
				seen[def.ID()] = true
			}
		}

		return true
	})

	d.processed = true
	return err
}

func (d *structdef) Dependencies() (defs []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	err = d.process()
	if err != nil {
		return defs, err
	}
	return d.deps, nil
}

func (d *structdef) ID() string {
	return d.id
}

func (d *structdef) Name() string {
	if d.tag != nil {
		return d.tag.Name
	}
	return d.name
}

func (d *structdef) OriginalName() string {
	return d.name
}

func (d *structdef) Path() string {
	return d.path
}

func (d *structdef) Exported() bool {
	return d.exported
}

func (d *structdef) Omitted() bool {
	return false
}

func (d *structdef) Node() *ast.TypeSpec {
	return d.node
}

func (d *structdef) Type() types.Type {
	return d.kind
}

func (d *structdef) Fields() (fields []Field) {
	for _, f := range d.fields {
		fields = append(fields, f)
	}
	return fields
}

func (d *structdef) Field(name string) Field {
	// log.Infof("name=%s numfields=%s", name, len(d.fields))
	for _, f := range d.fields {
		// log.Infof("name=%s field=%s", name, f.name)
		if f.name == name {
			return f
		}
	}
	return nil
}

func (d *structdef) Imports() map[string]string {
	return d.index.GetImports(d.path)
}

func (d *structdef) FromRuntime() bool {
	return false
}
