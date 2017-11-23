package window

import "github.com/matthewmueller/golly/js"

// NewFocusEvent fn
func NewFocusEvent(typearg string, eventinitdict *FocusEventInit) *FocusEvent {
	js.Rewrite("FocusEvent")
	return &FocusEvent{}
}

// FocusEvent struct
// js:"FocusEvent,omit"
type FocusEvent struct {
	UIEvent
}

// InitFocusEvent fn
func (*FocusEvent) InitFocusEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, relatedTargetArg EventTarget) {
	js.Rewrite("$<.initFocusEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, relatedTargetArg)
}

// RelatedTarget prop
func (*FocusEvent) RelatedTarget() (relatedTarget EventTarget) {
	js.Rewrite("$<.relatedTarget")
	return relatedTarget
}
