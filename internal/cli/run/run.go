package run

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/api"
	"github.com/matthewmueller/joy/internal/chrome"
	"github.com/matthewmueller/joy/internal/compiler/util"
	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("run", "run command")
	file := cmd.Arg("file", "go file to compile and run").Required().String()

	cmd.Action(func(_ *kingpin.ParseContext) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		filePath := path.Join(cwd, *file)

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
			FilePath:   filePath,
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(result)
		return nil
	})
}
