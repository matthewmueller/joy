package texttrackcue

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/documentfragment"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrack"
	"github.com/matthewmueller/golly/js"
)

type TextTrackCue struct {
	*eventtarget.EventTarget
}

func (*TextTrackCue) GetCueAsHTML() (d *documentfragment.DocumentFragment) {
	js.Rewrite("$<.getCueAsHTML()")
	return d
}

func (*TextTrackCue) GetEndTime() (endTime float32) {
	js.Rewrite("$<.endTime")
	return endTime
}

func (*TextTrackCue) SetEndTime(endTime float32) {
	js.Rewrite("$<.endTime = $1", endTime)
}

func (*TextTrackCue) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*TextTrackCue) SetID(id string) {
	js.Rewrite("$<.id = $1", id)
}

func (*TextTrackCue) GetOnenter() (enter *event.Event) {
	js.Rewrite("$<.onenter")
	return enter
}

func (*TextTrackCue) SetOnenter(enter *event.Event) {
	js.Rewrite("$<.onenter = $1", enter)
}

func (*TextTrackCue) GetOnexit() (exit *event.Event) {
	js.Rewrite("$<.onexit")
	return exit
}

func (*TextTrackCue) SetOnexit(exit *event.Event) {
	js.Rewrite("$<.onexit = $1", exit)
}

func (*TextTrackCue) GetPauseOnExit() (pauseOnExit bool) {
	js.Rewrite("$<.pauseOnExit")
	return pauseOnExit
}

func (*TextTrackCue) SetPauseOnExit(pauseOnExit bool) {
	js.Rewrite("$<.pauseOnExit = $1", pauseOnExit)
}

func (*TextTrackCue) GetStartTime() (startTime float32) {
	js.Rewrite("$<.startTime")
	return startTime
}

func (*TextTrackCue) SetStartTime(startTime float32) {
	js.Rewrite("$<.startTime = $1", startTime)
}

func (*TextTrackCue) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*TextTrackCue) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*TextTrackCue) GetTrack() (track *texttrack.TextTrack) {
	js.Rewrite("$<.track")
	return track
}
