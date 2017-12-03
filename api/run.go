package api

import (
	"context"
	"errors"

	"github.com/matthewmueller/joy/internal/chrome"
)

// RunSettings struct
type RunSettings struct {
	ChromePath string
	FilePath   string
}

// Run fn
// TODO: pass in a writer
func Run(ctx context.Context, settings *RunSettings) (result string, err error) {
	if settings.ChromePath == "" {
		return "", errors.New("the Google Chrome path needs to be set")
	}

	files, err := Build(ctx, &BuildSettings{
		Packages: []string{settings.FilePath},
	})
	if err != nil {
		return "", err
	} else if len(files) != 1 {
		return "", errors.New("only 1 file allowed")
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
