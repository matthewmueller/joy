package texttrack

import "github.com/matthewmueller/golly/dom2/window"

// js:"TextTrackCue,omit"
type TextTrackCue interface {
	window.EventTarget

	// GetCueAsHTML
	GetCueAsHTML() (w *window.DocumentFragment)

	// EndTime
	EndTime() (endTime float32)

	// EndTime
	SetEndTime(endTime float32)

	// ID
	ID() (id string)

	// ID
	SetID(id string)

	// Onenter
	Onenter() (onenter func(window.Event))

	// Onenter
	SetOnenter(onenter func(window.Event))

	// Onexit
	Onexit() (onexit func(window.Event))

	// Onexit
	SetOnexit(onexit func(window.Event))

	// PauseOnExit
	PauseOnExit() (pauseOnExit bool)

	// PauseOnExit
	SetPauseOnExit(pauseOnExit bool)

	// StartTime
	StartTime() (startTime float32)

	// StartTime
	SetStartTime(startTime float32)

	// Text
	Text() (text string)

	// Text
	SetText(text string)

	// Track
	Track() (track *TextTrack)
}
