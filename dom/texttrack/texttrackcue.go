package texttrack

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"TextTrackCue,omit"
type TextTrackCue interface {
	window.EventTarget

	// GetCueAsHTML
	// js:"getCueAsHTML,rewrite=getcueashtml"
	GetCueAsHTML() (w *window.DocumentFragment)

	// EndTime prop
	// js:"endTime,rewrite=endtime"
	EndTime() (endTime float32)

	// EndTime prop
	// js:"setendTime,rewrite=setendtime"
	SetEndTime(endTime float32)

	// ID prop
	// js:"id,rewrite=id"
	ID() (id string)

	// ID prop
	// js:"setid,rewrite=setid"
	SetID(id string)

	// Onenter prop
	// js:"onenter,rewrite=onenter"
	Onenter() (onenter func(window.Event))

	// Onenter prop
	// js:"setonenter,rewrite=setonenter"
	SetOnenter(onenter func(window.Event))

	// Onexit prop
	// js:"onexit,rewrite=onexit"
	Onexit() (onexit func(window.Event))

	// Onexit prop
	// js:"setonexit,rewrite=setonexit"
	SetOnexit(onexit func(window.Event))

	// PauseOnExit prop
	// js:"pauseOnExit,rewrite=pauseonexit"
	PauseOnExit() (pauseOnExit bool)

	// PauseOnExit prop
	// js:"setpauseOnExit,rewrite=setpauseonexit"
	SetPauseOnExit(pauseOnExit bool)

	// StartTime prop
	// js:"startTime,rewrite=starttime"
	StartTime() (startTime float32)

	// StartTime prop
	// js:"setstartTime,rewrite=setstarttime"
	SetStartTime(startTime float32)

	// Text prop
	// js:"text,rewrite=text"
	Text() (text string)

	// Text prop
	// js:"settext,rewrite=settext"
	SetText(text string)

	// Track prop
	// js:"track,rewrite=track"
	Track() (track *TextTrack)
}

// getcueashtml fn
func getcueashtml() (w *window.DocumentFragment) {
	js.Rewrite("$<.getCueAsHTML()")
	return w
}

// endtime prop
func endtime() (endTime float32) {
	js.Rewrite("$<.endTime")
	return endTime
}

// setendtime prop
func setendtime(endTime float32) {
	js.Rewrite("$<.endTime = endTime")
}

// id prop
func id() (id string) {
	js.Rewrite("$<.id")
	return id
}

// setid prop
func setid(id string) {
	js.Rewrite("$<.id = id")
}

// onenter prop
func onenter() (onenter func(window.Event)) {
	js.Rewrite("$<.onenter")
	return onenter
}

// setonenter prop
func setonenter(onenter func(window.Event)) {
	js.Rewrite("$<.onenter = onenter")
}

// onexit prop
func onexit() (onexit func(window.Event)) {
	js.Rewrite("$<.onexit")
	return onexit
}

// setonexit prop
func setonexit(onexit func(window.Event)) {
	js.Rewrite("$<.onexit = onexit")
}

// pauseonexit prop
func pauseonexit() (pauseOnExit bool) {
	js.Rewrite("$<.pauseOnExit")
	return pauseOnExit
}

// setpauseonexit prop
func setpauseonexit(pauseOnExit bool) {
	js.Rewrite("$<.pauseOnExit = pauseOnExit")
}

// starttime prop
func starttime() (startTime float32) {
	js.Rewrite("$<.startTime")
	return startTime
}

// setstarttime prop
func setstarttime(startTime float32) {
	js.Rewrite("$<.startTime = startTime")
}

// text prop
func text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// settext prop
func settext(text string) {
	js.Rewrite("$<.text = text")
}

// track prop
func track() (track *TextTrack) {
	js.Rewrite("$<.track")
	return track
}
