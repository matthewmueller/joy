package navigationeventwithreferrer

import "github.com/matthewmueller/golly/dom/navigationevent"

// js:"NavigationEventWithReferrer,omit"
type NavigationEventWithReferrer interface {
	navigationevent.NavigationEvent

	// Referer prop
	// js:"referer"
	Referer() (referer string)
}
