package crawler

import (
	"fmt"
	"go/ast"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/compiler/index"
	"golang.org/x/tools/go/loader"
)

// state struct
type state struct {
	info  *loader.PackageInfo
	index *index.Index
}

// Crawl fn
func Crawl(program *loader.Program) (idx *index.Index, err error) {
	defer log.Trace("index").Stop(&err)

	// get the runtime path
	// runtimePath, err := util.RuntimePath()
	// if err != nil {
	// 	return nil, err
	// }

	idx = index.New(program)

	// decls := []index.Declaration{}
	for _, info := range program.AllPackages {
		for _, file := range info.Files {
			for _, decl := range file.Decls {
				s := &state{info, idx}
				if e := declarations(s, decl); e != nil {
					return nil, e
				}
			}
		}

	}

	return idx, nil
	// return &Index{decls}, nil
}

func declarations(s *state, n ast.Decl) (err error) {
	switch t := n.(type) {
	case *ast.FuncDecl:
		return funcDecl(s, t)
	case *ast.GenDecl:
		return genDecl(s, t)
	default:
		return fmt.Errorf("unhandled declaration: %T", n)
	}
}

func genDecl(s *state, n *ast.GenDecl) (err error) {
	for _, sp := range n.Specs {
		if e := spec(s, sp); e != nil {
			return e
		}
	}
	return nil
}

func spec(s *state, n ast.Spec) (err error) {
	switch t := n.(type) {
	case *ast.ValueSpec:
		return valueSpec(s, t)
	case *ast.TypeSpec:
		return typeSpec(s, t)
	case *ast.ImportSpec:
		return importSpec(s, t)
	default:
		return fmt.Errorf("unhandled declaration: %T", spec)
	}
}

func valueSpec(s *state, n *ast.ValueSpec) (err error) {
	v, e := index.NewVariable(s.info, n)
	if e != nil {
		return e
	}
	s.index.Add(v)
	return nil
}

func typeSpec(s *state, n *ast.TypeSpec) (err error) {
	switch n.Type.(type) {
	case *ast.StructType:
		return structType(s, n)
	case *ast.InterfaceType:
		return interfaceType(s, n)
	}
	return nil
}

// TODO: add to index
func importSpec(s *state, n *ast.ImportSpec) (err error) {
	// imp, err := NewImport(s.info, n)
	// if err != nil {
	// 	return err
	// }
	// s.index.AddImport(imp)
	return nil
}

func structType(s *state, n *ast.TypeSpec) (err error) {
	st, e := index.NewStruct(s.info, n)
	if e != nil {
		return e
	}
	s.index.Add(st)
	return nil
}

func interfaceType(s *state, n *ast.TypeSpec) (err error) {
	iface, e := index.NewInterface(s.info, n)
	if e != nil {
		return e
	}
	s.index.Add(iface)
	return nil
}

func funcDecl(s *state, n *ast.FuncDecl) (err error) {
	// packagePath := s.info.Pkg.Path()
	// obj := s.info.ObjectOf(n.Name)
	// name := n.Name.Name
	// id := obj.String()

	// // get the list of parameters
	// var params []string
	// for _, param := range n.Type.Params.List {
	// 	for _, id := range param.Names {
	// 		params = append(params, id.Name)
	// 	}
	// }

	// // if it's a method don't export,
	// // if it's the main() function
	// // export either way
	// exported := obj.Exported()
	// if n.Recv != nil {
	// 	exported = false
	// } else if name == "main" {
	// 	exported = true
	// }

	if n.Recv != nil {
		return s.index.Method(s.info, n)
	}

	return s.index.Function(s.info, n)
	// 	from: packagePath,
	// 	node: n,
	// 	id:   id,
	// }
	// declarations[id] = &types.Declaration{
	// 	From:     packagePath,
	// 	Exported: exported,
	// 	Name:     name,
	// 	ID:       id,
	// 	Node:     n,
	// 	Params:   params,
	// }

	// declaration may satisfy an interface
	// so we hold onto it for possible
	// inclusion later
	// if n.Recv != nil {
	// 	recv := n.Recv.List[0]
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

	// // point human-friendly names to the declaration
	// if runtimePath == packagePath {
	// 	runtime[name] = declarations[id]
	// }

	// return nil
}
