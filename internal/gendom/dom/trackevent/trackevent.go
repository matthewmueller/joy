package trackevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type TrackEvent struct {
	*event.Event
}

func (*TrackEvent) GetTrack() (track interface{}) {
	js.Rewrite("$<.track")
	return track
}
