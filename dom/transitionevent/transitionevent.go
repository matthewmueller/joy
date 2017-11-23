package transitionevent

import (
	"github.com/matthewmueller/golly/dom2/transitioneventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *transitioneventinit.TransitionEventInit) *TransitionEvent {
	js.Rewrite("TransitionEvent")
	return &TransitionEvent{}
}

// TransitionEvent struct
// js:"TransitionEvent,omit"
type TransitionEvent struct {
	window.Event
}

// InitTransitionEvent fn
func (*TransitionEvent) InitTransitionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, propertyNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$<.initTransitionEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, propertyNameArg, elapsedTimeArg)
}

// ElapsedTime prop
func (*TransitionEvent) ElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}

// PropertyName prop
func (*TransitionEvent) PropertyName() (propertyName string) {
	js.Rewrite("$<.propertyName")
	return propertyName
}
