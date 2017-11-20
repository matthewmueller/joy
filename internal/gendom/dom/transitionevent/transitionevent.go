package transitionevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type TransitionEvent struct {
	*event.Event
}

func (*TransitionEvent) InitTransitionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, propertyNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$<.initTransitionEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, propertyNameArg, elapsedTimeArg)
}

func (*TransitionEvent) GetElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}

func (*TransitionEvent) GetPropertyName() (propertyName string) {
	js.Rewrite("$<.propertyName")
	return propertyName
}
