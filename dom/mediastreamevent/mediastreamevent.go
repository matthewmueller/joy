package mediastreamevent

import (
	"github.com/matthewmueller/golly/dom/mediastreameventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(kind string, eventinitdict *mediastreameventinit.MediaStreamEventInit) *MediaStreamEvent {
	js.Rewrite("MediaStreamEvent")
	return &MediaStreamEvent{}
}

// MediaStreamEvent struct
// js:"MediaStreamEvent,omit"
type MediaStreamEvent struct {
	window.Event
}

// Stream prop
func (*MediaStreamEvent) Stream() (stream *window.MediaStream) {
	js.Rewrite("$<.stream")
	return stream
}
