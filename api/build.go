package api

import (
	"context"

	"github.com/matthewmueller/golly/golang"
	"github.com/matthewmueller/golly/types"
)

// BuildSettings struct
type BuildSettings struct {
	Packages []string
}

// Build fn
func Build(ctx context.Context, settings *BuildSettings) (files []*types.File, err error) {
	compiler := golang.New()
	files, _, err = compiler.Compile(settings.Packages...)
	if err != nil {
		return nil, err
	}

	return files, nil
}
