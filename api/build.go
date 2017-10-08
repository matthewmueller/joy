package api

import (
	"github.com/matthewmueller/golly/golang"
	"github.com/matthewmueller/golly/types"
)

// Build fn
func Build(pkgs ...string) (files []*types.File, err error) {
	compiler := golang.New()
	files, err = compiler.Compile(pkgs...)
	if err != nil {
		return nil, err
	}

	return files, nil
}
