package navigationcompletedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/navigationevent"
	"github.com/matthewmueller/golly/js"
)

type NavigationCompletedEvent struct {
	*navigationevent.NavigationEvent
}

func (*NavigationCompletedEvent) GetIsSuccess() (isSuccess bool) {
	js.Rewrite("$<.isSuccess")
	return isSuccess
}

func (*NavigationCompletedEvent) GetWebErrorStatus() (webErrorStatus uint64) {
	js.Rewrite("$<.webErrorStatus")
	return webErrorStatus
}
