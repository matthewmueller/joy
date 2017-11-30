package navigationeventwithreferrer

import "github.com/matthewmueller/golly/dom/navigationevent"

// NavigationEventWithReferrer interface
// js:"NavigationEventWithReferrer"
type NavigationEventWithReferrer interface {
	navigationevent.NavigationEvent

	// Referer prop
	// js:"referer"
	Referer() (referer string)
}
