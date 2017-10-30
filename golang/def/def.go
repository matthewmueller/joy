package def

import (
	"go/ast"
	"go/types"
)

// Definition interface
type Definition interface {
	ID() string
	Path() string
	Name() string
	Exported() bool
	Omitted() bool
	Type() types.Type
	Kind() string
	Dependencies() ([]Edge, error)
	Imports() map[string]string
	FromRuntime() bool
}

// Field interface
type Field interface {
	Name() string
	Type() ast.Expr
}

// Rewrite interface
type Rewrite interface {
	Rewrite(arguments *ast.FieldList) (string, error)
}

// Edge interface
type Edge interface {
	Definition() Definition
	
	Type() string
}
