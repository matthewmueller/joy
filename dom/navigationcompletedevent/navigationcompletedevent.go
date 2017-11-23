package navigationcompletedevent

import (
	"github.com/matthewmueller/golly/dom2/navigationevent"
	"github.com/matthewmueller/golly/js"
)

// NavigationCompletedEvent struct
// js:"NavigationCompletedEvent,omit"
type NavigationCompletedEvent struct {
	navigationevent.NavigationEvent
}

// IsSuccess prop
func (*NavigationCompletedEvent) IsSuccess() (isSuccess bool) {
	js.Rewrite("$<.isSuccess")
	return isSuccess
}

// WebErrorStatus prop
func (*NavigationCompletedEvent) WebErrorStatus() (webErrorStatus uint64) {
	js.Rewrite("$<.webErrorStatus")
	return webErrorStatus
}
