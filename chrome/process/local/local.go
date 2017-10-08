package local

// TODO
// - there's a race when streaming stdout/stderr pipes
//   we need to make sure we've read everything
//   before calling Wait()

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// Local struct
type Local struct {
	cmd      *exec.Cmd
	settings *Settings
	eg       *errgroup.Group
	ctx      context.Context
	once     sync.Once
}

// Settings for the container
type Settings struct {
	Name   string
	Flags  []string
	Stdout chan<- []byte
	Stderr chan<- []byte
}

// New process to create
func New(parent context.Context, settings *Settings) (*Local, error) {
	// create the command
	cmd := exec.CommandContext(parent, settings.Name, settings.Flags...)

	eg, ctx := errgroup.WithContext(parent)

	l := &Local{
		settings: settings,
		cmd:      cmd,
		ctx:      ctx,
		eg:       eg,
	}

	// setup stdout if it was specified
	if settings.Stdout != nil {
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		eg.Go(func() error {
			e := l.stdio(settings.Stdout, stdout)
			return e
		})
		eg.Go(func() error {
			<-ctx.Done()
			return stdout.Close()
		})

	}

	// setup stderr if it was specified
	if settings.Stderr != nil {
		stderr, err := l.cmd.StderrPipe()
		if err != nil {
			return nil, err
		}
		eg.Go(func() error {
			e := l.stdio(settings.Stderr, stderr)
			return e
		})
		eg.Go(func() error {
			<-ctx.Done()
			return stderr.Close()
		})
	}

	if e := cmd.Start(); e != nil {
		return l, e
	}

	// if the command dies, trigger a shutdown
	eg.Go(func() error {
		return cmd.Wait()
	})

	// if anything else dies, trigger a shutdown
	eg.Go(func() error {
		<-ctx.Done()
		return l.doStop(5 * time.Second)
	})

	return l, nil
}

// Wait for the process to complete
func (l *Local) Wait() error {
	return l.eg.Wait()
}

// Stop the process
// We only trigger ctx.Done() here
// because we don't want to end up
// in a deadlock
func (l *Local) Stop() (err error) {
	l.eg.Go(func() error { return errors.New("stopped") })
	return l.eg.Wait()
}

func (l *Local) doStop(timeout time.Duration) (err error) {
	if e := l.cmd.Process.Signal(syscall.SIGINT); e != nil {
		return e
	}

	done := make(chan error, 1)
	go func() {
		done <- l.cmd.Wait()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		// kill after the context elapsed
		l.cmd.Process.Signal(syscall.SIGKILL)
		select {
		case err := <-done:
			return err
		case <-time.After(10 * time.Second):
			return errors.New("couldn't kill process after 10s")
		}
	}
}

func (l *Local) stdio(stdioCh chan<- []byte, reader io.ReadCloser) error {
	buf := bufio.NewReader(reader)

	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}

		select {
		case stdioCh <- line:
		// internally dying (from an error)
		case <-l.ctx.Done():
			return l.ctx.Err()
		}
	}
}
