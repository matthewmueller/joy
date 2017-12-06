package api

import (
	"context"

	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/compiler/script"
)

// BuildSettings struct
type BuildSettings struct {
	Packages    []string
	Development bool
	Root        string
}

// Build fn
func Build(ctx context.Context, settings *BuildSettings) (scripts []*script.Script, err error) {
	c := compiler.New(&compiler.Params{
		Development: settings.Development,
	})

	scripts, err = c.Compile(settings.Packages...)
	if err != nil {
		return nil, err
	}

	return scripts, nil
}
