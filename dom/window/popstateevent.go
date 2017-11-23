package window

import (
	"github.com/matthewmueller/golly/dom2/popstateeventinit"
	"github.com/matthewmueller/golly/js"
)

// NewPopStateEvent fn
func NewPopStateEvent(typearg string, eventinitdict *popstateeventinit.PopStateEventInit) *PopStateEvent {
	js.Rewrite("PopStateEvent")
	return &PopStateEvent{}
}

// PopStateEvent struct
// js:"PopStateEvent,omit"
type PopStateEvent struct {
	Event
}

// InitPopStateEvent fn
func (*PopStateEvent) InitPopStateEvent(typeArg string, canBubbleArg bool, cancelableArg bool, stateArg interface{}) {
	js.Rewrite("$<.initPopStateEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, stateArg)
}

// State prop
func (*PopStateEvent) State() (state interface{}) {
	js.Rewrite("$<.state")
	return state
}
