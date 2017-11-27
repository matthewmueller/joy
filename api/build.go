package api

import (
	"context"

	"github.com/matthewmueller/golly/internal/compiler"
	"github.com/matthewmueller/golly/internal/compiler/script"
)

// BuildSettings struct
type BuildSettings struct {
	Packages []string
	Root string
}

// Build fn
func Build(ctx context.Context, settings *BuildSettings) (scripts []*script.Script, err error) {
	c := compiler.New()
	scripts, err = c.Compile(settings.Packages...)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}
