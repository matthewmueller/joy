package types

import "github.com/matthewmueller/golly/jsast"

// Declaration interface
type Declaration interface {
	ID() string
	Name() string
	Path() string
	Exported() bool
	Translate() (jsast.IStatement, error)
	Dependencies() ([]Declaration, error)
}

// File interface
type File interface {
	Name() string
	Path() string
	Source() string
}
