package test

import (
	"context"

	"github.com/apex/log"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New test command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("test", "run go tests against headless chrome")
	verbose := cmd.Flag("verbose", "verbose flag").Short('v').Bool()
	pkgs := cmd.Arg("packages", "packages to run the tests on").Required().Strings()
	_, _ = verbose, pkgs
	cmd.Action(func(_ *kingpin.ParseContext) (err error) {
		// packages, err := mains.Find(*pkgs...)
		// if err != nil {
		// 	return err
		// }
		log.Infof("Headless testing is coming soon!")
		return nil
	})
}
