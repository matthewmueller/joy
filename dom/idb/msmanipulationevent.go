package idb

import "github.com/matthewmueller/golly/js"

// MSManipulationEvent struct
// js:"MSManipulationEvent,omit"
type MSManipulationEvent struct {
	UIEvent
}

// InitMSManipulationEvent
func (*MSManipulationEvent) InitMSManipulationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, lastState int, currentState int) {
	js.Rewrite("$<.InitMSManipulationEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, lastState, currentState)
}

// CurrentState
func (*MSManipulationEvent) CurrentState() (currentState int) {
	js.Rewrite("$<.CurrentState")
	return currentState
}

// InertiaDestinationX
func (*MSManipulationEvent) InertiaDestinationX() (inertiaDestinationX int) {
	js.Rewrite("$<.InertiaDestinationX")
	return inertiaDestinationX
}

// InertiaDestinationY
func (*MSManipulationEvent) InertiaDestinationY() (inertiaDestinationY int) {
	js.Rewrite("$<.InertiaDestinationY")
	return inertiaDestinationY
}

// LastState
func (*MSManipulationEvent) LastState() (lastState int) {
	js.Rewrite("$<.LastState")
	return lastState
}
