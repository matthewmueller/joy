package trackevent

import (
	"github.com/matthewmueller/golly/dom/trackeventinit"
	"github.com/matthewmueller/golly/dom/window"
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

// Track prop
func (*TrackEvent) Track() (track interface{}) {
	js.Rewrite("$<.track")
	return track
}
