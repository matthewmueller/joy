package window

import "github.com/matthewmueller/golly/js"

// NewMediaStreamTrackEvent fn
func NewMediaStreamTrackEvent(typearg string, eventinitdict *MediaStreamTrackEventInit) *MediaStreamTrackEvent {
	js.Rewrite("MediaStreamTrackEvent")
	return &MediaStreamTrackEvent{}
}

// MediaStreamTrackEvent struct
// js:"MediaStreamTrackEvent,omit"
type MediaStreamTrackEvent struct {
	Event
}

// Track prop
func (*MediaStreamTrackEvent) Track() (track *MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}
