package chrome

import (
	"context"
	"time"

	tomb "gopkg.in/tomb.v2"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/mafredri/cdp/devtool"
	"github.com/matthewmueller/golly/chrome/websocket"
	"github.com/pkg/errors"
)

// assert interface compliance with websocket.Websocket.
var _ websocket.Websocket = (*Target)(nil)

// Target is a thin wrapper around the websocket
// that ensures we clean up after our targets
type Target struct {
	addr   string
	ws     websocket.Websocket
	t      tomb.Tomb
	target *devtool.Target
}

func newTarget(parent context.Context, addr string) (*Target, error) {
	tar := &Target{
		addr: addr,
	}
	ctx := tar.t.Context(parent)

	devt := devtool.New(addr)
	target, err := devt.Create(ctx)
	if err != nil {
		return nil, err
	}
	tar.target = target

	if target.WebSocketDebuggerURL == "" {
		return nil, errors.New("target websocket url not ready - TODO: implement retry")
	}

	// dial the websocket
	ws, _, err := websocket.Dial(ctx, target.WebSocketDebuggerURL)
	if err != nil {
		return nil, errors.Wrap(err, "unable to dial websocket")
	}
	tar.ws = ws

	// start dying if the websocket goes
	tar.t.Go(ws.Wait)

	// if we start dying from above
	// try to close self gracefully
	tar.t.Go(func() error {
		<-ctx.Done()
		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return tar.close(c, "target is dying")
	})

	return tar, nil
}

// Write proxies to websocket
func (tar *Target) Write(ctx context.Context, method string, params interface{}) error {
	return tar.ws.Write(ctx, method, params)
}

// Send proxies to websocket
func (tar *Target) Send(ctx context.Context, method string, params interface{}, reply interface{}) error {
	return tar.ws.Send(ctx, method, params, reply)
}

// Subscribe proxies to websocket
func (tar *Target) Subscribe(ctx context.Context, method string) (*websocket.Listener, error) {
	return tar.ws.Subscribe(ctx, method)
}

// Wait proxies to websocket
func (tar *Target) Wait() error {
	return tar.t.Wait()
}

// Close closes the websocket and shuts down the target
func (tar *Target) Close() (err error) {
	// intentionally kill
	tar.t.Kill(nil)

	select {
	case <-tar.t.Dead():
		// if e := tar.t.Err(); e != nil {
		// 	log.WithError(e).Warnf("error from intentionally closing")
		// }
		return nil
	case <-time.After(5 * time.Second):
		return errors.New("couldn't close chrome after 5 seconds")
	}
}

func (tar *Target) close(ctx context.Context, reason string) (err error) {
	// try closing the websocket gracefully
	if e := tar.ws.Close(); e != nil {
		err = multierror.Append(err, e)
	}

	devt := devtool.New(tar.addr)

	// ignore context as we don't really want to cancel
	// this request either way
	if e := devt.Close(ctx, tar.target); e != nil {
		err = multierror.Append(err, e)
	}

	return err
}
