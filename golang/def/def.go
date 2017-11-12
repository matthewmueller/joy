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
	OriginalName() string
	Exported() bool
	Omitted() bool
	Type() types.Type
	Kind() string
	Dependencies() ([]Definition, error)
	Imports() map[string]string
	FromRuntime() bool
}

// Field interface
type Field interface {
	Name() string
	Type() ast.Expr
}

// Rewrite interface
// type Rewrite interface {
// 	Rewrite(arguments *ast.FieldList) (string, error)
// }

// Rewrite interface
type Rewrite interface {
	Definition() Definition
	Expression() string
	Vars() []RewriteVariable
	Variadic() bool
}

// RewriteVariable interface
type RewriteVariable interface {
	Definition() Definition
	Node() ast.Expr
	String() (string, error)
}

// Edge interface
type Edge interface {
	Definition() Definition

	Type() string
}

// FunctionResult interface
type FunctionResult interface {
	Name() string
	Definition() Definition
}
