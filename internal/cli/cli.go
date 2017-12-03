package cli

import (
	"context"
	"os"

	"github.com/matthewmueller/joy/internal/cli/build"
	"github.com/matthewmueller/joy/internal/cli/run"
	"github.com/matthewmueller/joy/internal/cli/serve"
	"github.com/matthewmueller/joy/internal/cli/upgrade"
	"github.com/matthewmueller/joy/internal/cli/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Run the CLI
func Run(ctx context.Context, ver string) error {
	// defer stats.Client.ConditionalFlush(100, 12*time.Hour)
	cmd := kingpin.New("joy", "Go to Javascript compiler")
	cmd.Version(ver)

	// commands
	upgrade.New(ctx, cmd, ver)
	version.New(ctx, cmd, ver)
	build.New(ctx, cmd)
	serve.New(ctx, cmd)
	run.New(ctx, cmd)

	_, err := cmd.Parse(os.Args[1:])

	return err
}
