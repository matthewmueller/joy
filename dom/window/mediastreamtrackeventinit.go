package window

import "github.com/matthewmueller/joy/dom/eventinit"

type MediaStreamTrackEventInit struct {
	*eventinit.EventInit

	track *MediaStreamTrack
}
