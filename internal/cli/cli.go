package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/prompt"
	"github.com/matthewmueller/joy/internal/stats"
	"github.com/matthewmueller/store"

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
func Run(ctx context.Context, ver string) (err error) {
	// setup stats
	defer func() {
		if err := stats.Client.MaybeFlush(100, 1*time.Minute); err != nil {
			log.WithError(err).Errorf("error flushing")
		}
	}()

	stats.Client.Set(map[string]interface{}{
		"os":      runtime.GOOS,
		"arch":    runtime.GOARCH,
		"version": ver,
	})

	cmd := kingpin.New("joy", "Joy – A Delightful Go to Javascript Compiler")
	cmd.Version(ver)

	// setup our local db
	db, err := store.New("joy")
	if err != nil {
		return errors.Wrapf(err, "unable to setup the storage")
	}

	runs, err := stats.Increment(db)
	if err != nil {
		return errors.Wrapf(err, "error incrementing stats")
	}

	if done, err := prompt.Prompt(db, runs); err != nil || done {
		return errors.Wrapf(err, "prompt error")
	}

	// special case: joy ./main.go
	if len(os.Args[1:]) == 1 && filepath.Ext(os.Args[1]) == ".go" {
		files, err := api.Build(ctx, &api.BuildSettings{
			Packages: []string{os.Args[1]},
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
