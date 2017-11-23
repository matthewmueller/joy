package window

import "github.com/matthewmueller/golly/dom/eventinit"

type MediaStreamTrackEventInit struct {
	*eventinit.EventInit

	track *MediaStreamTrack
}
