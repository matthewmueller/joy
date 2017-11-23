package animationevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"AnimationEvent,omit"
type AnimationEvent struct {
	window.Event
}

// InitAnimationEvent
func (*AnimationEvent) InitAnimationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, animationNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$<.InitAnimationEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, animationNameArg, elapsedTimeArg)
}

// AnimationName
func (*AnimationEvent) AnimationName() (animationName string) {
	js.Rewrite("$<.AnimationName")
	return animationName
}

// ElapsedTime
func (*AnimationEvent) ElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.ElapsedTime")
	return elapsedTime
}
