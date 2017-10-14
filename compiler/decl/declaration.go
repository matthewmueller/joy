package decl

import "github.com/matthewmueller/golly/jsast"

// Declaration interface
type Declaration interface {
	ID() string
	Translate() (jsast.IStatement, error)
	Dependencies() ([]Declaration, error)
}
