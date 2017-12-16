package test

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/stats"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New test command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("test", "run go tests against headless chrome")
	verbose := cmd.Flag("verbose", "verbose flag").Short('v').Bool()
	pkgs := cmd.Arg("packages", "packages to run the tests on").Required().Strings()
	_, _ = verbose, pkgs
	cmd.Action(func(_ *kingpin.ParseContext) (err error) {
		start := time.Now()

		// stats
		defer stats.TrackError("test", time.Now(), &err)

		// packages, err := mains.Find(*pkgs...)
		// if err != nil {
		// 	return err
		// }

		// stats
		stats.Track("test", map[string]interface{}{
			"duration": time.Since(start).Round(time.Millisecond).String(),
			"packages": len(*pkgs),
		})

		log.Infof("Headless testing is coming soon!")
		return nil
	})
}
