package navigationeventwithreferrer

import (
	"github.com/matthewmueller/golly/dom/navigationevent"
	"github.com/matthewmueller/golly/js"
)

// js:"NavigationEventWithReferrer,omit"
type NavigationEventWithReferrer interface {
	navigationevent.NavigationEvent

	// Referer prop
	// js:"referer,rewrite=referer"
	Referer() (referer string)
}

// referer prop
func referer() (referer string) {
	js.Rewrite("$<.referer")
	return referer
}
