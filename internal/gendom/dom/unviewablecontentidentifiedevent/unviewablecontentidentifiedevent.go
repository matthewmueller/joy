package unviewablecontentidentifiedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/navigationeventwithreferrer"
	"github.com/matthewmueller/golly/js"
)

type UnviewableContentIdentifiedEvent struct {
	*navigationeventwithreferrer.NavigationEventWithReferrer
}

func (*UnviewableContentIdentifiedEvent) GetMediaType() (mediaType string) {
	js.Rewrite("$<.mediaType")
	return mediaType
}
