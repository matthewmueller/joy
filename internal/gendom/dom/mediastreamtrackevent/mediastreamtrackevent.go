package mediastreamtrackevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrack"
	"github.com/matthewmueller/golly/js"
)

type MediaStreamTrackEvent struct {
	*event.Event
}

func (*MediaStreamTrackEvent) GetTrack() (track *mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}
