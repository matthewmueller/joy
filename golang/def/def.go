package def

import (
	"go/types"
)

// Definition interface
type Definition interface {
	ID() string
	Path() string
	Name() string
	Exported() bool
	Omitted() bool
	// Node() ast.Node
	Type() types.Type
	Dependencies() ([]Definition, error)
	// Translate() (jsast.IStatement, error)
}
