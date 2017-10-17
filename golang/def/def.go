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
	Dependencies() ([]Definition, error)
	Exported() bool
	Omitted() bool
	Node() ast.Node
	Type() types.Type
}
