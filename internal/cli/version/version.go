package version

import (
	"context"

	"github.com/apex/log"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application, version string) {
	cmd := root.Command("version", "get the current version")
	cmd.Action(func(_ *kingpin.ParseContext) error {
		log.Infof(version)
		return nil
	})
}
