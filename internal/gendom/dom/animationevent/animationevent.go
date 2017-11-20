package animationevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type AnimationEvent struct {
	*event.Event
}

func (*AnimationEvent) InitAnimationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, animationNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$<.initAnimationEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, animationNameArg, elapsedTimeArg)
}

func (*AnimationEvent) GetAnimationName() (animationName string) {
	js.Rewrite("$<.animationName")
	return animationName
}

func (*AnimationEvent) GetElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}
