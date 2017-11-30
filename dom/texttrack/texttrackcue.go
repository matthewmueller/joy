package texttrack

import "github.com/matthewmueller/golly/dom/window"

// TextTrackCue interface
// js:"TextTrackCue"
type TextTrackCue interface {
	window.EventTarget

	// GetCueAsHTML
	// js:"getCueAsHTML"
	GetCueAsHTML() (w *window.DocumentFragment)

	// EndTime prop
	// js:"endTime"
	EndTime() (endTime float32)

	// SetEndTime prop
	// js:"endTime"
	SetEndTime(endTime float32)

	// ID prop
	// js:"id"
	ID() (id string)

	// SetID prop
	// js:"id"
	SetID(id string)

	// Onenter prop
	// js:"onenter"
	Onenter() (onenter func(window.Event))

	// SetOnenter prop
	// js:"onenter"
	SetOnenter(onenter func(window.Event))

	// Onexit prop
	// js:"onexit"
	Onexit() (onexit func(window.Event))

	// SetOnexit prop
	// js:"onexit"
	SetOnexit(onexit func(window.Event))

	// PauseOnExit prop
	// js:"pauseOnExit"
	PauseOnExit() (pauseOnExit bool)

	// SetPauseOnExit prop
	// js:"pauseOnExit"
	SetPauseOnExit(pauseOnExit bool)

	// StartTime prop
	// js:"startTime"
	StartTime() (startTime float32)

	// SetStartTime prop
	// js:"startTime"
	SetStartTime(startTime float32)

	// Text prop
	// js:"text"
	Text() (text string)

	// SetText prop
	// js:"text"
	SetText(text string)

	// Track prop
	// js:"track"
	Track() (track *TextTrack)
}
