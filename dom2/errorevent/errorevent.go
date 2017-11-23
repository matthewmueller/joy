package errorevent

import (
	"github.com/matthewmueller/golly/dom2/erroreventinit"
	"github.com/matthewmueller/golly/dom2/window"
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

// InitErrorEvent
func (*ErrorEvent) InitErrorEvent(typeArg string, canBubbleArg bool, cancelableArg bool, messageArg string, filenameArg string, linenoArg uint) {
	js.Rewrite("$<.InitErrorEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, messageArg, filenameArg, linenoArg)
}

// Colno
func (*ErrorEvent) Colno() (colno uint) {
	js.Rewrite("$<.Colno")
	return colno
}

// Error
func (*ErrorEvent) Error() (err interface{}) {
	js.Rewrite("$<.Error")
	return err
}

// Filename
func (*ErrorEvent) Filename() (filename string) {
	js.Rewrite("$<.Filename")
	return filename
}

// Lineno
func (*ErrorEvent) Lineno() (lineno uint) {
	js.Rewrite("$<.Lineno")
	return lineno
}

// Message
func (*ErrorEvent) Message() (message string) {
	js.Rewrite("$<.Message")
	return message
}
