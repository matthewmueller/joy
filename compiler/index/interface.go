package index

import (
	"go/ast"

	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/jsast"
	"golang.org/x/tools/go/loader"
)

var _ types.Declaration = (*Interface)(nil)

// NewInterface creates a new interface
func NewInterface(info *loader.PackageInfo, n *ast.TypeSpec) (*Interface, error) {
	return &Interface{}, nil
}

// Interface struct
type Interface struct {
	info *loader.PackageInfo
	node *ast.TypeSpec
}

// ID string
func (i *Interface) ID() string {
	return ""
}

// Name of the declaration
func (i *Interface) Name() string {
	return ""
}

// Exported fn
// TODO
func (i *Interface) Exported() bool {
	return true
}

// Path to the declaration
func (i *Interface) Path() string {
	return i.info.Pkg.Path()
}

// Translate fn
func (i *Interface) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (i *Interface) Dependencies() (deps []types.Declaration, err error) {
	return deps, nil
}
