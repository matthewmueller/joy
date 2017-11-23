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

// IsSuccess
func (*NavigationCompletedEvent) IsSuccess() (isSuccess bool) {
	js.Rewrite("$<.IsSuccess")
	return isSuccess
}

// WebErrorStatus
func (*NavigationCompletedEvent) WebErrorStatus() (webErrorStatus uint64) {
	js.Rewrite("$<.WebErrorStatus")
	return webErrorStatus
}
