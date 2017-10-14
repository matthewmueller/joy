package compiler

import (
	"github.com/dop251/goja/file"
	"github.com/matthewmueller/golly/compiler/assembler"
	"github.com/matthewmueller/golly/compiler/indexer"
	"github.com/matthewmueller/golly/compiler/loader"
	"github.com/matthewmueller/golly/compiler/parser"
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
func (c *Compiler) Compile() (files []file.File, err error) {
	program, err := loader.Load(&loader.Settings{Packages: c.settings.Packages})
	if err != nil {
		return files, errors.Wrap(err, "error loading the packages")
	}

	idx, err := indexer.New(program)
	if err != nil {
		return files, errors.Wrap(err, "error indexing the packages")
	}

	graph, err := parser.Parse(idx)
	if err != nil {
		return files, errors.Wrap(err, "error parsing the program")
	}

	files, err = assembler.Assemble(graph)
	if err != nil {
		return files, errors.Wrap(err, "error assembling the scripts")
	}

	return files, nil
}
