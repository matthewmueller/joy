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

// Declaration struct
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
	JSFieldTags map[string]*structtag.Tag

	Dependencies []*Declaration
	Rewrite      *Rewrite
	Async        bool
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
