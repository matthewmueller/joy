package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/matthewmueller/joy/internal/chrome"
	"github.com/matthewmueller/joy/internal/compiler"
)

// RunSettings struct
type RunSettings struct {
	ChromePath  string
	FilePath    string
	Development bool
}

// Run fn
func Run(ctx context.Context, settings *RunSettings) (result string, err error) {
	if settings.ChromePath == "" {
		return "", errors.New("the Google Chrome path needs to be set")
	}

	c := compiler.New(&compiler.Params{
		Development: settings.Development,
	})

	files, err := c.Compile(settings.FilePath)
	if err != nil {
		return "", err
	} else if len(files) != 1 {
		return "", fmt.Errorf("joy run expects only 1 main file, but received %d files", len(files))
	}

	ch, err := chrome.New(ctx, &chrome.Settings{
		ExecutablePath: settings.ChromePath,
	})
	if err != nil {
		return "", err
	}
	defer ch.Close()

	tar, err := ch.Target()
	if err != nil {
		return "", err
	}
	defer tar.Close()

	return tar.Run(files[0].Source())
}
