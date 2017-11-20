package webglcontextevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type WebGLContextEvent struct {
	*event.Event
}

func (*WebGLContextEvent) GetStatusMessage() (statusMessage string) {
	js.Rewrite("$<.statusMessage")
	return statusMessage
}
