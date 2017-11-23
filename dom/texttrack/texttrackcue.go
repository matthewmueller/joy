package texttrack

import "github.com/matthewmueller/golly/dom2/window"

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
	SetEndTime(endTime float32)

	// ID prop
	// js:"id"
	ID() (id string)

	// ID prop
	SetID(id string)

	// Onenter prop
	// js:"onenter"
	Onenter() (onenter func(window.Event))

	// Onenter prop
	SetOnenter(onenter func(window.Event))

	// Onexit prop
	// js:"onexit"
	Onexit() (onexit func(window.Event))

	// Onexit prop
	SetOnexit(onexit func(window.Event))

	// PauseOnExit prop
	// js:"pauseOnExit"
	PauseOnExit() (pauseOnExit bool)

	// PauseOnExit prop
	SetPauseOnExit(pauseOnExit bool)

	// StartTime prop
	// js:"startTime"
	StartTime() (startTime float32)

	// StartTime prop
	SetStartTime(startTime float32)

	// Text prop
	// js:"text"
	Text() (text string)

	// Text prop
	SetText(text string)

	// Track prop
	// js:"track"
	Track() (track *TextTrack)
}
