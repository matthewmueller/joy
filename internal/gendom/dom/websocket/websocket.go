package websocket

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/closeevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/messageevent"
	"github.com/matthewmueller/golly/js"
)

type WebSocket struct {
	*eventtarget.EventTarget
}

func (*WebSocket) Close(code uint8, reason string) {
	js.Rewrite("$<.close($1, $2)", code, reason)
}

func (*WebSocket) Send(data interface{}) {
	js.Rewrite("$<.send($1)", data)
}

func (*WebSocket) GetBinaryType() (binaryType string) {
	js.Rewrite("$<.binaryType")
	return binaryType
}

func (*WebSocket) SetBinaryType(binaryType string) {
	js.Rewrite("$<.binaryType = $1", binaryType)
}

func (*WebSocket) GetBufferedAmount() (bufferedAmount uint) {
	js.Rewrite("$<.bufferedAmount")
	return bufferedAmount
}

func (*WebSocket) GetExtensions() (extensions string) {
	js.Rewrite("$<.extensions")
	return extensions
}

func (*WebSocket) GetOnclose() (cls *closeevent.CloseEvent) {
	js.Rewrite("$<.onclose")
	return cls
}

func (*WebSocket) SetOnclose(cls *closeevent.CloseEvent) {
	js.Rewrite("$<.onclose = $1", cls)
}

func (*WebSocket) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*WebSocket) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*WebSocket) GetOnmessage() (message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*WebSocket) SetOnmessage(message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

func (*WebSocket) GetOnopen() (open *event.Event) {
	js.Rewrite("$<.onopen")
	return open
}

func (*WebSocket) SetOnopen(open *event.Event) {
	js.Rewrite("$<.onopen = $1", open)
}

func (*WebSocket) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*WebSocket) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*WebSocket) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}
