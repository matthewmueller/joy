package trackevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"TrackEvent,omit"
type TrackEvent struct {
	window.Event
}

// Track
func (*TrackEvent) Track() (track interface{}) {
	js.Rewrite("$<.Track")
	return track
}
