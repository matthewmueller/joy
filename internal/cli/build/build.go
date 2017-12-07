package build

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/compiler"
	"github.com/matthewmueller/joy/internal/mains"
	"github.com/matthewmueller/joy/internal/stats"
	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("build", "build Go packages into Javascript")
	packages := cmd.Arg("packages", "packages to build").Required().Strings()
	dev := cmd.Flag("dev", "generate a development build").Short('d').Bool()
	output := cmd.Flag("output", "directory to output files").Short('o').String()

	cmd.Action(func(_ *kingpin.ParseContext) (err error) {
		start := time.Now()

		// stats
		defer stats.TrackError("build", time.Now(), &err)

		if !*dev {
			log.Infof("Production builds coming soon! for now use `joy build --dev <packages>...` and run regenerator and uglify manually")
			return nil
		} else if *output != "" {
			log.Infof("Output templates coming soon! For now you can use `joy build --dev <packages>...`")
			return nil
		}

		gofiles, err := mains.Find(*packages)
		if err != nil {
			return err
		}

		c := compiler.New(&compiler.Params{
			Development: *dev,
		})
		jsfiles, e := c.Compile(gofiles...)
		if e != nil {
			return errors.Wrap(e, "error building packages")
		}

		cwd, err := os.Getwd()
		if err != nil {
			return errors.Wrapf(err, "error getting cwd")
		}

		var loc int
		for _, file := range jsfiles {
			filename := path.Base(file.Name())
			if err := ioutil.WriteFile(path.Join(cwd, filename+".js"), []byte(file.Source()), 0644); err != nil {
				return errors.Wrapf(err, "error writing %s", file.Path())
			}

			//  count lines of code to be able
			// to map filesize to performance
			loc += len(file.Source())
		}

		// stats
		stats.Track("build", map[string]interface{}{
			"duration": time.Since(start).Round(time.Millisecond).String(),
			"files":    len(jsfiles),
			"loc":      loc,
		})

		return nil
	})
}
