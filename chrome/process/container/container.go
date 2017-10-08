package container

import (
	"bufio"
	"context"
	"errors"
	"io"
	"strings"
	"time"

	dockrun "github.com/matthewmueller/go-dockrun"
	uuid "github.com/satori/go.uuid"
	tomb "gopkg.in/tomb.v2"
)

// Container struct
type Container struct {
	settings *Settings
	c        *dockrun.Runner
	t        tomb.Tomb
}

// Settings for the container
type Settings struct {
	Image  string
	Expose []string
	Stdout chan<- []byte
	Stderr chan<- []byte
}

// New container
func New(parent context.Context, settings *Settings) (*Container, error) {
	co := &Container{
		settings: settings,
	}

	client, err := dockrun.New()
	if err != nil {
		return nil, err
	}

	name := strings.Replace(settings.Image, "/", "-", -1)
	name = strings.Replace(name, ":", "-", -1)

	containerName := name + "-" + uuid.NewV4().String()
	container := client.Container(settings.Image, containerName)

	for _, portMap := range settings.Expose {
		container.Expose(portMap)
	}

	// cancellable downstream
	ctx := co.t.Context(parent)

	// start the container
	c, err := container.Run(ctx)
	if err != nil {
		return nil, err
	}
	co.c = c

	// if we've set stdout, pass that along
	if co.settings.Stdout != nil {
		co.t.Go(func() error {
			return co.stdout(ctx, c)
		})
	}

	// if we've set stderr, pass that along
	if co.settings.Stderr != nil {
		co.t.Go(func() error {
			return co.stderr(ctx, c)
		})
	}

	// wait until the container exits
	co.t.Go(c.Wait)

	// if we're dying, we want to
	// tell the container to exit
	// TODO: update container to die
	// when the context gets cancelled
	// parents shouldn't be telling
	// children to exit, the context
	// passed through should inform
	// the children to exit
	co.t.Go(func() error {
		<-ctx.Done()
		return co.doStop(5 * time.Second)
	})

	return co, nil
}

func (co *Container) stdout(ctx context.Context, c *dockrun.Runner) error {
	reader, writer := io.Pipe()
	buf := bufio.NewReader(reader)
	if e := c.Stdout(ctx, writer); e != nil {
		return e
	}

	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}

		select {
		case co.settings.Stdout <- line:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (co *Container) stderr(ctx context.Context, c *dockrun.Runner) error {
	reader, writer := io.Pipe()
	buf := bufio.NewReader(reader)
	if e := c.Stderr(ctx, writer); e != nil {
		return e
	}

	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}

		select {
		case co.settings.Stderr <- line:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Wait for the container to finish
func (co *Container) Wait() error {
	return co.t.Wait()
}

// Stop the container
func (co *Container) Stop() error {
	co.t.Kill(errors.New("stopped"))
	return co.Wait()
}

func (co *Container) doStop(timeout time.Duration) error {
	co.c.Stop(uint(timeout))
	return co.Wait()
}
