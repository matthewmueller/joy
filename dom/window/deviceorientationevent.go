package window

import "github.com/matthewmueller/golly/js"

var _ Event = (*DeviceOrientationEvent)(nil)

// DeviceOrientationEvent struct
// js:"DeviceOrientationEvent,omit"
type DeviceOrientationEvent struct {
}

// InitDeviceOrientationEvent fn
// js:"initDeviceOrientationEvent"
func (*DeviceOrientationEvent) InitDeviceOrientationEvent(kind string, bubbles bool, cancelable bool, alpha float32, beta float32, gamma float32, absolute bool) {
	js.Rewrite("$_.initDeviceOrientationEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, alpha, beta, gamma, absolute)
}

// InitEvent fn
// js:"initEvent"
func (*DeviceOrientationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*DeviceOrientationEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*DeviceOrientationEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*DeviceOrientationEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Absolute prop
// js:"absolute"
func (*DeviceOrientationEvent) Absolute() (absolute bool) {
	js.Rewrite("$_.absolute")
	return absolute
}

// Alpha prop
// js:"alpha"
func (*DeviceOrientationEvent) Alpha() (alpha float32) {
	js.Rewrite("$_.alpha")
	return alpha
}

// Beta prop
// js:"beta"
func (*DeviceOrientationEvent) Beta() (beta float32) {
	js.Rewrite("$_.beta")
	return beta
}

// Gamma prop
// js:"gamma"
func (*DeviceOrientationEvent) Gamma() (gamma float32) {
	js.Rewrite("$_.gamma")
	return gamma
}

// Bubbles prop
// js:"bubbles"
func (*DeviceOrientationEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*DeviceOrientationEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*DeviceOrientationEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*DeviceOrientationEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*DeviceOrientationEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*DeviceOrientationEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*DeviceOrientationEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*DeviceOrientationEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*DeviceOrientationEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*DeviceOrientationEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*DeviceOrientationEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*DeviceOrientationEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*DeviceOrientationEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*DeviceOrientationEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
