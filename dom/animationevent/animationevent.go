package animationevent

import (
	"github.com/matthewmueller/golly/dom/animationeventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *animationeventinit.AnimationEventInit) *AnimationEvent {
	js.Rewrite("AnimationEvent")
	return &AnimationEvent{}
}

// AnimationEvent struct
// js:"AnimationEvent,omit"
type AnimationEvent struct {
	window.Event
}

// InitAnimationEvent fn
func (*AnimationEvent) InitAnimationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, animationNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$<.initAnimationEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, animationNameArg, elapsedTimeArg)
}

// AnimationName prop
func (*AnimationEvent) AnimationName() (animationName string) {
	js.Rewrite("$<.animationName")
	return animationName
}

// ElapsedTime prop
func (*AnimationEvent) ElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}
