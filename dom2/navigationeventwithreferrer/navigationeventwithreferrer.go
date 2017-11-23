package navigationeventwithreferrer

// js:"NavigationEventWithReferrer,omit"
type NavigationEventWithReferrer interface {
	navigationevent.NavigationEvent

	// Referer
	Referer() (referer string)
}
