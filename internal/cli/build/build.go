package build

import (
	"context"
	"fmt"

	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/mains"
	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("build", "build command")
	packages := cmd.Arg("package", "packages to build").Required().Strings()

	cmd.Action(func(_ *kingpin.ParseContext) error {
		gofiles, err := mains.Find(*packages)
		if err != nil {
			return err
		}

		// start := time.Now()
		c := compiler.New()
		jsfiles, e := c.Compile(gofiles...)
		if e != nil {
			return errors.Wrap(e, "error building packages")
		}

		for _, file := range jsfiles {
			fmt.Println(file.Source())
		}

		return nil
	})
}
