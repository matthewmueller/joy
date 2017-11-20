package videotracklist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/trackevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/videotrack"
	"github.com/matthewmueller/golly/js"
)

type VideoTrackList struct {
	*eventtarget.EventTarget
}

func (*VideoTrackList) GetTrackByID(id string) (v *videotrack.VideoTrack) {
	js.Rewrite("$<.getTrackById($1)", id)
	return v
}

func (*VideoTrackList) Item(index uint) (v *videotrack.VideoTrack) {
	js.Rewrite("$<.item($1)", index)
	return v
}

func (*VideoTrackList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*VideoTrackList) GetOnaddtrack() (addtrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*VideoTrackList) SetOnaddtrack(addtrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}

func (*VideoTrackList) GetOnchange() (change *event.Event) {
	js.Rewrite("$<.onchange")
	return change
}

func (*VideoTrackList) SetOnchange(change *event.Event) {
	js.Rewrite("$<.onchange = $1", change)
}

func (*VideoTrackList) GetOnremovetrack() (removetrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onremovetrack")
	return removetrack
}

func (*VideoTrackList) SetOnremovetrack(removetrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onremovetrack = $1", removetrack)
}

func (*VideoTrackList) GetSelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}
