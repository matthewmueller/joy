package trackevent

import (
	"github.com/matthewmueller/golly/dom2/trackeventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *trackeventinit.TrackEventInit) *TrackEvent {
	js.Rewrite("TrackEvent")
	return &TrackEvent{}
}

// TrackEvent struct
// js:"TrackEvent,omit"
type TrackEvent struct {
	window.Event
}

// Track
func (*TrackEvent) Track() (track interface{}) {
	js.Rewrite("$<.Track")
	return track
}
