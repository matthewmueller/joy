package process

import (
	"bufio"
	"context"
	"io"
	"os/exec"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// Process struct
type Process struct {
	cmd      *exec.Cmd
	settings *Settings
	eg       *errgroup.Group
	ctx      context.Context
}

// Settings for the container
type Settings struct {
	Path   string
	Flags  []string
	Stdout chan<- []byte
	Stderr chan<- []byte
}

// New process to create
func New(parent context.Context, settings *Settings) (*Process, error) {
	eg, ctx := errgroup.WithContext(parent)

	// create the command
	cmd := exec.CommandContext(ctx, settings.Path, settings.Flags...)

	p := &Process{
		settings: settings,
		cmd:      cmd,
		eg:       eg,
		ctx:      ctx,
	}

	// setup stdout if it was specified
	if settings.Stdout != nil {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		eg.Go(func() error {
			return p.stdio(settings.Stdout, stdout)
		})
	}

	// setup stderr if it was specified
	if settings.Stderr != nil {
		stderr, err := p.cmd.StderrPipe()
		if err != nil {
			return nil, err
		}
		eg.Go(func() error {
			return p.stdio(settings.Stderr, stderr)
		})
	}

	if e := cmd.Start(); e != nil {
		return nil, e
	}

	return p, nil
}

// Wait for the process to complete
func (p *Process) Wait() error {
	return p.eg.Wait()
}

// Stop the process
func (p *Process) Stop(timeout context.Context) error {
	p.cmd.Process.Signal(syscall.SIGINT)
	select {
	case <-p.ctx.Done():
		return p.eg.Wait()
	case <-timeout.Done():
		// kill after the context elapsed
		p.cmd.Process.Signal(syscall.SIGKILL)
		return p.eg.Wait()
	}
}

func (p *Process) stdio(stdioCh chan<- []byte, reader io.ReadCloser) error {
	buf := bufio.NewReader(reader)
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}

		select {
		case stdioCh <- line:
		// internally dying (from an error)
		case <-p.ctx.Done():
			return p.ctx.Err()
		}
	}
}
