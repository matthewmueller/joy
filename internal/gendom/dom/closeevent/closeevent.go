package closeevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type CloseEvent struct {
	*event.Event
}

func (*CloseEvent) InitCloseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, wasCleanArg bool, codeArg uint8, reasonArg string) {
	js.Rewrite("$<.initCloseEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, wasCleanArg, codeArg, reasonArg)
}

func (*CloseEvent) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*CloseEvent) GetReason() (reason string) {
	js.Rewrite("$<.reason")
	return reason
}

func (*CloseEvent) GetWasClean() (wasClean bool) {
	js.Rewrite("$<.wasClean")
	return wasClean
}
