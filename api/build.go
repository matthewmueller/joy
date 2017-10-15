package api

import (
	"context"

	"github.com/matthewmueller/golly/compiler"
	"github.com/matthewmueller/golly/compiler/types"
)

// BuildSettings struct
type BuildSettings struct {
	Packages []string
}

// Build fn
func Build(ctx context.Context, settings *BuildSettings) (files []types.File, err error) {
	c := compiler.New(&compiler.Settings{
		Packages: settings.Packages,
	})
	files, err = c.Compile()
	if err != nil {
		return nil, err
	}

	return files, nil
}
