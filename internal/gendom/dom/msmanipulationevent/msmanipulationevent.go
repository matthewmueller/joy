package msmanipulationevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type MSManipulationEvent struct {
	*uievent.UIEvent
}

func (*MSManipulationEvent) InitMSManipulationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, lastState int, currentState int) {
	js.Rewrite("$<.initMSManipulationEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, lastState, currentState)
}

func (*MSManipulationEvent) GetCurrentState() (currentState int) {
	js.Rewrite("$<.currentState")
	return currentState
}

func (*MSManipulationEvent) GetInertiaDestinationX() (inertiaDestinationX int) {
	js.Rewrite("$<.inertiaDestinationX")
	return inertiaDestinationX
}

func (*MSManipulationEvent) GetInertiaDestinationY() (inertiaDestinationY int) {
	js.Rewrite("$<.inertiaDestinationY")
	return inertiaDestinationY
}

func (*MSManipulationEvent) GetLastState() (lastState int) {
	js.Rewrite("$<.lastState")
	return lastState
}
