package websocket

import (
	"github.com/matthewmueller/joy/dom/closeevent"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*WebSocket)(nil)

// New fn
func New(url string, protocols *interface{}) *WebSocket {
	js.Rewrite("WebSocket")
	return &WebSocket{}
}

// WebSocket struct
// js:"WebSocket,omit"
type WebSocket struct {
}

// Close fn
// js:"close"
func (*WebSocket) Close(code *uint8, reason *string) {
	js.Rewrite("$_.close($1, $2)", code, reason)
}

// Send fn
// js:"send"
func (*WebSocket) Send(data interface{}) {
	js.Rewrite("$_.send($1)", data)
}

// AddEventListener fn
// js:"addEventListener"
func (*WebSocket) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*WebSocket) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*WebSocket) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// BinaryType prop
// js:"binaryType"
func (*WebSocket) BinaryType() (binaryType string) {
	js.Rewrite("$_.binaryType")
	return binaryType
}

// SetBinaryType prop
// js:"binaryType"
func (*WebSocket) SetBinaryType(binaryType string) {
	js.Rewrite("$_.binaryType = $1", binaryType)
}

// BufferedAmount prop
// js:"bufferedAmount"
func (*WebSocket) BufferedAmount() (bufferedAmount uint) {
	js.Rewrite("$_.bufferedAmount")
	return bufferedAmount
}

// Extensions prop
// js:"extensions"
func (*WebSocket) Extensions() (extensions string) {
	js.Rewrite("$_.extensions")
	return extensions
}

// Onclose prop
// js:"onclose"
func (*WebSocket) Onclose() (onclose func(*closeevent.CloseEvent)) {
	js.Rewrite("$_.onclose")
	return onclose
}

// SetOnclose prop
// js:"onclose"
func (*WebSocket) SetOnclose(onclose func(*closeevent.CloseEvent)) {
	js.Rewrite("$_.onclose = $1", onclose)
}

// Onerror prop
// js:"onerror"
func (*WebSocket) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*WebSocket) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onmessage prop
// js:"onmessage"
func (*WebSocket) Onmessage() (onmessage func(*window.MessageEvent)) {
	js.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*WebSocket) SetOnmessage(onmessage func(*window.MessageEvent)) {
	js.Rewrite("$_.onmessage = $1", onmessage)
}

// Onopen prop
// js:"onopen"
func (*WebSocket) Onopen() (onopen func(window.Event)) {
	js.Rewrite("$_.onopen")
	return onopen
}

// SetOnopen prop
// js:"onopen"
func (*WebSocket) SetOnopen(onopen func(window.Event)) {
	js.Rewrite("$_.onopen = $1", onopen)
}

// Protocol prop
// js:"protocol"
func (*WebSocket) Protocol() (protocol string) {
	js.Rewrite("$_.protocol")
	return protocol
}

// ReadyState prop
// js:"readyState"
func (*WebSocket) ReadyState() (readyState uint8) {
	js.Rewrite("$_.readyState")
	return readyState
}

// URL prop
// js:"url"
func (*WebSocket) URL() (url string) {
	js.Rewrite("$_.url")
	return url
}
