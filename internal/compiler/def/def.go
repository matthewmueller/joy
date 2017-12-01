package def

import (
	"go/ast"
	"go/types"
)

// ID interface
type ID interface {
	ID() string
}

// Definition interface
type Definition interface {
	ID

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

// Rewritee is the original method or function getting rewritten
type Rewritee interface {
	ID() string
	Params() []string
	Rewrite() Rewrite
	IsVariadic() bool
}

// Rewrite interface
type Rewrite interface {
	Rewritee() Rewritee
	Expression() string
	Vars() []RewriteVariable
}

// RewriteVariable interface
type RewriteVariable interface {
	Definition() Definition
	Node() ast.Expr
	String() (string, error)
}

// FunctionResult interface
type FunctionResult interface {
	Name() string
	Definition() Definition
}

// InterfaceMethod interface
type InterfaceMethod interface {
	Rewritee

	OriginalName() string
	Name() string
}
