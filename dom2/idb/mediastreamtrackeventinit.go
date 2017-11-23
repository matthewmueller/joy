package idb

import "github.com/matthewmueller/golly/dom2/eventinit"

type MediaStreamTrackEventInit struct {
	*eventinit.EventInit

	track *MediaStreamTrack
}
