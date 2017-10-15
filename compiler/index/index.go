package index

import (
	"fmt"

	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/compiler/util"
	"golang.org/x/tools/go/loader"
)

// Index struct
type Index struct {
	decls   map[string]types.Declaration
	program *loader.Program
}

// New fn
func New(program *loader.Program) *Index {
	return &Index{
		decls:   map[string]types.Declaration{},
		program: program,
	}
}

// Add fn
func (i *Index) Add(d types.Declaration) {
	i.decls[d.ID()] = d
}

// AddImport fn
// func (i *Index) AddImport(d types.Import) {

// }

// Mains finds all the main() functions inside our initial packages
func (i *Index) Mains() (decls []types.Declaration, err error) {
	runtimePath, err := util.RuntimePath()
	if err != nil {
		return decls, err
	}

	for _, info := range i.program.InitialPackages() {
		packagePath := info.Pkg.Path()

		// ignore the runtime
		if runtimePath == packagePath {
			continue
		}

		// find the main objects using the package scope
		obj := info.Pkg.Scope().Lookup("main")
		if obj == nil {
			return decls, fmt.Errorf("main not found in %s", packagePath)
		}

		// find the declaration in the index
		main := i.decls[obj.String()]
		if main == nil {
			return decls, fmt.Errorf("main declaration not found: %s", obj.String())
		}

		decls = append(decls, main)
	}
	
	return decls, nil
}
