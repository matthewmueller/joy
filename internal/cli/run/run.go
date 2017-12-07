package run

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/api"
	"github.com/matthewmueller/joy/internal/chrome"
	"github.com/matthewmueller/joy/internal/compiler/util"
	"github.com/matthewmueller/joy/internal/stats"
	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("run", "compile and run a Go file")
	file := cmd.Arg("file", "Go file to compile and run").Required().String()

	cmd.Action(func(_ *kingpin.ParseContext) (err error) {
		start := time.Now()

		// stats
		defer stats.TrackError("run", time.Now(), &err)

		root, err := util.JoyPath()
		if err != nil {
			return errors.Wrapf(err, "error getting root path")
		}

	exists:
		chromePath, err := chrome.Exists(path.Join(root, "chrome"))
		if err != nil {
			return errors.Wrapf(err, "unable to get chrome path")
		} else if chromePath == "" {
			log.Infof("downloading headless chrome (this only needs to be done once)")
			if err := chrome.Download(path.Join(root, "chrome")); err != nil {
				return errors.Wrapf(err, "error downloading headless chrome")
			}
			goto exists
		}

		result, err := api.Run(ctx, &api.RunSettings{
			ChromePath: chromePath,
			FilePath:   *file,
		})
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(result)

		// stats
		stats.Track("run", map[string]interface{}{
			"duration": time.Since(start).Round(time.Millisecond).String(),
		})

		return nil
	})
}
