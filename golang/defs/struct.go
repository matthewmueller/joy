package defs

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

// Structer interface
type Structer interface {
	def.Definition
	Node() *ast.TypeSpec
	Fields() []def.Field
	Field(name string) def.Field
	OriginalName() string
	Methods() []def.Definition
}

var _ Structer = (*structdef)(nil)

type structdef struct {
	exported  bool
	path      string
	name      string
	id        string
	index     *index.Index
	gen       *ast.GenDecl
	node      *ast.TypeSpec
	kind      types.Type
	tag       *structtag.Tag
	fields    []def.Field
	deps      []def.Definition
	processed bool
	omit      bool
}

// Struct fn
func Struct(index *index.Index, info *loader.PackageInfo, gn *ast.GenDecl, n *ast.TypeSpec) (def.Definition, error) {
	// typeof := info.TypeOf(n.Name)
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}
	id := strings.Join(idParts, " ")

	tag, e := util.JSTag(gn.Doc)
	if e != nil {
		return nil, e
	}

	st, ok := n.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("Struct: expected a *ast.StructType, but got a %T", n.Type)
	}

	var fields []def.Field
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
		name:     name,
		id:       id,
		index:    index,
		node:     n,
		kind:     info.TypeOf(n.Name),
		tag:      tag,
		fields:   fields,
	}, nil
}

func (d *structdef) process() (err error) {
	state, e := process(d.index, d, d.node)
	if e != nil {
		return e
	}

	// copy state into function
	d.processed = true
	d.deps = state.deps
	d.omit = state.omit
	d.tag = state.tag
	d.fields = state.fields

	return nil
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

func (d *structdef) Kind() string {
	return "STRUCT"
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
	if d.tag != nil {
		return d.tag.HasOption("omit")
	}
	return false
}

func (d *structdef) Node() *ast.TypeSpec {
	return d.node
}

func (d *structdef) Type() types.Type {
	return d.kind
}

func (d *structdef) Fields() (fields []def.Field) {
	for _, f := range d.fields {
		fields = append(fields, f)
	}
	return fields
}

func (d *structdef) Field(name string) def.Field {
	// log.Infof("name=%s numfields=%s", name, len(d.fields))
	for _, f := range d.fields {
		// log.Infof("name=%s field=%s", name, f.name)
		if f.Name() == name {
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

// Methods functions
func (d *structdef) Methods() (defs []def.Definition) {
	return defs
}
