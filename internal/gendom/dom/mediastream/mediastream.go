package mediastream

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrackevent"
	"github.com/matthewmueller/golly/js"
)

type MediaStream struct {
	*eventtarget.EventTarget
}

func (*MediaStream) AddTrack(track *mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.addTrack($1)", track)
}

func (*MediaStream) Clone() (m *MediaStream) {
	js.Rewrite("$<.clone()")
	return m
}

func (*MediaStream) GetAudioTracks() (m []*mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.getAudioTracks()")
	return m
}

func (*MediaStream) GetTrackByID(trackId string) (m *mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.getTrackById($1)", trackId)
	return m
}

func (*MediaStream) GetTracks() (m []*mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.getTracks()")
	return m
}

func (*MediaStream) GetVideoTracks() (m []*mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.getVideoTracks()")
	return m
}

func (*MediaStream) RemoveTrack(track *mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.removeTrack($1)", track)
}

func (*MediaStream) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MediaStream) GetActive() (active bool) {
	js.Rewrite("$<.active")
	return active
}

func (*MediaStream) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*MediaStream) GetOnactive() (active *event.Event) {
	js.Rewrite("$<.onactive")
	return active
}

func (*MediaStream) SetOnactive(active *event.Event) {
	js.Rewrite("$<.onactive = $1", active)
}

func (*MediaStream) GetOnaddtrack() (addtrack *mediastreamtrackevent.MediaStreamTrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*MediaStream) SetOnaddtrack(addtrack *mediastreamtrackevent.MediaStreamTrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}

func (*MediaStream) GetOninactive() (inactive *event.Event) {
	js.Rewrite("$<.oninactive")
	return inactive
}

func (*MediaStream) SetOninactive(inactive *event.Event) {
	js.Rewrite("$<.oninactive = $1", inactive)
}

func (*MediaStream) GetOnremovetrack() (removetrack *mediastreamtrackevent.MediaStreamTrackEvent) {
	js.Rewrite("$<.onremovetrack")
	return removetrack
}

func (*MediaStream) SetOnremovetrack(removetrack *mediastreamtrackevent.MediaStreamTrackEvent) {
	js.Rewrite("$<.onremovetrack = $1", removetrack)
}
