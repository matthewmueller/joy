package test

import (
	"context"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New test command
func New(ctx context.Context, root *kingpin.Application) {
	cmd := root.Command("test", "run go tests against headless chrome")
	cmd.Action(func(_ *kingpin.ParseContext) error {
		println("Headless testing is coming soon!")
		return nil
	})
}
