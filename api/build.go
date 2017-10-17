package api

import (
	"context"

	"github.com/matthewmueller/golly/golang"
	"github.com/matthewmueller/golly/golang/script"
)

// BuildSettings struct
type BuildSettings struct {
	Packages []string
}

// Build fn
func Build(ctx context.Context, settings *BuildSettings) (scripts []*script.Script, err error) {
	compiler := golang.New()
	scripts, err = compiler.Compile(settings.Packages...)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}
