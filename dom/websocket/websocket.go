package websocket

import (
	"github.com/matthewmueller/golly/dom/closeevent"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(url string, protocols *interface{}) *WebSocket {
	js.Rewrite("WebSocket")
	return &WebSocket{}
}

// WebSocket struct
// js:"WebSocket,omit"
type WebSocket struct {
	window.EventTarget
}

// Close fn
func (*WebSocket) Close(code *uint8, reason *string) {
	js.Rewrite("$<.close($1, $2)", code, reason)
}

// Send fn
func (*WebSocket) Send(data interface{}) {
	js.Rewrite("$<.send($1)", data)
}

// BinaryType prop
func (*WebSocket) BinaryType() (binaryType string) {
	js.Rewrite("$<.binaryType")
	return binaryType
}

// BinaryType prop
func (*WebSocket) SetBinaryType(binaryType string) {
	js.Rewrite("$<.binaryType = $1", binaryType)
}

// BufferedAmount prop
func (*WebSocket) BufferedAmount() (bufferedAmount uint) {
	js.Rewrite("$<.bufferedAmount")
	return bufferedAmount
}

// Extensions prop
func (*WebSocket) Extensions() (extensions string) {
	js.Rewrite("$<.extensions")
	return extensions
}

// Onclose prop
func (*WebSocket) Onclose() (onclose func(*closeevent.CloseEvent)) {
	js.Rewrite("$<.onclose")
	return onclose
}

// Onclose prop
func (*WebSocket) SetOnclose(onclose func(*closeevent.CloseEvent)) {
	js.Rewrite("$<.onclose = $1", onclose)
}

// Onerror prop
func (*WebSocket) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*WebSocket) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onmessage prop
func (*WebSocket) Onmessage() (onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.onmessage")
	return onmessage
}

// Onmessage prop
func (*WebSocket) SetOnmessage(onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.onmessage = $1", onmessage)
}

// Onopen prop
func (*WebSocket) Onopen() (onopen func(window.Event)) {
	js.Rewrite("$<.onopen")
	return onopen
}

// Onopen prop
func (*WebSocket) SetOnopen(onopen func(window.Event)) {
	js.Rewrite("$<.onopen = $1", onopen)
}

// Protocol prop
func (*WebSocket) Protocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

// ReadyState prop
func (*WebSocket) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// URL prop
func (*WebSocket) URL() (url string) {
	js.Rewrite("$<.url")
	return url
}
