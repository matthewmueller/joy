package window

import "github.com/matthewmueller/golly/js"

// MSManipulationEvent struct
// js:"MSManipulationEvent,omit"
type MSManipulationEvent struct {
	UIEvent
}

// InitMSManipulationEvent fn
func (*MSManipulationEvent) InitMSManipulationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, lastState int, currentState int) {
	js.Rewrite("$<.initMSManipulationEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, lastState, currentState)
}

// CurrentState prop
func (*MSManipulationEvent) CurrentState() (currentState int) {
	js.Rewrite("$<.currentState")
	return currentState
}

// InertiaDestinationX prop
func (*MSManipulationEvent) InertiaDestinationX() (inertiaDestinationX int) {
	js.Rewrite("$<.inertiaDestinationX")
	return inertiaDestinationX
}

// InertiaDestinationY prop
func (*MSManipulationEvent) InertiaDestinationY() (inertiaDestinationY int) {
	js.Rewrite("$<.inertiaDestinationY")
	return inertiaDestinationY
}

// LastState prop
func (*MSManipulationEvent) LastState() (lastState int) {
	js.Rewrite("$<.lastState")
	return lastState
}
