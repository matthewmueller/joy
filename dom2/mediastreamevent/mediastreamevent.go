package mediastreamevent

import (
	"github.com/matthewmueller/golly/dom2/mediastreameventinit"
	"github.com/matthewmueller/golly/dom2/window"
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

// Stream
func (*MediaStreamEvent) Stream() (stream *window.MediaStream) {
	js.Rewrite("$<.Stream")
	return stream
}
