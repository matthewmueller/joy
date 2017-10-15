package compiler

import (
	"github.com/matthewmueller/golly/compiler/assembler"
	"github.com/matthewmueller/golly/compiler/crawler"
	"github.com/matthewmueller/golly/compiler/loader"
	"github.com/matthewmueller/golly/compiler/parser"
	"github.com/matthewmueller/golly/compiler/types"
	"github.com/pkg/errors"
)

// Compiler struct
type Compiler struct {
	settings *Settings
}

// Settings struct
type Settings struct {
	Packages []string
}

// New fn
func New(settings *Settings) *Compiler {
	return &Compiler{
		settings: settings,
	}
}

// Compile fn
func (c *Compiler) Compile() (files []types.File, err error) {
	// Phase I: load the program and type check (heavy lifting done by golang's loader)
	program, err := loader.Load(&loader.Settings{Packages: c.settings.Packages})
	if err != nil {
		return files, errors.Wrap(err, "error loading the packages")
	}

	// Phase II: Shallow walk through the list of all declarations to build our index
	index, err := crawler.Crawl(program)
	if err != nil {
		return files, errors.Wrap(err, "error indexing the packages")
	}

	// Phase III: Deep walk through the declarations, using the index to add dependencies and build a dependency graph
	graph, err := parser.Parse(index)
	if err != nil {
		return files, errors.Wrap(err, "error parsing the program")
	}

	// Phase IV: Using the graph, sort it and assembles the declarations into actual javascript files.
	files, err = assembler.Assemble(graph)
	if err != nil {
		return files, errors.Wrap(err, "error assembling the files")
	}

	return files, nil
}
