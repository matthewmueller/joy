package texttrack

import "github.com/matthewmueller/golly/dom/window"

// js:"TextTrackCue,omit"
type TextTrackCue interface {
	window.EventTarget

	// GetCueAsHTML
	// js:"getCueAsHTML"
	GetCueAsHTML() (w *window.DocumentFragment)

	// EndTime prop
	// js:"endTime"
	EndTime() (endTime float32)

	// EndTime prop
	// js:"setendTime"
	SetEndTime(endTime float32)

	// ID prop
	// js:"id"
	ID() (id string)

	// ID prop
	// js:"setid"
	SetID(id string)

	// Onenter prop
	// js:"onenter"
	Onenter() (onenter func(window.Event))

	// Onenter prop
	// js:"setonenter"
	SetOnenter(onenter func(window.Event))

	// Onexit prop
	// js:"onexit"
	Onexit() (onexit func(window.Event))

	// Onexit prop
	// js:"setonexit"
	SetOnexit(onexit func(window.Event))

	// PauseOnExit prop
	// js:"pauseOnExit"
	PauseOnExit() (pauseOnExit bool)

	// PauseOnExit prop
	// js:"setpauseOnExit"
	SetPauseOnExit(pauseOnExit bool)

	// StartTime prop
	// js:"startTime"
	StartTime() (startTime float32)

	// StartTime prop
	// js:"setstartTime"
	SetStartTime(startTime float32)

	// Text prop
	// js:"text"
	Text() (text string)

	// Text prop
	// js:"settext"
	SetText(text string)

	// Track prop
	// js:"track"
	Track() (track *TextTrack)
}
