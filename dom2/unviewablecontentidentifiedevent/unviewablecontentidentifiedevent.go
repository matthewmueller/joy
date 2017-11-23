package unviewablecontentidentifiedevent

import (
	"github.com/matthewmueller/golly/dom2/navigationeventwithreferrer"
	"github.com/matthewmueller/golly/js"
)

// UnviewableContentIdentifiedEvent struct
// js:"UnviewableContentIdentifiedEvent,omit"
type UnviewableContentIdentifiedEvent struct {
	navigationeventwithreferrer.NavigationEventWithReferrer
}

// MediaType
func (*UnviewableContentIdentifiedEvent) MediaType() (mediaType string) {
	js.Rewrite("$<.MediaType")
	return mediaType
}
