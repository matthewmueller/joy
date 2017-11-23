package listeningstatechangedevent

import (
	"github.com/matthewmueller/golly/dom2/listeningstate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// ListeningStateChangedEvent struct
// js:"ListeningStateChangedEvent,omit"
type ListeningStateChangedEvent struct {
	window.Event
}

// Label
func (*ListeningStateChangedEvent) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// State
func (*ListeningStateChangedEvent) State() (state *listeningstate.ListeningState) {
	js.Rewrite("$<.State")
	return state
}
