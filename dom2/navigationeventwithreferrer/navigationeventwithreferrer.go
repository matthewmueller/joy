package navigationeventwithreferrer

import "github.com/matthewmueller/golly/dom2/navigationevent"

// js:"NavigationEventWithReferrer,omit"
type NavigationEventWithReferrer interface {
	navigationevent.NavigationEvent

	// Referer
	Referer() (referer string)
}
