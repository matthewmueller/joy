package event

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/js"
)

type Event struct {
}

func (*Event) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$<.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*Event) PreventDefault() {
	js.Rewrite("$<.preventDefault()")
}

func (*Event) StopImmediatePropagation() {
	js.Rewrite("$<.stopImmediatePropagation()")
}

func (*Event) StopPropagation() {
	js.Rewrite("$<.stopPropagation()")
}

func (*Event) GetBubbles() (bubbles bool) {
	js.Rewrite("$<.bubbles")
	return bubbles
}

func (*Event) GetCancelable() (cancelable bool) {
	js.Rewrite("$<.cancelable")
	return cancelable
}

func (*Event) GetCancelBubble() (cancelBubble bool) {
	js.Rewrite("$<.cancelBubble")
	return cancelBubble
}

func (*Event) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$<.cancelBubble = $1", cancelBubble)
}

func (*Event) GetCurrentTarget() (currentTarget *eventtarget.EventTarget) {
	js.Rewrite("$<.currentTarget")
	return currentTarget
}

func (*Event) GetDefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$<.defaultPrevented")
	return defaultPrevented
}

func (*Event) GetEventPhase() (eventPhase uint8) {
	js.Rewrite("$<.eventPhase")
	return eventPhase
}

func (*Event) GetIsTrusted() (isTrusted bool) {
	js.Rewrite("$<.isTrusted")
	return isTrusted
}

func (*Event) GetReturnValue() (returnValue bool) {
	js.Rewrite("$<.returnValue")
	return returnValue
}

func (*Event) SetReturnValue(returnValue bool) {
	js.Rewrite("$<.returnValue = $1", returnValue)
}

func (*Event) GetSrcElement() (srcElement *element.Element) {
	js.Rewrite("$<.srcElement")
	return srcElement
}

func (*Event) GetTarget() (target *eventtarget.EventTarget) {
	js.Rewrite("$<.target")
	return target
}

func (*Event) GetTimeStamp() (timeStamp uint64) {
	js.Rewrite("$<.timeStamp")
	return timeStamp
}

func (*Event) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
