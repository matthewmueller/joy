package listeningstatechangedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/listeningstate"
	"github.com/matthewmueller/golly/js"
)

type ListeningStateChangedEvent struct {
	*event.Event
}

func (*ListeningStateChangedEvent) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*ListeningStateChangedEvent) GetState() (state *listeningstate.ListeningState) {
	js.Rewrite("$<.state")
	return state
}
