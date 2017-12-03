package datacue

import (
	"github.com/matthewmueller/joy/dom/texttrack"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ texttrack.TextTrackCue = (*DataCue)(nil)
var _ window.EventTarget = (*DataCue)(nil)

// DataCue struct
// js:"DataCue,omit"
type DataCue struct {
}

// GetCueAsHTML fn
// js:"getCueAsHTML"
func (*DataCue) GetCueAsHTML() (w *window.DocumentFragment) {
	macro.Rewrite("$_.getCueAsHTML()")
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*DataCue) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*DataCue) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*DataCue) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Data prop
// js:"data"
func (*DataCue) Data() (data []byte) {
	macro.Rewrite("$_.data")
	return data
}

// SetData prop
// js:"data"
func (*DataCue) SetData(data []byte) {
	macro.Rewrite("$_.data = $1", data)
}

// EndTime prop
// js:"endTime"
func (*DataCue) EndTime() (endTime float32) {
	macro.Rewrite("$_.endTime")
	return endTime
}

// SetEndTime prop
// js:"endTime"
func (*DataCue) SetEndTime(endTime float32) {
	macro.Rewrite("$_.endTime = $1", endTime)
}

// ID prop
// js:"id"
func (*DataCue) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*DataCue) SetID(id string) {
	macro.Rewrite("$_.id = $1", id)
}

// Onenter prop
// js:"onenter"
func (*DataCue) Onenter() (onenter func(window.Event)) {
	macro.Rewrite("$_.onenter")
	return onenter
}

// SetOnenter prop
// js:"onenter"
func (*DataCue) SetOnenter(onenter func(window.Event)) {
	macro.Rewrite("$_.onenter = $1", onenter)
}

// Onexit prop
// js:"onexit"
func (*DataCue) Onexit() (onexit func(window.Event)) {
	macro.Rewrite("$_.onexit")
	return onexit
}

// SetOnexit prop
// js:"onexit"
func (*DataCue) SetOnexit(onexit func(window.Event)) {
	macro.Rewrite("$_.onexit = $1", onexit)
}

// PauseOnExit prop
// js:"pauseOnExit"
func (*DataCue) PauseOnExit() (pauseOnExit bool) {
	macro.Rewrite("$_.pauseOnExit")
	return pauseOnExit
}

// SetPauseOnExit prop
// js:"pauseOnExit"
func (*DataCue) SetPauseOnExit(pauseOnExit bool) {
	macro.Rewrite("$_.pauseOnExit = $1", pauseOnExit)
}

// StartTime prop
// js:"startTime"
func (*DataCue) StartTime() (startTime float32) {
	macro.Rewrite("$_.startTime")
	return startTime
}

// SetStartTime prop
// js:"startTime"
func (*DataCue) SetStartTime(startTime float32) {
	macro.Rewrite("$_.startTime = $1", startTime)
}

// Text prop
// js:"text"
func (*DataCue) Text() (text string) {
	macro.Rewrite("$_.text")
	return text
}

// SetText prop
// js:"text"
func (*DataCue) SetText(text string) {
	macro.Rewrite("$_.text = $1", text)
}

// Track prop
// js:"track"
func (*DataCue) Track() (track *texttrack.TextTrack) {
	macro.Rewrite("$_.track")
	return track
}
