package texttrack

import "github.com/matthewmueller/golly/dom/window"

// TextTrackCue interface
// js:"TextTrackCue"
type TextTrackCue interface {
	window.EventTarget

	// GetCueAsHTML
	// js:"getCueAsHTML"
	// jsrewrite:"$_.getCueAsHTML()"
	GetCueAsHTML() (w *window.DocumentFragment)

	// EndTime prop
	// js:"endTime"
	// jsrewrite:"$_.endTime"
	EndTime() (endTime float32)

	// SetEndTime prop
	// js:"endTime"
	// jsrewrite:"$_.endTime = $1"
	SetEndTime(endTime float32)

	// ID prop
	// js:"id"
	// jsrewrite:"$_.id"
	ID() (id string)

	// SetID prop
	// js:"id"
	// jsrewrite:"$_.id = $1"
	SetID(id string)

	// Onenter prop
	// js:"onenter"
	// jsrewrite:"$_.onenter"
	Onenter() (onenter func(window.Event))

	// SetOnenter prop
	// js:"onenter"
	// jsrewrite:"$_.onenter = $1"
	SetOnenter(onenter func(window.Event))

	// Onexit prop
	// js:"onexit"
	// jsrewrite:"$_.onexit"
	Onexit() (onexit func(window.Event))

	// SetOnexit prop
	// js:"onexit"
	// jsrewrite:"$_.onexit = $1"
	SetOnexit(onexit func(window.Event))

	// PauseOnExit prop
	// js:"pauseOnExit"
	// jsrewrite:"$_.pauseOnExit"
	PauseOnExit() (pauseOnExit bool)

	// SetPauseOnExit prop
	// js:"pauseOnExit"
	// jsrewrite:"$_.pauseOnExit = $1"
	SetPauseOnExit(pauseOnExit bool)

	// StartTime prop
	// js:"startTime"
	// jsrewrite:"$_.startTime"
	StartTime() (startTime float32)

	// SetStartTime prop
	// js:"startTime"
	// jsrewrite:"$_.startTime = $1"
	SetStartTime(startTime float32)

	// Text prop
	// js:"text"
	// jsrewrite:"$_.text"
	Text() (text string)

	// SetText prop
	// js:"text"
	// jsrewrite:"$_.text = $1"
	SetText(text string)

	// Track prop
	// js:"track"
	// jsrewrite:"$_.track"
	Track() (track *TextTrack)
}
