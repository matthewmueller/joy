package defs

import (
	"go/ast"

	"github.com/fatih/structtag"
)

// field of a struct
type field struct {
	name     string
	tag      *structtag.Tag
	kind     ast.Expr
	embedded bool
}

func (f *field) Name() string {
	if f.tag != nil {
		return f.tag.Name
	}
	return f.name
}

func (f *field) Type() ast.Expr {
	return f.kind
}

func (f *field) Embedded() bool {
	return f.embedded
}
