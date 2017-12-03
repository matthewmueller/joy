package serve

import (
	"context"
	"strconv"

	"github.com/matthewmueller/joy/api"
	"github.com/matthewmueller/joy/internal/mains"
	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("serve", "serve command")
	pkgs := cmd.Arg("packages", "packages to bundle").Required().Strings()
	port := cmd.Flag("port", "port to serve from").Default("8080").String()

	cmd.Action(func(_ *kingpin.ParseContext) error {
		packages, err := mains.Find(*pkgs)
		if err != nil {
			return err
		}

		port, e := strconv.Atoi(*port)
		if e != nil {
			return errors.Wrap(e, "invalid port")
		}

		return api.Serve(ctx, &api.ServeSettings{
			Packages: packages,
			Port:     port,
		})
	})
}
