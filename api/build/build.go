package build

import (
	"context"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/mains"
	"github.com/matthewmueller/joy/internal/paths"

	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/compiler/script"
	"github.com/pkg/errors"
)

// Config struct
type Config struct {
	Context     context.Context
	Packages    []string
	Development bool
	JoyPath     string
	Log         log.Interface // Log (optional)
}

func (c *Config) defaults() error {
	if c.Context == nil {
		c.Context = context.Background()
	}

	if c.Log == nil {
		c.Log = log.Log
	}

	if c.JoyPath == "" {
		p, err := paths.Joy()
		if err != nil {
			return errors.Wrapf(err, "error getting joy's root path")
		}
		c.JoyPath = p
	}

	return nil
}

// Build fn
func Build(cfg *Config) (scripts []*script.Script, err error) {
	if err := cfg.defaults(); err != nil {
		return scripts, errors.Wrapf(err, "error setting defaults")
	}

	packages, err := mains.Find(cfg.Packages...)
	if err != nil {
		return scripts, errors.Wrapf(err, "error finding mains")
	}

	scripts, err = compiler.Compile(&compiler.Config{
		Packages:    packages,
		Development: cfg.Development,
		JoyPath:     cfg.JoyPath,
	})
	if err != nil {
		return scripts, errors.Wrapf(err, "error compiling")
	}

	return scripts, nil
}
