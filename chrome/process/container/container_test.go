package container_test

import (
	"context"
	"net"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/mafredri/cdp/devtool"
	"github.com/matthewmueller/voodooo/chrome/process/container"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestContainerRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port, err := findPort()
	if err != nil {
		t.Fatal(err)
	}

	c, err := container.New(ctx, &container.Settings{
		Image:  "yukinying/chrome-headless-browser:latest",
		Expose: []string{strconv.Itoa(port) + ":9222"},
	})
	if err != nil {
		t.Fatal(err)
	}

	addr, err := url.Parse("http://localhost:" + strconv.Itoa(port))
	if err != nil {
		t.Fatal(err)
	}

	// check if we can connect to the page
	ctxTimeout, cancelTimeout := context.WithTimeout(ctx, 2*time.Second)
	defer cancelTimeout()
	if e := ready(ctxTimeout, addr); e != nil {
		t.Fatal(e)
	}

	cancel()
	c.Wait()
}

func TestContainerCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port, err := findPort()
	if err != nil {
		t.Fatal(err)
	}

	c, err := container.New(ctx, &container.Settings{
		Image:  "yukinying/chrome-headless-browser:latest",
		Expose: []string{strconv.Itoa(port) + ":9222"},
	})
	if err != nil {
		t.Fatal(err)
	}

	cancel()
	assert.EqualError(t, c.Wait(), "container exited with error code: 137")
}

// get an available port
func findPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

func ready(ctx context.Context, addr *url.URL) error {
	b := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	for {
		if e := getURL(ctx, addr.String()); e == nil {
			return nil
		}

		sleep := b.NextBackOff()
		if sleep == backoff.Stop {
			return errors.New("connection timeout")
		}

		time.Sleep(sleep)
	}
}

// get the URL to connect to without actually connecting
func getURL(ctx context.Context, url string) (err error) {
	devt := devtool.New(url)
	target, err := devt.Get(ctx, devtool.Page)

	if err != nil {
		return errors.Wrap(err, "unable to create the target")
	} else if target.WebSocketDebuggerURL == "" {
		return errors.Wrap(err, "no websocket debugger url yet")
	}

	return nil
}
