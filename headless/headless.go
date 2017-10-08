package headless

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/apex/log"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/rpcc"

	"github.com/matthewmueller/golly/headless/process"
)

// Headless struct
type Headless struct {
	process *process.Process
	ws      *rpcc.Conn
}

// New headless process
func New(chromePath string) (*Headless, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// find an open port
	port, err := freePort()
	if err != nil {
		return nil, err
	}

	p, e := process.New(ctx, &process.Settings{
		Path: chromePath,
		Flags: []string{
			"--headless",
			"--remote-debugging-port=" + strconv.Itoa(port),
			"--disable-gpu",
		},
	})
	if e != nil {
		return nil, e
	}

	time.Sleep(2 * time.Second)

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New("http://127.0.0.1:" + strconv.Itoa(port))
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Initiate a new RPC connection to the Chrome Debugging Protocol target.
	ws, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return nil, err
	}

	return &Headless{
		process: p,
		ws:      ws,
	}, nil
}

// Wait until the process dies
func (h *Headless) Wait() error {
	return h.process.Wait()
}

// Run fn
func (h *Headless) Run(ctx context.Context, expression string) (string, error) {
	awaitPromise := true

	c := cdp.NewClient(h.ws)

	e := c.Page.Enable(ctx)
	if e != nil {
		log.Infof("page enable")
		return "", e
	}

	e = c.Runtime.Enable(ctx)
	if e != nil {
		return "", e
	}

	log.Infof("expression %s", expression)

	res, e := c.Runtime.Evaluate(ctx, &runtime.EvaluateArgs{
		AwaitPromise: &awaitPromise,
		Expression:   expression,
	})
	if e != nil {
		return "", e
	}

	return res.Result.String(), nil
}

// Close our headless connection
func (h *Headless) Close(timeout time.Duration) error {
	h.ws.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return h.process.Stop(ctx)
}

// get an available port
func freePort() (int, error) {
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
