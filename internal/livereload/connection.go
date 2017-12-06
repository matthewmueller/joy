package livereload

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/apex/log"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type cmd struct {
	Command   string   `json:"command"`
	Protocols []string `json:"protocols"`
}

var protocols = []string{
	"http://livereload.com/protocols/official-7",
	"http://livereload.com/protocols/official-8",
	"http://livereload.com/protocols/official-9",
	"http://livereload.com/protocols/2.x-origin-version-negotiation",
	"http://livereload.com/protocols/2.x-remote-control",
}

// Upgrade the connection
func upgrade(ctx context.Context, w http.ResponseWriter, r *http.Request) (*conn, error) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	eg, ctx := errgroup.WithContext(ctx)

	conn := &conn{
		conn:    c,
		actionc: make(chan func() (json.RawMessage, error)),
		ctx:     ctx,
		eg:      eg,
		ready:   false,
	}

	// read from the browser
	eg.Go(func() error {
		for {
			msgType, reader, err := c.NextReader()
			if err != nil {
				return err
			}

			// Close if binary instead of text
			if msgType == websocket.BinaryMessage {
				return errors.New("text message required recieved a binary message")
			}

			// Close if it's not JSON
			command := &cmd{}
			err = json.NewDecoder(reader).Decode(command)
			if err != nil {
				return errors.Wrapf(err, "error decoding command")
			}

			// ignore empty commands
			if command.Command == "" {
				return fmt.Errorf("server command shouldn't be empty")
			}

			// ensure we're ready to go
			if !conn.ready {
				if valid := validateHello(command); !valid {
					return fmt.Errorf("invalid handshake")
				}
				conn.ready = true
				log.Debugf("ready!")
			}
		}
	})

	// write to the browser
	eg.Go(func() error {
		for {
			select {
			case fn := <-conn.actionc:
				// if !conn.ready {
				// 	return fmt.Errorf("bad handshake")
				// }

				buf, err := fn()
				if err != nil {
					return errors.Wrapf(err, "error running fn")
				}

				log.Debugf("sending: %s", buf)
				c.SetWriteDeadline(time.Now().Add(5 * time.Second))
				if err := c.WriteJSON(buf); err != nil {
					return errors.Wrapf(err, "error writing JSON")
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	// kick us off
	conn.actionc <- func() (json.RawMessage, error) {
		return json.Marshal(struct {
			Command    string   `json:"command"`
			Protocols  []string `json:"protocols"`
			ServerName string   `json:"serverName"`
		}{
			Command:    "hello",
			Protocols:  protocols,
			ServerName: "joy",
		})
	}

	return conn, nil
}

// conn is a single websocket connection
type conn struct {
	actionc chan func() (json.RawMessage, error)
	conn    *websocket.Conn
	ctx     context.Context
	eg      *errgroup.Group
	ready   bool
}

// Reload a file
func (c *conn) Reload(file string) {
	c.actionc <- func() (json.RawMessage, error) {
		return json.Marshal(struct {
			Command string `json:"command,omitempty"`
			Path    string `json:"path,omitempty"`
			LiveCSS bool   `json:"live_css,omitempty"`
		}{
			Command: "reload",
			Path:    file,
			LiveCSS: true,
		})
	}
}

// Wait until the connection is over
func (c *conn) Wait() error {
	return c.eg.Wait()
}

// validate the client's handshake
func validateHello(cmd *cmd) bool {
	if cmd.Command != "hello" {
		return false
	}
	for _, c := range cmd.Protocols {
		for _, s := range protocols {
			if c == s {
				return true
			}
		}
	}
	return false
}
