package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/matthewmueller/joy/internal/mains"

	"github.com/pkg/errors"

	"github.com/matthewmueller/joy/api"
	"github.com/matthewmueller/joy/internal/cli/build"
	"github.com/matthewmueller/joy/internal/cli/run"
	"github.com/matthewmueller/joy/internal/cli/serve"
	"github.com/matthewmueller/joy/internal/cli/test"
	"github.com/matthewmueller/joy/internal/cli/upgrade"
	"github.com/matthewmueller/joy/internal/cli/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Run the CLI
func Run(ctx context.Context, ver string) error {
	// defer stats.Client.ConditionalFlush(100, 12*time.Hour)
	cmd := kingpin.New("joy", "Go to Javascript compiler")
	cmd.Version(ver)

	// special case: joy ./main.go
	if len(os.Args[1:]) == 1 && filepath.Ext(os.Args[1]) == ".go" {
		pkgs, err := mains.Find([]string{os.Args[1]})
		if err != nil {
			return errors.Wrapf(err, "error finding mains")
		} else if len(pkgs) == 0 {
			return errors.Wrapf(err, "no main file")
		}

		files, err := api.Build(ctx, &api.BuildSettings{
			Packages: pkgs,
		})

		if err != nil {
			return errors.Wrapf(err, "error building code")
		} else if len(files) != 1 {
			return fmt.Errorf("joy run expects only 1 main file, but received %d files", len(files))
		}

		fmt.Println(files[0].Source())
		return nil
	}

	// commands
	upgrade.New(ctx, cmd, ver)
	version.New(ctx, cmd, ver)
	build.New(ctx, cmd)
	serve.New(ctx, cmd)
	test.New(ctx, cmd)
	run.New(ctx, cmd)

	// run the command
	_, err := cmd.Parse(os.Args[1:])
	return err
}
