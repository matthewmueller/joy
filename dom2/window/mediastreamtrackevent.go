package window

import "github.com/matthewmueller/golly/js"

// js:"MediaStreamTrackEvent,omit"
type MediaStreamTrackEvent struct {
	Event
}

// Track
func (*MediaStreamTrackEvent) Track() (track *MediaStreamTrack) {
	js.Rewrite("$<.Track")
	return track
}
