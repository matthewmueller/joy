package errorevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type ErrorEvent struct {
	*event.Event
}

func (*ErrorEvent) InitErrorEvent(typeArg string, canBubbleArg bool, cancelableArg bool, messageArg string, filenameArg string, linenoArg uint) {
	js.Rewrite("$<.initErrorEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, messageArg, filenameArg, linenoArg)
}

func (*ErrorEvent) GetColno() (colno uint) {
	js.Rewrite("$<.colno")
	return colno
}

func (*ErrorEvent) GetError() (err interface{}) {
	js.Rewrite("$<.error")
	return err
}

func (*ErrorEvent) GetFilename() (filename string) {
	js.Rewrite("$<.filename")
	return filename
}

func (*ErrorEvent) GetLineno() (lineno uint) {
	js.Rewrite("$<.lineno")
	return lineno
}

func (*ErrorEvent) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}
