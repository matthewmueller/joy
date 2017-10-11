package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/apex/log"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/rpcc"
)

type evaluate struct {
	AwaitPromise          bool   `json:"await_promise,omitempty"`
	Expression            string `json:"expression,omitempty"`
	GeneratePreview       bool   `json:"generatePreview,omitempty"`
	IncludeCommandLineAPI bool   `json:"includeCommandLineAPI,omitempty"`
}

type evaluateResult struct {
	Result           remoteObject    `json:"result,omitempty"`
	ExceptionDetails json.RawMessage `json:"exceptionDetails,omitempty"`
}

type console struct {
	Type string         `json:"type,omitempty"`
	Args []remoteObject `json:"args,omitempty"`
}

type remoteObject struct {
	ObjectID   string          `json:"objectId,omitempty"`
	Type       string          `json:"type,omitempty"`
	Subtype    string          `json:"subtype,omitempty"`
	Value      json.RawMessage `json:"value,omitempty"`
	Preview    *remoteObject   `json:"preview,omitempty"`
	Properties []*remoteObject `json:"properties,omitempty"`
}

type awaitPromise struct {
	PromiseObjectID int  `json:"promiseObjectId,omitempty"`
	ReturnByValue   bool `json:"returnByValue,omitempty"`
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

var source = `
new Promise(function (resolve, reject) {
	setTimeout(() => resolve("hi world!"), 5000)
})
`

func main() {
	ctx := context.Background()

	// stdout := make(chan []byte, 100)
	// stderr := make(chan []byte, 100)

	// go func() {
	// 	for {
	// 		line := <-stdout
	// 		fmt.Println(string(line))
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		line := <-stderr
	// 		fmt.Println(string(line))
	// 	}
	// }()

	// ch, err := chrome.New(ctx, &chrome.Settings{
	// 	ExecutablePath: os.Getenv("GOLLY_CHROME_PATH"),
	// 	Stderr:         stderr,
	// 	Stdout:         stdout,
	// })
	// if err != nil {
	// 	log.WithError(err).Fatal("error creating new headless")
	// }
	// defer ch.Close()

	// ws, err := ch.Lease(ctx)
	// if err != nil {
	// 	log.WithError(err).Fatal("error leasing target")
	// }
	// defer ws.Close()

	// if e := ws.Send(ctx, "Runtime.enable", nil, nil); e != nil {
	// 	log.WithError(e).Fatal("error running runtime.enable")
	// }

	// listener, err := ws.Subscribe(ctx, "Runtime.consoleAPICalled")
	// if err != nil {
	// 	log.WithError(err).Fatal("error subscribing")
	// }
	// defer listener.Close()

	// go func() {
	// 	for {
	// 		var msg console
	// 		e := listener.Wait(&msg)
	// 		if e != nil {
	// 			return
	// 		}
	// 		for _, arg := range msg.Args {
	// 			log.Infof("got arg %+v", arg)
	// 		}
	// 	}
	// }()

	// var obj json.RawMessage
	// start := time.Now()
	// if e := ws.Send(ctx, "Runtime.evaluate", &evaluate{
	// 	Expression:            source,
	// 	GeneratePreview:       true,
	// 	AwaitPromise:          true,
	// 	IncludeCommandLineAPI: true,
	// }, &obj); e != nil {
	// 	log.WithError(e).Fatal("error evaluating")
	// }
	// log.Infof("done in %s %+v", time.Since(start), string(obj))

	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			log.WithError(err).Errorf("conn")
			return
		}
	}

	// Initiate a new RPC connection to the Chrome Debugging Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		log.WithError(err).Errorf("dial")
		return
	}
	defer conn.Close() // Leaving connections open will leak memory.

	c := cdp.NewClient(conn)

	if e := c.Runtime.Enable(ctx); e != nil {
		log.WithError(e).Errorf("enable error")
		return
	}

	console, err := c.Runtime.ConsoleAPICalled(ctx)
	if err != nil {
		log.WithError(err).Errorf("console")
		return
	}
	defer console.Close()

	go func() {
		for {
			msg := <-console.Ready()
			log.Infof("%+v", msg)
		}
	}()

	awaitPromise := true
	generatePreview := true
	start := time.Now()
	res, err := c.Runtime.Evaluate(ctx, &runtime.EvaluateArgs{
		Expression:      source,
		AwaitPromise:    &awaitPromise,
		GeneratePreview: &generatePreview,
	})
	if err != nil {
		log.WithError(err).Errorf("err eval")
	}
	log.Infof("result %s after %s", res.Result.String(), time.Since(start))
}
