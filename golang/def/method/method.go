package method

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/fatih/structtag"
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
	index     *index.Index
	id        string
	path      string
	name      string
	kind      types.Type
	node      *ast.FuncDecl
	exported  bool
	params    []string
	recv      string
	async     bool
	processed bool
	deps      []def.Definition
	tag       *structtag.Tag
	runtime   bool
	imports   map[string]string
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

	// check if this declaration is from the runtime
	fromRuntime := false
	runtimePath, e := util.RuntimePath()
	if e != nil {
		return nil, e
	}
	if packagePath == runtimePath {
		fromRuntime = true
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

	tag, e := util.JSTag(n.Doc)
	if e != nil {
		return nil, e
	}

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
		tag:      tag,
		runtime:  fromRuntime,
	}, nil
}

func (d *methoddef) process() (err error) {
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

		case *ast.ChanType:
			deps, e := d.index.Runtime("Channel", "send", "Send", "Recv")
			if e != nil {
				err = e
				return false
			}
			d.deps = append(d.deps, deps...)
			for _, dep := range deps {
				d.imports["runtime"] = dep.Path()
				seen[dep.ID()] = true
			}
		}

		return true
	})

	d.processed = true
	return err
}

func (d *methoddef) Dependencies() (defs []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	err = d.process()
	if err != nil {
		return defs, err
	}
	return d.deps, nil
}

func (d *methoddef) IsAsync() (bool, error) {
	if d.processed {
		return d.async, nil
	}
	e := d.process()
	if e != nil {
		return false, e
	}
	return d.async, nil
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

func (d *methoddef) Recv() def.Definition {
	return d.index.Get(d.recv)
}

func (d *methoddef) Imports() map[string]string {
	// combine def imports with file imports
	imports := map[string]string{}
	for alias, path := range d.imports {
		imports[alias] = path
	}
	for alias, path := range d.index.GetImports(d.path) {
		imports[alias] = path
	}
	return imports
}

func (d *methoddef) FromRuntime() bool {
	return d.runtime
}
