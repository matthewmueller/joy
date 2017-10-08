package chrome

import (
	"context"
	"net"
	"net/url"
	"strconv"
	"time"

	"gopkg.in/tomb.v2"

	"github.com/cenkalti/backoff"
	"github.com/mafredri/cdp/devtool"
	"github.com/matthewmueller/golly/chrome/process"
	"github.com/matthewmueller/golly/chrome/process/container"
	"github.com/matthewmueller/golly/chrome/process/local"
	"github.com/matthewmueller/golly/chrome/websocket"
	"github.com/pkg/errors"
)

// Chrome struct
type Chrome struct {
	settings *Settings
	addr     url.URL
	p        process.Process
	t        tomb.Tomb
}

// Settings for chrome
type Settings struct {
	Image          string
	ExecutablePath string
	Addr           string
	Stdout         chan<- []byte
	Stderr         chan<- []byte
}

// ErrConnectionTimeout timeout while trying to connect
var ErrConnectionTimeout = errors.New("connection to the chrome instance timed out")

// New Chrome
//
// Chrome owns the process because it needs
// to be able to kill it if things go awry
func New(parent context.Context, settings *Settings) (*Chrome, error) {
	// resolve the address
	addr, err := resolveAddr(settings.Addr)
	if err != nil {
		return nil, err
	}

	ch := &Chrome{
		settings: settings,
		addr:     addr,
	}

	// if we're dying, cancel downstream
	ctx := ch.t.Context(parent)

	// start the process
	p, err := startProcess(ctx, settings, addr)
	if err != nil {
		return nil, err
	}
	ch.p = p

	// wait until chrome is ready
	if e := ready(ctx, addr); e != nil {
		return nil, e
	}

	// handle downstream process dying
	// in the future
	ch.t.Go(p.Wait)

	return ch, nil
}

// Lease a target
// TODO: test out multiple targets at some point again
func (ch *Chrome) Lease(parent context.Context) (websocket.Websocket, error) {
	// also cancel context if chrome is dying
	ctx := ch.t.Context(parent)
	// TODO: right now dead targets don't kill chrome,
	// but perhaps they should if target.Wait()
	// would only return unexpected errors
	return newTarget(ctx, ch.addr.String())
}

// Wait until chrome exits
func (ch *Chrome) Wait() error {
	return ch.t.Wait()
}

// Close chrome
func (ch *Chrome) Close() error {
	return ch.p.Stop()
}

// resolves an chrome devtools address
func resolveAddr(address string) (u url.URL, err error) {
	var port int

	hostname := "127.0.0.1"
	defaultURL := url.URL{
		Scheme: "http",
		Host:   "127.0.0.1:9222",
	}

	if address == "" {
		port, err = findPort()
		if err != nil {
			return u, errors.Wrap(err, "unable to find an available port")
		}
		defaultURL.Host = hostname + ":" + strconv.Itoa(port)
		return defaultURL, nil
	}

	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return u, errors.Wrap(err, "did not understand the address")
	}

	if addr.IP != nil {
		hostname = addr.IP.String()
	}

	if addr.Port == 0 {
		port, err = findPort()
		if err != nil {
			return u, errors.Wrap(err, "unable to find an available port")
		}
	} else {
		port = addr.Port
	}

	defaultURL.Host = hostname + ":" + strconv.Itoa(port)
	return defaultURL, nil
}

func startProcess(ctx context.Context, settings *Settings, addr url.URL) (process.Process, error) {
	if settings.ExecutablePath == "" && settings.Image == "" {
		return nil, errors.New("need either an executable path or an image")
	}

	// local process
	if settings.ExecutablePath != "" {

		type Settings struct {
			Name   string
			Flags  []string
			Stdout chan<- []byte
			Stderr chan<- []byte
		}

		return local.New(ctx, &local.Settings{
			Name: settings.ExecutablePath,
			Flags: []string{
				"--headless",
				"--remote-debugging-port=" + addr.Port(),
				"--disable-gpu",
			},
			Stdout: settings.Stdout,
			Stderr: settings.Stderr,
		})
	}

	return container.New(ctx, &container.Settings{
		Image:  settings.Image,
		Expose: []string{addr.Port() + ":9222"},
		Stdout: settings.Stdout,
		Stderr: settings.Stderr,
	})
}

func ready(ctx context.Context, addr url.URL) error {
	b := backo(ctx)
	for {
		if e := getURL(ctx, addr.String()); e == nil {
			return nil
		}

		sleep := b.NextBackOff()
		if sleep == backoff.Stop {
			return ErrConnectionTimeout
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

func backo(ctx context.Context) backoff.BackOffContext {
	return backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
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
