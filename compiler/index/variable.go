package index

import (
	"go/ast"

	"golang.org/x/tools/go/loader"

	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/jsast"
)

var _ types.Declaration = (*Variable)(nil)

// NewVariable fn
func NewVariable(info *loader.PackageInfo, n *ast.ValueSpec) (*Variable, error) {
	return &Variable{}, nil
}

// Variable struct
type Variable struct {
	node *ast.ValueSpec
}

// ID string
func (d *Variable) ID() string {
	return ""
}

// Name of the declaration
func (d *Variable) Name() string {
	return ""
}

// Exported fn
// TODO
func (d *Variable) Exported() bool {
	return true
}

// Path string
func (d *Variable) Path() string {
	return ""
}

// Translate fn
func (d *Variable) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Variable) Dependencies() (deps []types.Declaration, err error) {
	return deps, nil
}
