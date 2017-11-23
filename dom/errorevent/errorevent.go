package errorevent

import (
	"github.com/matthewmueller/golly/dom/erroreventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *erroreventinit.ErrorEventInit) *ErrorEvent {
	js.Rewrite("ErrorEvent")
	return &ErrorEvent{}
}

// ErrorEvent struct
// js:"ErrorEvent,omit"
type ErrorEvent struct {
	window.Event
}

// InitErrorEvent fn
func (*ErrorEvent) InitErrorEvent(typeArg string, canBubbleArg bool, cancelableArg bool, messageArg string, filenameArg string, linenoArg uint) {
	js.Rewrite("$<.initErrorEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, messageArg, filenameArg, linenoArg)
}

// Colno prop
func (*ErrorEvent) Colno() (colno uint) {
	js.Rewrite("$<.colno")
	return colno
}

// Error prop
func (*ErrorEvent) Error() (err interface{}) {
	js.Rewrite("$<.error")
	return err
}

// Filename prop
func (*ErrorEvent) Filename() (filename string) {
	js.Rewrite("$<.filename")
	return filename
}

// Lineno prop
func (*ErrorEvent) Lineno() (lineno uint) {
	js.Rewrite("$<.lineno")
	return lineno
}

// Message prop
func (*ErrorEvent) Message() (message string) {
	js.Rewrite("$<.message")
	return message
}
