package types

import (
	"go/ast"

	"github.com/fatih/structtag"
)

// Script struct
type Script struct {
	Name     string
	Packages []*Package
	RawFiles []*File
}

// Package struct
type Package struct {
	Path         string
	Declarations []*Declaration
	Exports      []*Declaration

	// maps[alias]: packagePath
	Dependencies map[string]*Package
}

// Declaration is a fragment of code that
// may show up in the built code. A declaration
// could be a variable, constant, function,
// method or struct (but not an interface)
//
// TODO: this should be broken up into
// separate types and turned into an interface
// Right now fields can be present or not present
// based on type conditions.
//
// For now, try to keep this as data that only
// the translator depends on and not stuff that
// the inspector will use.
type Declaration struct {
	Exported bool
	From     string
	ID       string
	Node     ast.Decl
	Name     string

	// Just names for now because
	// we only need the names right
	// now for js.Rewrite. This is
	// only non-nil for FuncDecl's
	Params []string

	// This is nonnil when there is a comment above
	// the declaration
	// e.g. // js:"..."
	JSTag *structtag.Tag

	// This is a map of the struct tags on a field
	// e.g. Name string `js:"name"`
	Fields []Field

	Dependencies []*Declaration
	Rewrite      *Rewrite
	Async        bool
}

// Field contains struct field information
type Field struct {
	Name string
	Type string
	Tag  *structtag.Tag
}

// Rewrite struct for js.Rewrite
type Rewrite struct {
	Expression string
	Variables  []string
}

// File struct
type File struct {
	Name   string
	Source string
}
