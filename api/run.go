package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/matthewmueller/golly/chrome"
)

type evaluate struct {
	AwaitPromise    bool   `json:"await_promise,omitempty"`
	Expression      string `json:"expression,omitempty"`
	GeneratePreview bool   `json:"generatePreview,omitempty"`
}

type console struct {
	Type string         `json:"type,omitempty"`
	Args []remoteObject `json:"args,omitempty"`
}

type remoteObject struct {
	Type       string          `json:"type,omitempty"`
	Subtype    string          `json:"subtype,omitempty"`
	Value      json.RawMessage `json:"value,omitempty"`
	Preview    *remoteObject   `json:"preview,omitempty"`
	Properties []*remoteObject `json:"properties,omitempty"`
}

// type evaluate struct {
// 	AwaitPromise bool   `json:"await_promise,omitempty"`
// 	Expression   string `json:"expression,omitempty"`
// }

// Run fn
// TODO: pass in a writer
func Run(ctx context.Context, filepath string) error {
	// eg, ctx := errgroup.WithContext(ctx)

	chromePath := os.Getenv("GOLLY_CHROME_PATH")
	if chromePath == "" {
		return errors.New("GOLLY_CHROME_PATH not set")
	}

	files, err := Build(filepath)
	if err != nil {
		return err
	} else if len(files) != 1 {
		return errors.New("not 1 file")
	}

	stdout := make(chan []byte, 100)
	stderr := make(chan []byte, 100)

	go func() {
		for {
			line := <-stdout
			fmt.Println(line)
		}
	}()

	go func() {
		for {
			line := <-stderr
			fmt.Println(line)
		}
	}()

	ch, err := chrome.New(ctx, &chrome.Settings{
		ExecutablePath: chromePath,
		Stderr:         stderr,
		Stdout:         stdout,
	})
	if err != nil {
		return err
	}
	defer ch.Close()

	ws, err := ch.Lease(ctx)
	if err != nil {
		return err
	}
	defer ws.Close()

	if e := ws.Send(ctx, "Runtime.enable", nil, nil); e != nil {
		return e
	}

	listener, err := ws.Subscribe(ctx, "Runtime.consoleAPICalled")
	if err != nil {
		return err
	}
	defer listener.Close()

	go func() {
		for {
			var msg console
			e := listener.Wait(&msg)
			if e != nil {
				return
			}
			for _, arg := range msg.Args {
				fmt.Println(formatValue(&arg))
			}
		}
	}()

	var res json.RawMessage
	if e := ws.Send(ctx, "Runtime.evaluate", &evaluate{
		Expression:      files[0].Source,
		GeneratePreview: true,
		AwaitPromise:    true,
	}, &res); e != nil {
		return e
	}

	return nil
}

func formatValue(obj *remoteObject) string {
	preview := obj.Preview
	value := obj.Value

	switch obj.Type {
	case "object":
		switch obj.Subtype {
		case "array":
			var arr []string
			for _, prop := range preview.Properties {
				arr = append(arr, formatValue(prop))
			}
			return "[ " + strings.Join(arr, ", ") + " ]"
		default:
			bytes, _ := json.Marshal(preview)
			return string(bytes)
		}
	case "string":
		return string(value[1 : len(value)-1])
	default:
		return string(value)
	}
}
