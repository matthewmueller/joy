package unviewablecontentidentifiedevent

import (
	"github.com/matthewmueller/golly/dom/navigationeventwithreferrer"
	"github.com/matthewmueller/golly/js"
)

// UnviewableContentIdentifiedEvent struct
// js:"UnviewableContentIdentifiedEvent,omit"
type UnviewableContentIdentifiedEvent struct {
	navigationeventwithreferrer.NavigationEventWithReferrer
}

// MediaType prop
func (*UnviewableContentIdentifiedEvent) MediaType() (mediaType string) {
	js.Rewrite("$<.mediaType")
	return mediaType
}
