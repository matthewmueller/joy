package websocket

import (
	"context"
	"encoding/json"
	"math/rand"
	"net"
	"net/http"
	"time"

	tomb "gopkg.in/tomb.v2"

	alog "github.com/apex/log"
	"github.com/cenkalti/backoff"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var log = alog.WithField("pkg", "websocket")

// Listener struct
type Listener struct {
	method     string
	unregister chan *listenerRequest
	eventCh    chan json.RawMessage
	t          *tomb.Tomb
}

// Websocket interface
type Websocket interface {
	Write(ctx context.Context, method string, params interface{}) error
	Send(ctx context.Context, method string, params interface{}, reply interface{}) error
	Subscribe(ctx context.Context, method string) (*Listener, error)
	Wait() error
	Close() error
}

// ErrClosed happens when we explicitly close the connection
var ErrClosed = errors.New("connection has been closed")

// ErrSendTimeout error while waiting for a response
var ErrSendTimeout = errors.New("timeout while waiting for a response")

// ErrReadTimeout error while waiting for a response
var ErrReadTimeout = errors.New("timeout while trying to read a response")

type writeRequest struct {
	Type     int
	JSON     interface{}
	Binary   []byte
	Response chan error
	Deadline time.Time
}

type listenerRequest struct {
	Method  string `json:"method,omitempty"`
	EventCh chan json.RawMessage
}

type sendListener struct {
	ID         int
	ResponseCh chan json.RawMessage
}

type message struct {
	ID     int         `json:"id,omitempty"`
	Method string      `json:"method,omitempty"`
	Params interface{} `json:"params,omitempty"`
}

type response struct {
	ID     *int            `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Method string          `json:"method,omitempty"`
	Params json.RawMessage `json:"params,omitempty"`
	Error  *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// WS struct
type WS struct {
	url    string
	dialer *websocket.Dialer
	conn   *websocket.Conn

	readCh   chan json.RawMessage
	writeCh  chan *writeRequest
	closedCh chan bool

	registerListener   chan *listenerRequest
	unregisterListener chan *listenerRequest
	sendCh             chan *sendListener

	t tomb.Tomb
}

// Dial the websocket
func Dial(parent context.Context, url string) (Websocket, *http.Response, error) {
	dialer := &websocket.Dialer{
		// 128kb, NOTE: this fixes the issue with
		// chrome headless terminating after
		// 3903 bytes. 128kb is the framesize
		// of the browser by default.
		WriteBufferSize:  128000,
		ReadBufferSize:   128000,
		HandshakeTimeout: 2 * time.Second,
		// give us context support
		NetDial: func(network, addr string) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(parent, network, addr)
		},
	}

	conn, res, err := dial(parent, dialer, url)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to connect to websocket")
	}

	ws := WS{
		conn: conn,

		readCh:  make(chan json.RawMessage),
		writeCh: make(chan *writeRequest),

		sendCh:             make(chan *sendListener, 1),
		registerListener:   make(chan *listenerRequest, 1),
		unregisterListener: make(chan *listenerRequest, 1),
	}

	// we'll still want to cancel these out
	// if an error comes from above
	ctx := ws.t.Context(parent)

	// start the read loop for incoming messages
	ws.t.Go(func() error {
		return ws.readLoop(ctx)
	})

	// start the write loop for outgoing messages
	ws.t.Go(func() error {
		return ws.writeLoop(ctx)
	})

	// setup the event loop (for coordination)
	ws.t.Go(func() error {
		return ws.eventLoop(ctx)
	})

	// close the connection once context
	// is cancelled or we start dying
	ws.t.Go(func() error {
		<-ctx.Done()
		return ws.conn.Close()
	})

	return &ws, res, nil
}

// read loop
func (ws *WS) readLoop(ctx context.Context) error {
	conn := ws.conn
	for {
		var raw json.RawMessage
		if e := conn.ReadJSON(&raw); e != nil {
			return errors.Wrap(e, "read error")
		}

		select {
		case ws.readCh <- raw:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// write loop
func (ws *WS) writeLoop(ctx context.Context) error {
	conn := ws.conn
	for {
		select {
		case outgoing := <-ws.writeCh:
			var err error
			switch outgoing.Type {
			case websocket.TextMessage:
				conn.SetWriteDeadline(outgoing.Deadline)
				err = conn.WriteJSON(outgoing.JSON)
			case websocket.BinaryMessage:
				conn.SetWriteDeadline(outgoing.Deadline)
				err = conn.WriteMessage(websocket.BinaryMessage, outgoing.Binary)
			case websocket.CloseMessage:
				err = conn.WriteControl(websocket.CloseMessage, outgoing.Binary, outgoing.Deadline)
			default:
				err = errors.New("write does not yet support that type")
			}
			if err != nil {
				outgoing.Response <- err
				return err
			}
			// all good
			close(outgoing.Response)
		case <-ctx.Done():
			// if we've closed end the loop
			// log.Info("ending write loop")
			return ctx.Err()
		}
	}
}

func (ws *WS) eventLoop(ctx context.Context) error {
	listeners := map[string]chan json.RawMessage{}
	senders := map[int]*sendListener{}

	for {
		select {
		case msg := <-ws.readCh:
			var res response
			if e := json.Unmarshal(msg, &res); e != nil {
				return e
			}

			// check if it has an ID
			if res.ID != nil && senders[*res.ID] != nil {
				// send to root
				// log.Warnf("sending to root")
				senders[*res.ID].ResponseCh <- res.Result
				// log.Warnf("sent to root")
				delete(senders, *res.ID)
				continue
			}

			if listeners[res.Method] != nil {
				// log.Warnf("writing to listener %s", res.Method)
				select {
				case listeners[res.Method] <- msg:
					log.Debugf("wrote to listener %s", res.Method)
					continue
				case <-ctx.Done():
					// log.Infof("ending event loop")
					return ctx.Err()
				}
			}

			// log.
			// 	WithField("message", string(msg)).
			// 	Warnf("ignoring read message with no sender or listener")

		case listener := <-ws.registerListener:
			listeners[listener.Method] = listener.EventCh
		case listener := <-ws.unregisterListener:
			delete(listeners, listener.Method)
		case sender := <-ws.sendCh:
			senders[sender.ID] = sender
		case <-ctx.Done():
			// log.Infof("ending event loop")
			return ctx.Err()
		}
	}
}

// Write to the websocket
func (ws *WS) Write(ctx context.Context, method string, params interface{}) error {
	responseErrorCh := make(chan error, 1)

	var deadline time.Time
	if d, ok := ctx.Deadline(); ok {
		deadline = d
	}

	return ws.write(ctx, &writeRequest{
		Type: websocket.TextMessage,
		JSON: &message{
			ID:     int(rand.Int31()),
			Method: method,
			Params: params,
		},
		Response: responseErrorCh,
		Deadline: deadline,
	})
}

// WriteBinary to the websocket
func (ws *WS) WriteBinary(ctx context.Context, binary []byte) error {
	responseErrorCh := make(chan error, 1)

	var deadline time.Time
	if d, ok := ctx.Deadline(); ok {
		deadline = d
	}

	return ws.write(ctx, &writeRequest{
		Type:     websocket.BinaryMessage,
		Binary:   binary,
		Response: responseErrorCh,
		Deadline: deadline,
	})
}

// WriteControl to the websocket
func (ws *WS) WriteControl(ctx context.Context, messageType int, binary []byte) error {
	responseErrorCh := make(chan error, 1)

	var deadline time.Time
	if d, ok := ctx.Deadline(); ok {
		deadline = d
	}

	return ws.write(ctx, &writeRequest{
		Type:     messageType,
		Binary:   binary,
		Response: responseErrorCh,
		Deadline: deadline,
	})
}

// Send a message on the websocket and wait for a response
func (ws *WS) Send(ctx context.Context, method string, params interface{}, response interface{}) error {
	responseCh := make(chan json.RawMessage)
	responseErrorCh := make(chan error, 1)
	id := int(rand.Int31())

	listener := &sendListener{
		ID:         id,
		ResponseCh: responseCh,
	}

	select {
	case ws.sendCh <- listener:
	// websocket failure
	case <-ws.t.Dying():
		return ws.t.Err()
	// current send error
	case <-ctx.Done():
		return ctx.Err()
	}

	var deadline time.Time
	if d, ok := ctx.Deadline(); ok {
		deadline = d
	}

	if e := ws.write(ctx, &writeRequest{
		Type: websocket.TextMessage,
		JSON: &message{
			ID:     id,
			Method: method,
			Params: params,
		},
		Response: responseErrorCh,
		Deadline: deadline,
	}); e != nil {
		return e
	}

	// wait for a response
	select {
	case msg := <-responseCh:
		if response == nil {
			return nil
		}
		return json.Unmarshal(msg, response)
	// websocket error
	case <-ws.t.Dying():
		return ws.t.Err()
		// send error
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (ws *WS) write(ctx context.Context, req *writeRequest) error {
	// write request
	select {
	case ws.writeCh <- req:
	case <-ws.t.Dying():
		return ws.t.Err()
	case <-ctx.Done():
		return ctx.Err()
	}

	// ensure we wrote to the socket
	select {
	case err := <-req.Response:
		return err
	case <-ws.t.Dying():
		return ws.t.Err()
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Wait until the websocket errors out
func (ws *WS) Wait() error {
	return ws.t.Wait()
}

// Close the websocket gracefully
func (ws *WS) Close() error {
	// if we're already dying, wait until we exit
	select {
	case <-ws.t.Dying():
		return ws.t.Wait()
	default:
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// try to close gracefully give it 5 seconds
	if e := ws.WriteControl(ctx, websocket.CloseMessage, []byte("shutting down gracefully")); e != nil {
		ws.t.Kill(e)
	} else {
		ws.t.Kill(nil)
	}

	return ws.t.Wait()
}

// Subscribe to events from a method
func (ws *WS) Subscribe(ctx context.Context, method string) (*Listener, error) {
	// TODO: re-evaluate if we need this 50 here
	eventCh := make(chan json.RawMessage, 50)
	listener := &listenerRequest{
		Method:  method,
		EventCh: eventCh,
	}

	select {
	case ws.registerListener <- listener:
	case <-ws.t.Dying():
		return nil, ws.t.Err()
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return &Listener{
		unregister: ws.unregisterListener,
		method:     method,
		eventCh:    eventCh,
		t:          &ws.t,
	}, nil
}

// Wait for a new event
func (l *Listener) Wait(v interface{}) error {
	select {
	case raw := <-l.eventCh:
		var res response
		if e := json.Unmarshal(raw, &res); e != nil {
			return e
		}
		return json.Unmarshal(res.Params, v)
	case <-l.t.Dying():
		return l.t.Err()
	}
}

// Close the listener
func (l *Listener) Close() error {
	listener := &listenerRequest{
		Method:  l.method,
		EventCh: l.eventCh,
	}

	select {
	case l.unregister <- listener:
		return nil
	case <-l.t.Dying():
		return l.t.Err()
	}
}

func dial(ctx context.Context, dialer *websocket.Dialer, url string) (*websocket.Conn, *http.Response, error) {
	b := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	for {
		conn, res, err := dialer.Dial(url, nil)
		if err == nil {
			return conn, res, err
		}

		sleep := b.NextBackOff()
		if sleep == backoff.Stop {
			return nil, nil, errors.New("dial timed out")
		}

		time.Sleep(sleep)
	}
}
