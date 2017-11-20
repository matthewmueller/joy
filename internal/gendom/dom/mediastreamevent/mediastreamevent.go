package mediastreamevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastream"
	"github.com/matthewmueller/golly/js"
)

type MediaStreamEvent struct {
	*event.Event
}

func (*MediaStreamEvent) GetStream() (stream *mediastream.MediaStream) {
	js.Rewrite("$<.stream")
	return stream
}
