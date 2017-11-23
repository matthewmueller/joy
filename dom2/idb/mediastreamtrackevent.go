package idb

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

// Track
func (*MediaStreamTrackEvent) Track() (track *MediaStreamTrack) {
	js.Rewrite("$<.Track")
	return track
}
