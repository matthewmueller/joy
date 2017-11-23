package webglcontextevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"WebGLContextEvent,omit"
type WebGLContextEvent struct {
	window.Event
}

// StatusMessage
func (*WebGLContextEvent) StatusMessage() (statusMessage string) {
	js.Rewrite("$<.StatusMessage")
	return statusMessage
}
