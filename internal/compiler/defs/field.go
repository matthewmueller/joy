package defs

import (
	"go/ast"

	"github.com/matthewmueller/joy/internal/compiler/util"
)

// field of a struct
type field struct {
	name     string
	tag      util.JSTag
	kind     ast.Expr
	embedded bool
}

func (f *field) Name() string {
	if f.tag.Rename != "" {
		return f.tag.Rename
	}
	return f.name
}

func (f *field) Type() ast.Expr {
	return f.kind
}

func (f *field) Embedded() bool {
	return f.embedded
}
