// Package run compiles your Go file and runs it on headless chrome
//
// This is the programmatic API to running Joy. Run will be called
// by the CLI:
//
//    $ joy run main.go
//
package run

import (
	"context"
	"fmt"
	"path"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/mains"

	"github.com/matthewmueller/joy/internal/chrome"
	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/paths"
	"github.com/pkg/errors"
)

// Config struct
type Config struct {
	FilePath    string          // Go file to compile (required)
	Context     context.Context // Cancellable context (optional, default: background)
	Development bool            // Run the development bundle (optional, default: true)
	JoyPath     string          // Root path to Joy's state files (optional)
	Log         log.Interface   // Logger to use (optional)
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

// Run fn
func Run(cfg *Config) (result string, err error) {
	if err := cfg.defaults(); err != nil {
		return result, errors.Wrapf(err, "error getting defaults")
	}

	packages, err := mains.Find(cfg.FilePath)
	if err != nil {
		return result, errors.Wrapf(err, "error getting mains for filepath")
	}

	files, err := compiler.Compile(&compiler.Config{
		Packages:    packages,
		JoyPath:     cfg.JoyPath,
		Development: cfg.Development,
	})
	if err != nil {
		return result, errors.Wrapf(err, "unable to compile file")
	} else if len(files) == 0 {
		return result, fmt.Errorf("a main file requires a main function to build")
	} else if len(files) != 1 {
		return result, fmt.Errorf("joy run expects only 1 main file, but received %d files", len(files))
	}

	// download chrome if it doesn't already exist
	ch, err := chrome.Start(cfg.Context, path.Join(cfg.JoyPath, "chrome"))
	if err != nil {
		return result, errors.Wrapf(err, "error starting chrome")
	}
	defer ch.Close()

	tar, err := ch.Target()
	if err != nil {
		return result, errors.Wrapf(err, "error leasing target")
	}
	defer tar.Close()

	result, err = tar.Run(files[0].Source())
	if err != nil {
		return result, errors.Wrapf(err, "error running chrome target")
	}

	return result, nil
}
