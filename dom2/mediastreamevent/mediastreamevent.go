package mediastreamevent

import "github.com/matthewmueller/golly/js"

// js:"MediaStreamEvent,omit"
type MediaStreamEvent struct {
	window.Event
}

// Stream
func (*MediaStreamEvent) Stream() (stream *window.MediaStream) {
	js.Rewrite("$<.Stream")
	return stream
}
