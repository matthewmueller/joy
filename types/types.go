package types

import (
	"go/ast"

	"github.com/fatih/structtag"
)

// Script struct
type Script struct {
	Name     string
	Packages []*Package
}

// Package struct
type Package struct {
	Path         string
	Declarations []*Declaration

	// available after inspection
	// maps[alias]: packagePath
	Dependencies map[string]string
}

// Declaration struct
type Declaration struct {
	Exported bool
	From     string
	ID       string
	Node     ast.Decl
	Name     string

	// Note that this represents *all*
	// imports in the file, not just
	// the ones that the declaration
	// uses. We use this information
	// to fill in the package dependencies
	// are *only* the imports the
	// declarations use
	Imports map[string]string

	// The following are available after inspection
	JSTag        *structtag.Tag
	Dependencies []*Declaration
	Async        bool
}

// File struct
type File struct {
	Name   string
	Source string
}
