package window

import (
	"github.com/matthewmueller/golly/dom/deviceacceleration"
	"github.com/matthewmueller/golly/dom/deviceaccelerationdict"
	"github.com/matthewmueller/golly/dom/devicerotationrate"
	"github.com/matthewmueller/golly/dom/devicerotationratedict"
	"github.com/matthewmueller/golly/js"
)

var _ Event = (*DeviceMotionEvent)(nil)

// DeviceMotionEvent struct
// js:"DeviceMotionEvent,omit"
type DeviceMotionEvent struct {
}

// InitDeviceMotionEvent fn
// js:"initDeviceMotionEvent"
func (*DeviceMotionEvent) InitDeviceMotionEvent(kind string, bubbles bool, cancelable bool, acceleration *deviceaccelerationdict.DeviceAccelerationDict, accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict, rotationRate *devicerotationratedict.DeviceRotationRateDict, interval float32) {
	js.Rewrite("$_.initDeviceMotionEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, acceleration, accelerationIncludingGravity, rotationRate, interval)
}

// InitEvent fn
// js:"initEvent"
func (*DeviceMotionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*DeviceMotionEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*DeviceMotionEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*DeviceMotionEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Acceleration prop
// js:"acceleration"
func (*DeviceMotionEvent) Acceleration() (acceleration *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$_.acceleration")
	return acceleration
}

// AccelerationIncludingGravity prop
// js:"accelerationIncludingGravity"
func (*DeviceMotionEvent) AccelerationIncludingGravity() (accelerationIncludingGravity *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$_.accelerationIncludingGravity")
	return accelerationIncludingGravity
}

// Interval prop
// js:"interval"
func (*DeviceMotionEvent) Interval() (interval float32) {
	js.Rewrite("$_.interval")
	return interval
}

// RotationRate prop
// js:"rotationRate"
func (*DeviceMotionEvent) RotationRate() (rotationRate *devicerotationrate.DeviceRotationRate) {
	js.Rewrite("$_.rotationRate")
	return rotationRate
}

// Bubbles prop
// js:"bubbles"
func (*DeviceMotionEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*DeviceMotionEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*DeviceMotionEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*DeviceMotionEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*DeviceMotionEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*DeviceMotionEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*DeviceMotionEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*DeviceMotionEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*DeviceMotionEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*DeviceMotionEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*DeviceMotionEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*DeviceMotionEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*DeviceMotionEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*DeviceMotionEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
