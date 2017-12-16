package run

import (
	"context"
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/api/run"
	"github.com/matthewmueller/joy/internal/stats"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("run", "compile and run a Go file")
	file := cmd.Arg("file", "Go file to compile and run").Required().String()

	cmd.Action(func(_ *kingpin.ParseContext) (err error) {
		start := time.Now()

		// stats
		defer stats.TrackError("run", time.Now(), &err)

		result, err := run.Run(&run.Config{
			Context:     ctx,
			FilePath:    *file,
			Development: true, // TODO: change
		})
		if err != nil {
			log.WithError(err).Error("error running script")
			return err
		}

		// stats
		stats.Track("run", map[string]interface{}{
			"duration": time.Since(start).Round(time.Millisecond).String(),
		})

		fmt.Println(result)
		return nil
	})
}
