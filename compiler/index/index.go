package index

import (
	"go/ast"

	"github.com/matthewmueller/golly/jsast"
	"golang.org/x/tools/go/loader"
)

// Declaration interface
type Declaration interface {
	ID() string
	Translate() (jsast.IStatement, error)
	Dependencies() ([]Declaration, error)
}

// Index struct
type Index struct {
	declarations []Declaration
}

// New function
func New() *Index {
	return &Index{}
}

// Function creates a function declaration
func (i *Index) Function(info *loader.PackageInfo, n *ast.FuncDecl) error {
	return nil
}

// Method creates a method declaration
func (i *Index) Method(info *loader.PackageInfo, n *ast.FuncDecl) error {
	return nil
}

// Mains gets the main entrypoints of the program
func (i *Index) Mains() (decls []Declaration, err error) {
	return decls, nil
}
