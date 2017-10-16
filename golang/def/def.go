package def

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/util"

	"golang.org/x/tools/go/loader"
)

// Kind of definition
// type Kind int

// // enums
// const (
// 	// FUNCTION Kind
// 	FUNCTION = Kind(iota)
// 	// METHOD Kind
// 	METHOD
// 	// STRUCT Kind
// 	STRUCT
// 	// INTERFACE Kind
// 	INTERFACE
// 	// VALUE Kind
// 	VALUE
// )

// Definition interface
type Definition interface {
	ID() string
	Path() string
	Name() string
	Dependencies() []Definition
	Exported() bool
	Omit() bool
	Node() ast.Node
	Type() types.Type
}

// Interface method
type Interface interface {
	Definition
	ImplementedBy(method string) []Method
}

// Function interface
type Function interface {
	Definition
	IsAsync() bool
}

// Method interface
type Method interface {
	Function
	Recv() Definition
}

var _ Function = (*funcdef)(nil)
var _ Method = (*methoddef)(nil)
var _ Interface = (*ifacedef)(nil)
var _ Definition = (*structdef)(nil)
var _ Definition = (*valuedef)(nil)

type funcdef struct {
	idx      map[string]Definition
	id       string
	path     string
	name     string
	kind     types.Type
	node     *ast.FuncDecl
	exported bool
	params   []string
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

func (d *funcdef) Dependencies() []Definition {
	return nil
}

func (d *funcdef) Exported() bool {
	return d.exported
}

func (d *funcdef) Omit() bool {
	return false
}

func (d *funcdef) Node() ast.Node {
	return d.node
}

func (d *funcdef) Type() types.Type {
	return d.kind
}

func (d *funcdef) IsAsync() bool {
	return false
}

// NewFunction fn
func NewFunction(index map[string]Definition, info *loader.PackageInfo, n *ast.FuncDecl) (Definition, error) {
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
		id:       id,
		exported: exported,
		path:     packagePath,
		name:     name,
		node:     n,
		params:   params,
		kind:     info.TypeOf(n.Name),
	}, nil
}

type methoddef struct {
	idx      map[string]Definition
	id       string
	path     string
	name     string
	kind     types.Type
	node     *ast.FuncDecl
	exported bool
	params   []string
	recv     string
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

func (d *methoddef) Dependencies() []Definition {
	return nil
}

func (d *methoddef) Exported() bool {
	return d.exported
}

func (d *methoddef) Omit() bool {
	return false
}
func (d *methoddef) Node() ast.Node {
	return d.node
}

func (d *methoddef) Type() types.Type {
	return d.kind
}

func (d *methoddef) IsAsync() bool {
	return false
}

func (d *methoddef) Recv() Definition {
	return d.idx[d.recv]
}

// NewMethod fn
func NewMethod(index map[string]Definition, info *loader.PackageInfo, n *ast.FuncDecl) (Definition, error) {
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
		idx:      index,
		exported: exported,
		path:     packagePath,
		name:     name,
		node:     n,
		params:   params,
		kind:     info.TypeOf(n.Name),
		recv:     recv,
	}, nil
}

type structdef struct {
	exported bool
	path     string
	name     string
	id       string
	idx      map[string]Definition
	node     *ast.TypeSpec
	kind     types.Type
}

func (d *structdef) ID() string {
	return d.id
}

func (d *structdef) Name() string {
	return d.name
}

func (d *structdef) Path() string {
	return d.path
}

func (d *structdef) Dependencies() []Definition {
	return nil
}

func (d *structdef) Exported() bool {
	return d.exported
}

func (d *structdef) Omit() bool {
	return false
}

func (d *structdef) Node() ast.Node {
	return d.node
}

func (d *structdef) Type() types.Type {
	return d.kind
}

// NewStruct fn
func NewStruct(index map[string]Definition, info *loader.PackageInfo, n *ast.TypeSpec) (Definition, error) {
	// typeof := info.TypeOf(n.Name)
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	idParts := []string{packagePath, n.Name.Name}
	id := strings.Join(idParts, " ")

	return &structdef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     n.Name.Name,
		id:       id,
		idx:      index,
		node:     n,
		kind:     info.TypeOf(n.Name),
	}, nil
}

type valuedef struct {
	exported bool
	path     string
	name     string
	id       string
	idx      map[string]Definition
	node     *ast.Ident
	kind     types.Type
}

func (d *valuedef) ID() string {
	return d.id
}

func (d *valuedef) Name() string {
	return d.name
}

func (d *valuedef) Path() string {
	return d.path
}

func (d *valuedef) Dependencies() []Definition {
	return nil
}

func (d *valuedef) Exported() bool {
	return d.exported
}

func (d *valuedef) Omit() bool {
	return false
}

func (d *valuedef) Node() ast.Node {
	return d.node
}

func (d *valuedef) Type() types.Type {
	return d.kind
}

// NewValue fn
func NewValue(index map[string]Definition, info *loader.PackageInfo, n *ast.Ident) (Definition, error) {
	obj := info.ObjectOf(n)
	packagePath := obj.Pkg().Path()
	idParts := []string{packagePath, n.Name}
	id := strings.Join(idParts, " ")

	return &valuedef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     n.Name,
		id:       id,
		idx:      index,
		node:     n,
		// kind:     VALUE,
	}, nil
}

type ifacedef struct {
	exported bool
	path     string
	name     string
	id       string
	idx      map[string]Definition
	methods  []string
	node     *ast.TypeSpec
	kind     *types.Interface
}

func (d *ifacedef) ID() string {
	return d.id
}

func (d *ifacedef) Name() string {
	return d.name
}

func (d *ifacedef) Path() string {
	return d.path
}

func (d *ifacedef) Dependencies() []Definition {
	return nil
}

func (d *ifacedef) Exported() bool {
	return d.exported
}
func (d *ifacedef) Omit() bool {
	return false
}

func (d *ifacedef) Node() ast.Node {
	return d.node
}

func (d *ifacedef) Type() types.Type {
	return d.kind
}

// TODO: optimize
func (d *ifacedef) ImplementedBy(method string) (defs []Method) {
	for _, n := range d.idx {
		m, ok := n.(Method)
		if !ok {
			continue
		}

		if m.Name() != method {
			continue
		}

		if types.Implements(m.Recv().Type(), d.kind) {
			defs = append(defs, m)
		}
	}

	return defs
}

// NewInterface fn
func NewInterface(index map[string]Definition, info *loader.PackageInfo, n *ast.TypeSpec) (Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	idParts := []string{packagePath, n.Name.Name}
	id := strings.Join(idParts, " ")
	methods := []string{}

	iface := n.Type.(*ast.InterfaceType)
	for _, list := range iface.Methods.List {
		for _, method := range list.Names {
			methods = append(methods, method.Name)
		}
	}

	kind, ok := info.TypeOf(n.Type).(*types.Interface)
	if !ok {
		return nil, fmt.Errorf("NewInterface: expected n.Type to be an interface")
	}

	return &ifacedef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     n.Name.Name,
		id:       id,
		idx:      index,
		node:     n,
		methods:  methods,
		kind:     kind,
	}, nil
}
