package run

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/mains"

	"github.com/matthewmueller/joy/internal/chrome"
	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/paths"
	"github.com/pkg/errors"
)

// Config struct
type Config struct {
	Context     context.Context
	FilePath    string
	Development bool
	ChromePath  string
	MacroPath   string
	RuntimePath string
	StdPath     string
	Log         log.Interface // Log (optional)
}

func (c *Config) defaults() error {
	if c.Context == nil {
		c.Context = context.Background()
	}

	if c.Log == nil {
		c.Log = log.Log
	}

	if c.MacroPath == "" {
		p, err := paths.Macro()
		if err != nil {
			return errors.Wrapf(err, "error getting macro path")
		}
		c.MacroPath = p
	}

	if c.RuntimePath == "" {
		p, err := paths.Runtime()
		if err != nil {
			return errors.Wrapf(err, "error getting runtime path")
		}
		c.RuntimePath = p
	}

	if c.StdPath == "" {
		p, err := paths.Stdlib()
		if err != nil {
			return errors.Wrapf(err, "error getting std path")
		}
		c.StdPath = p
	}

	if c.ChromePath == "" {
		p, err := paths.Chrome()
		if err != nil {
			return errors.Wrapf(err, "error getting chrome path")
		}
		c.ChromePath = p
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
		Development: cfg.Development,
		MacroPath:   cfg.MacroPath,
		RuntimePath: cfg.RuntimePath,
		StdPath:     cfg.StdPath,
	})
	if err != nil {
		return result, errors.Wrapf(err, "unable to compile file")
	} else if len(files) != 1 {
		return result, fmt.Errorf("joy run expects only 1 main file, but received %d files", len(files))
	}

	// download chrome if it doesn't already exist
exists:
	fullpath, err := chrome.Exists(cfg.ChromePath)
	if err != nil {
		return result, errors.Wrapf(err, "unable to get chrome path")
	} else if fullpath == "" {
		cfg.Log.Infof("downloading headless chrome (this only needs to be done once)")
		if err := chrome.Download(cfg.ChromePath); err != nil {
			return result, errors.Wrapf(err, "error downloading headless chrome")
		}
		goto exists
	}

	ch, err := chrome.New(cfg.Context, &chrome.Settings{
		ExecutablePath: cfg.ChromePath,
	})
	if err != nil {
		return result, err
	}
	defer ch.Close()

	tar, err := ch.Target()
	if err != nil {
		return result, err
	}
	defer tar.Close()

	return tar.Run(files[0].Source())
}
