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

// Label prop
func (*ListeningStateChangedEvent) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// State prop
func (*ListeningStateChangedEvent) State() (state *listeningstate.ListeningState) {
	js.Rewrite("$<.state")
	return state
}
