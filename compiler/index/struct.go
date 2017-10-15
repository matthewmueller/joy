package index

import (
	"go/ast"

	"golang.org/x/tools/go/loader"

	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/jsast"
)

var _ types.Declaration = (*Struct)(nil)

// NewStruct fn
func NewStruct(info *loader.PackageInfo, n *ast.TypeSpec) (*Struct, error) {
	return &Struct{}, nil
}

// Struct struct
type Struct struct {
	node *ast.TypeSpec
}

// ID string
func (d *Struct) ID() string {
	return ""
}

// Name of the declaration
func (d *Struct) Name() string {
	return ""
}

// Exported fn
// TODO
func (d *Struct) Exported() bool {
	return true
}

// Path string
func (d *Struct) Path() string {
	return ""
}

// Translate fn
func (d *Struct) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Struct) Dependencies() (deps []types.Declaration, err error) {
	return deps, nil
}
