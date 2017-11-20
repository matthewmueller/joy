package texttracklist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/trackevent"
	"github.com/matthewmueller/golly/js"
)

type TextTrackList struct {
	*eventtarget.EventTarget
}

func (*TextTrackList) Item(index uint) (t *texttrack.TextTrack) {
	js.Rewrite("$<.item($1)", index)
	return t
}

func (*TextTrackList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*TextTrackList) GetOnaddtrack() (addtrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*TextTrackList) SetOnaddtrack(addtrack *trackevent.TrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}
