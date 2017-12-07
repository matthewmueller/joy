package version

import (
	"context"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/stats"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application, version string) {
	cmd := root.Command("version", "get the current version")
	cmd.Action(func(_ *kingpin.ParseContext) error {
		stats.Track("version", map[string]interface{}{})
		log.Infof(version)
		return nil
	})
}
