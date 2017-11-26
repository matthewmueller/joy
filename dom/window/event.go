package window

import "github.com/matthewmueller/golly/js"

// js:"Event,omit"
type Event interface {

	// InitEvent
	// js:"initEvent,rewrite=initevent"
	InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool)

	// PreventDefault
	// js:"preventDefault,rewrite=preventdefault"
	PreventDefault()

	// StopImmediatePropagation
	// js:"stopImmediatePropagation,rewrite=stopimmediatepropagation"
	StopImmediatePropagation()

	// StopPropagation
	// js:"stopPropagation,rewrite=stoppropagation"
	StopPropagation()

	// Bubbles prop
	// js:"bubbles,rewrite=bubbles"
	Bubbles() (bubbles bool)

	// Cancelable prop
	// js:"cancelable,rewrite=cancelable"
	Cancelable() (cancelable bool)

	// CancelBubble prop
	// js:"cancelBubble,rewrite=cancelbubble"
	CancelBubble() (cancelBubble bool)

	// CancelBubble prop
	// js:"setcancelBubble,rewrite=setcancelbubble"
	SetCancelBubble(cancelBubble bool)

	// CurrentTarget prop
	// js:"currentTarget,rewrite=currenttarget"
	CurrentTarget() (currentTarget EventTarget)

	// DefaultPrevented prop
	// js:"defaultPrevented,rewrite=defaultprevented"
	DefaultPrevented() (defaultPrevented bool)

	// EventPhase prop
	// js:"eventPhase,rewrite=eventphase"
	EventPhase() (eventPhase uint8)

	// IsTrusted prop
	// js:"isTrusted,rewrite=istrusted"
	IsTrusted() (isTrusted bool)

	// ReturnValue prop
	// js:"returnValue,rewrite=returnvalue"
	ReturnValue() (returnValue bool)

	// ReturnValue prop
	// js:"setreturnValue,rewrite=setreturnvalue"
	SetReturnValue(returnValue bool)

	// SrcElement prop
	// js:"srcElement,rewrite=srcelement"
	SrcElement() (srcElement Element)

	// Target prop
	// js:"target,rewrite=target"
	Target() (target EventTarget)

	// TimeStamp prop
	// js:"timeStamp,rewrite=timestamp"
	TimeStamp() (timeStamp uint64)

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind string)
}

// initevent fn
func initevent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$<.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// preventdefault fn
func preventdefault() {
	js.Rewrite("$<.preventDefault()")
}

// stopimmediatepropagation fn
func stopimmediatepropagation() {
	js.Rewrite("$<.stopImmediatePropagation()")
}

// stoppropagation fn
func stoppropagation() {
	js.Rewrite("$<.stopPropagation()")
}

// bubbles prop
func bubbles() (bubbles bool) {
	js.Rewrite("$<.bubbles")
	return bubbles
}

// cancelable prop
func cancelable() (cancelable bool) {
	js.Rewrite("$<.cancelable")
	return cancelable
}

// cancelbubble prop
func cancelbubble() (cancelBubble bool) {
	js.Rewrite("$<.cancelBubble")
	return cancelBubble
}

// setcancelbubble prop
func setcancelbubble(cancelBubble bool) {
	js.Rewrite("$<.cancelBubble = cancelBubble")
}

// currenttarget prop
func currenttarget() (currentTarget EventTarget) {
	js.Rewrite("$<.currentTarget")
	return currentTarget
}

// defaultprevented prop
func defaultprevented() (defaultPrevented bool) {
	js.Rewrite("$<.defaultPrevented")
	return defaultPrevented
}

// eventphase prop
func eventphase() (eventPhase uint8) {
	js.Rewrite("$<.eventPhase")
	return eventPhase
}

// istrusted prop
func istrusted() (isTrusted bool) {
	js.Rewrite("$<.isTrusted")
	return isTrusted
}

// returnvalue prop
func returnvalue() (returnValue bool) {
	js.Rewrite("$<.returnValue")
	return returnValue
}

// setreturnvalue prop
func setreturnvalue(returnValue bool) {
	js.Rewrite("$<.returnValue = returnValue")
}

// srcelement prop
func srcelement() (srcElement Element) {
	js.Rewrite("$<.srcElement")
	return srcElement
}

// target prop
func target() (target EventTarget) {
	js.Rewrite("$<.target")
	return target
}

// timestamp prop
func timestamp() (timeStamp uint64) {
	js.Rewrite("$<.timeStamp")
	return timeStamp
}

// // kind prop
// func kind() (kind string) {
// 	js.Rewrite("$<.type")
// 	return kind
// }
