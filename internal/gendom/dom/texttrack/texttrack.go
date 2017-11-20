package texttrack

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrackcue"
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrackcuelist"
	"github.com/matthewmueller/golly/js"
)

type TextTrack struct {
	*eventtarget.EventTarget
}

func (*TextTrack) AddCue(cue *texttrackcue.TextTrackCue) {
	js.Rewrite("$<.addCue($1)", cue)
}

func (*TextTrack) RemoveCue(cue *texttrackcue.TextTrackCue) {
	js.Rewrite("$<.removeCue($1)", cue)
}

func (*TextTrack) GetActiveCues() (activeCues *texttrackcuelist.TextTrackCueList) {
	js.Rewrite("$<.activeCues")
	return activeCues
}

func (*TextTrack) GetCues() (cues *texttrackcuelist.TextTrackCueList) {
	js.Rewrite("$<.cues")
	return cues
}

func (*TextTrack) GetInBandMetadataTrackDispatchType() (inBandMetadataTrackDispatchType string) {
	js.Rewrite("$<.inBandMetadataTrackDispatchType")
	return inBandMetadataTrackDispatchType
}

func (*TextTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*TextTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*TextTrack) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*TextTrack) GetMode() (mode interface{}) {
	js.Rewrite("$<.mode")
	return mode
}

func (*TextTrack) SetMode(mode interface{}) {
	js.Rewrite("$<.mode = $1", mode)
}

func (*TextTrack) GetOncuechange() (cuechange *event.Event) {
	js.Rewrite("$<.oncuechange")
	return cuechange
}

func (*TextTrack) SetOncuechange(cuechange *event.Event) {
	js.Rewrite("$<.oncuechange = $1", cuechange)
}

func (*TextTrack) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*TextTrack) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*TextTrack) GetOnload() (load *event.Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*TextTrack) SetOnload(load *event.Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*TextTrack) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}
