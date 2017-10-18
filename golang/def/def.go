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
	Type() types.Type
	Dependencies() ([]Definition, error)
	Imports() map[string]string
	FromRuntime() bool
}
