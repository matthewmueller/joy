package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/matthewmueller/joy/internal/prompt"
	"github.com/matthewmueller/store"

	"github.com/pkg/errors"

	apibuild "github.com/matthewmueller/joy/api/build"
	"github.com/matthewmueller/joy/internal/cli/build"
	"github.com/matthewmueller/joy/internal/cli/run"
	"github.com/matthewmueller/joy/internal/cli/serve"
	"github.com/matthewmueller/joy/internal/cli/test"
	"github.com/matthewmueller/joy/internal/cli/upgrade"
	"github.com/matthewmueller/joy/internal/cli/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Run the CLI
func Run(ctx context.Context, ver string) (err error) {
	cmd := kingpin.New("joy", "Joy – A Delightful Go to Javascript Compiler")
	cmd.Version(ver)

	// setup our local db
	db, err := store.New("joy")
	if err != nil {
		return errors.Wrapf(err, "unable to setup the storage")
	}

	if ver != "master" {
		if done, err := prompt.Prompt(db); err != nil || done {
			return errors.Wrapf(err, "prompt error")
		}
	}

	// special case: joy ./main & ./main.go
	if len(os.Args[1:]) == 1 {
		if _, err := os.Stat(os.Args[1]); !os.IsNotExist(err) {
			files, err := apibuild.Build(&apibuild.Config{
				Context:  ctx,
				Packages: []string{os.Args[1]},
			})
			if err != nil {
				return errors.Wrapf(err, "error building code")
			} else if len(files) == 0 {
				return fmt.Errorf("a main file requires a main function to build")
			} else if len(files) != 1 {
				return fmt.Errorf("joy run expects only 1 main file, but received %d files", len(files))
			}

			fmt.Println(files[0].Source())
			return nil
		}
	}

	// commands
	run.New(ctx, cmd)
	build.New(ctx, cmd)
	serve.New(ctx, cmd)
	test.New(ctx, cmd)
	upgrade.New(ctx, cmd, ver)
	version.New(ctx, cmd, ver)

	// run the command
	_, err = cmd.Parse(os.Args[1:])
	return err
}
