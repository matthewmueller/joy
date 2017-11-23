package webglcontextevent

import (
	"github.com/matthewmueller/golly/dom2/webglcontexteventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *webglcontexteventinit.WebGLContextEventInit) *WebGLContextEvent {
	js.Rewrite("WebGLContextEvent")
	return &WebGLContextEvent{}
}

// WebGLContextEvent struct
// js:"WebGLContextEvent,omit"
type WebGLContextEvent struct {
	window.Event
}

// StatusMessage prop
func (*WebGLContextEvent) StatusMessage() (statusMessage string) {
	js.Rewrite("$<.statusMessage")
	return statusMessage
}
