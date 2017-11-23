package websocket

import "github.com/matthewmueller/golly/js"

// js:"WebSocket,omit"
type WebSocket struct {
	window.EventTarget
}

// Close
func (*WebSocket) Close(code *uint8, reason *string) {
	js.Rewrite("$<.Close($1, $2)", code, reason)
}

// Send
func (*WebSocket) Send(data interface{}) {
	js.Rewrite("$<.Send($1)", data)
}

// BinaryType
func (*WebSocket) BinaryType() (binaryType string) {
	js.Rewrite("$<.BinaryType")
	return binaryType
}

// BinaryType
func (*WebSocket) SetBinaryType(binaryType string) {
	js.Rewrite("$<.BinaryType = $1", binaryType)
}

// BufferedAmount
func (*WebSocket) BufferedAmount() (bufferedAmount uint) {
	js.Rewrite("$<.BufferedAmount")
	return bufferedAmount
}

// Extensions
func (*WebSocket) Extensions() (extensions string) {
	js.Rewrite("$<.Extensions")
	return extensions
}

// Onclose
func (*WebSocket) Onclose() (onclose func(*closeevent.CloseEvent)) {
	js.Rewrite("$<.Onclose")
	return onclose
}

// Onclose
func (*WebSocket) SetOnclose(onclose func(*closeevent.CloseEvent)) {
	js.Rewrite("$<.Onclose = $1", onclose)
}

// Onerror
func (*WebSocket) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*WebSocket) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onmessage
func (*WebSocket) Onmessage() (onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.Onmessage")
	return onmessage
}

// Onmessage
func (*WebSocket) SetOnmessage(onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.Onmessage = $1", onmessage)
}

// Onopen
func (*WebSocket) Onopen() (onopen func(window.Event)) {
	js.Rewrite("$<.Onopen")
	return onopen
}

// Onopen
func (*WebSocket) SetOnopen(onopen func(window.Event)) {
	js.Rewrite("$<.Onopen = $1", onopen)
}

// Protocol
func (*WebSocket) Protocol() (protocol string) {
	js.Rewrite("$<.Protocol")
	return protocol
}

// ReadyState
func (*WebSocket) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// URL
func (*WebSocket) URL() (url string) {
	js.Rewrite("$<.URL")
	return url
}
