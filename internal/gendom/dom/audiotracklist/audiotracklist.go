package audiotracklist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiotrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/trackevent"
	"github.com/matthewmueller/golly/js"
)

type AudioTrackList struct {
	*eventtarget.EventTarget
}

func (*AudioTrackList) GetTrackByID(id string) (a *audiotrack.AudioTrack) {
	js.Rewrite("$<.getTrackById($1)", id)
	return a
}

func (*AudioTrackList) Item(index uint) (a *audiotrack.AudioTrack) {
	js.Rewrite("$<.item($1)", index)
	return a
}

func (*AudioTrackList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*AudioTrackList) GetOnaddtrack() (addtrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*AudioTrackList) SetOnaddtrack(addtrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}

func (*AudioTrackList) GetOnchange() (change *event.Event) {
	js.Rewrite("$<.onchange")
	return change
}

func (*AudioTrackList) SetOnchange(change *event.Event) {
	js.Rewrite("$<.onchange = $1", change)
}

func (*AudioTrackList) GetOnremovetrack() (removetrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onremovetrack")
	return removetrack
}

func (*AudioTrackList) SetOnremovetrack(removetrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onremovetrack = $1", removetrack)
}
