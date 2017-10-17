package method

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/def/fn"
	"github.com/matthewmueller/golly/golang/index"
	"github.com/matthewmueller/golly/golang/util"
	"golang.org/x/tools/go/loader"
)

// Method interface
type Method interface {
	fn.Function
	Recv() def.Definition
}

var _ Method = (*methoddef)(nil)

type methoddef struct {
	index    *index.Index
	id       string
	path     string
	name     string
	kind     types.Type
	node     *ast.FuncDecl
	exported bool
	params   []string
	recv     string
}

// NewMethod fn
func NewMethod(index *index.Index, info *loader.PackageInfo, n *ast.FuncDecl) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}
	var recv string

	// if it has a receiver, include that in the ID
	if n.Recv != nil && len(n.Recv.List) > 0 {
		id, e := util.GetIdentifier(n.Recv.List[0].Type)
		if e != nil {
			return nil, e
		}
		idParts = append(idParts, id.Name)
		recv = packagePath + " " + id.Name
	}

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

	return &methoddef{
		id:       id,
		index:    index,
		exported: exported,
		path:     packagePath,
		name:     name,
		node:     n,
		params:   params,
		kind:     info.TypeOf(n.Name),
		recv:     recv,
	}, nil
}

func (d *methoddef) ID() string {
	return d.id
}

func (d *methoddef) Name() string {
	return d.name
}

func (d *methoddef) Path() string {
	return d.path
}

func (d *methoddef) Dependencies() ([]def.Definition, error) {
	return nil, nil
}

func (d *methoddef) Exported() bool {
	return d.exported
}

func (d *methoddef) Omitted() bool {
	return false
}
func (d *methoddef) Node() *ast.FuncDecl {
	return d.node
}

func (d *methoddef) Type() types.Type {
	return d.kind
}

func (d *methoddef) IsAsync() bool {
	return false
}

func (d *methoddef) Recv() def.Definition {
	return d.index.Get(d.recv)
}

func (d *methoddef) InRuntime() bool {
	return false
}
