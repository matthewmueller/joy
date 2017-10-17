package fn

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/util"

	"golang.org/x/tools/go/loader"
)

// Function interface
type Function interface {
	def.Definition
	IsAsync() bool
	InRuntime() bool
	Node() *ast.FuncDecl
}

type funcdef struct {
	info     *loader.PackageInfo
	index    *index.Index
	id       string
	path     string
	name     string
	kind     types.Type
	node     *ast.FuncDecl
	exported bool
	params   []string
	tag      *structtag.Tag
}

var _ Function = (*funcdef)(nil)

// NewFunction fn
func NewFunction(index *index.Index, info *loader.PackageInfo, n *ast.FuncDecl) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}

	// if it has a receiver, include that in the ID
	// if n.Recv != nil && len(n.Recv.List) > 0 {
	// 	id, e := util.GetIdentifier(n.Recv.List[0].Type)
	// 	if e != nil {
	// 		return nil, e
	// 	}
	// 	idParts = append(idParts, id.Name)
	// }

	id := strings.Join(idParts, " ")

	var params []string
	for _, param := range n.Type.Params.List {
		for _, ident := range param.Names {
			params = append(params, ident.Name)
		}
	}

	// if it's a method don't export,
	// if it's the main() function
	// export either way
	exported := obj.Exported()
	if n.Recv != nil {
		exported = false
	} else if name == "main" {
		exported = true
	}

	// declarations[id] = &fndef{
	// 	Exported: exported,
	// 	From:     packagePath,
	// 	Name:     name,
	// 	ID:       id,
	// 	Node:     decl,
	// 	Params:   params,
	// }

	// declaration may satisfy an interface
	// so we hold onto it for possible
	// inclusion later
	// if t.Recv != nil {
	// 	recv := t.Recv.List[0]
	// 	if receivers[name] == nil {
	// 		receivers[name] = []*receiver{}
	// 	}
	// 	receivers[name] = append(
	// 		receivers[name],
	// 		&receiver{
	// 			Type:     info.TypeOf(recv.Type),
	// 			Function: declarations[id],
	// 		},
	// 	)
	// }

	// point human-friendly names to the declaration

	return &funcdef{
		index:    index,
		info:     info,
		id:       id,
		exported: exported,
		path:     packagePath,
		name:     name,
		node:     n,
		params:   params,
		kind:     info.TypeOf(n.Name),
	}, nil
}

func (d *funcdef) ID() string {
	return d.id
}

func (d *funcdef) Name() string {
	return d.name
}

func (d *funcdef) Path() string {
	return d.path
}

func (d *funcdef) Dependencies() (defs []def.Definition, err error) {
	seen := map[string]bool{}

	ast.Inspect(d.node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			tag, e := util.JSTag(t.Doc)
			if e != nil {
				err = e
				return false
			}
			d.tag = tag

		case *ast.SelectorExpr:
			def, e := d.index.DefinitionOf(d.info, t)
			if e != nil {
				err = e
				return false
			} else if def != nil && !seen[def.ID()] {
				defs = append(defs, def)
				seen[def.ID()] = true
			}

		case *ast.Ident:
			def, e := d.index.DefinitionOf(d.info, t)
			if e != nil {
				err = e
				return false
			} else if def != nil && !seen[def.ID()] {
				defs = append(defs, def)
				seen[def.ID()] = true
			}
		}
		return true
	})

	return defs, nil
}

func (d *funcdef) Exported() bool {
	return d.exported
}

func (d *funcdef) Omitted() bool {
	return false
}

func (d *funcdef) Node() *ast.FuncDecl {
	return d.node
}

func (d *funcdef) Type() types.Type {
	return d.kind
}

func (d *funcdef) IsAsync() bool {
	return false
}

func (d *funcdef) InRuntime() bool {
	return false
}
