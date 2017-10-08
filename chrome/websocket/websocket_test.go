package websocket_test

import (
	"context"
	"encoding/json"
	"net"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cenkalti/backoff"
	"github.com/mafredri/cdp/devtool"
	"github.com/matthewmueller/voodooo/chrome/process/container"
	"github.com/matthewmueller/voodooo/websocket"
	"github.com/pkg/errors"
)

func TestSend(t *testing.T) {
	// start a chrome
	ctx, cancel := context.WithCancel(context.Background())
	addr, wait, err := chrome(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer wait()
	defer cancel()

	// test websocket
	ws, _, err := websocket.Dial(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}

	// test send
	var raw json.RawMessage
	if e := ws.Send(ctx, "Page.getLayoutMetrics", nil, &raw); e != nil {
		t.Fatal(e)
	}
	assert.Contains(t, string(raw), "contentSize")

	// cleanup
	cancel()
	wait()
	assert.Equal(t, ws.Wait(), context.Canceled)
}

func TestWrite(t *testing.T) {
	// start a chrome
	ctx, cancel := context.WithCancel(context.Background())
	addr, wait, err := chrome(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer wait()
	defer cancel()

	// test websocket
	ws, _, err := websocket.Dial(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}

	// test send
	raw := json.RawMessage(`{"url":"https://google.com"}`)
	if e := ws.Write(ctx, "Page.navigate", raw); e != nil {
		t.Fatal(e)
	}

	// cleanup
	cancel()
	wait()
	assert.Equal(t, ws.Wait(), context.Canceled)
}

func TestSubscribe(t *testing.T) {
	// start a chrome
	ctx, cancel := context.WithCancel(context.Background())
	addr, wait, err := chrome(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer wait()
	defer cancel()

	// test websocket
	ws, _, err := websocket.Dial(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}

	if e := ws.Send(ctx, "Page.enable", nil, nil); e != nil {
		t.Fatal(e)
	}

	args := json.RawMessage(`{"format":"png","quality":80}`)
	if e := ws.Send(ctx, "Page.startScreencast", args, nil); e != nil {
		t.Fatal(e)
	}

	sub, err := ws.Subscribe(ctx, "Page.screencastFrame")
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Close()

	var frame struct {
		SessionID int             `json:"sessionId"`
		Metadata  json.RawMessage `json:"metadata"`
	}
	if e := sub.Wait(&frame); e != nil {
		t.Fatal(e)
	}
	assert.Contains(t, string(frame.Metadata), "deviceWidth")

	args = json.RawMessage(`{"sessionId":` + strconv.Itoa(frame.SessionID) + `}`)
	if e := ws.Send(ctx, "Page.screencastFrameAck", args, nil); e != nil {
		t.Fatal(e)
	}

	// cleanup
	cancel()
	wait()
	assert.Equal(t, ws.Wait(), context.Canceled)
}

func TestClose(t *testing.T) {
	// start a chrome
	ctx, cancel := context.WithCancel(context.Background())
	addr, wait, err := chrome(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer wait()
	defer cancel()

	// test websocket
	ws, _, err := websocket.Dial(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}

	// test send
	raw := json.RawMessage(`{"url":"https://google.com"}`)
	if e := ws.Write(ctx, "Page.navigate", raw); e != nil {
		t.Fatal(e)
	}

	// cleanup with close
	assert.Equal(t, ws.Close(), context.Canceled)
	assert.Equal(t, ws.Wait(), context.Canceled)
}

func chrome(ctx context.Context) (string, func(), error) {
	port, err := findPort()
	if err != nil {
		return "", nil, err
	}

	c, err := container.New(ctx, &container.Settings{
		Image:  "yukinying/chrome-headless-browser:latest",
		Expose: []string{strconv.Itoa(port) + ":9222"},
	})
	if err != nil {
		return "", nil, err
	}

	addr, err := url.Parse("http://localhost:" + strconv.Itoa(port))
	if err != nil {
		return "", nil, err
	}

	if e := ready(ctx, addr); e != nil {
		return "", nil, e
	}

	devt := devtool.New(addr.String())
	target, err := devt.Create(ctx)
	if err != nil {
		return "", nil, err
	}

	if target.WebSocketDebuggerURL == "" {
		return "", nil, errors.New("no websocket debugger url")
	}

	return target.WebSocketDebuggerURL, func() {
		c.Wait()
	}, nil
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
