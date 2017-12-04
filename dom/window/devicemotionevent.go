package window

import (
	"github.com/matthewmueller/joy/dom/deviceacceleration"
	"github.com/matthewmueller/joy/dom/deviceaccelerationdict"
	"github.com/matthewmueller/joy/dom/devicemotioneventinit"
	"github.com/matthewmueller/joy/dom/devicerotationrate"
	"github.com/matthewmueller/joy/dom/devicerotationratedict"
	"github.com/matthewmueller/joy/macro"
)

var _ Event = (*DeviceMotionEvent)(nil)

// NewDeviceMotionEvent fn
func NewDeviceMotionEvent(typearg string, eventinitdict *devicemotioneventinit.DeviceMotionEventInit) *DeviceMotionEvent {
	macro.Rewrite("new DeviceMotionEvent($1, $2)", typearg, eventinitdict)
	return &DeviceMotionEvent{}
}

// DeviceMotionEvent struct
// js:"DeviceMotionEvent,omit"
type DeviceMotionEvent struct {
}

// InitDeviceMotionEvent fn
// js:"initDeviceMotionEvent"
func (*DeviceMotionEvent) InitDeviceMotionEvent(kind string, bubbles bool, cancelable bool, acceleration *deviceaccelerationdict.DeviceAccelerationDict, accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict, rotationRate *devicerotationratedict.DeviceRotationRateDict, interval float32) {
	macro.Rewrite("$_.initDeviceMotionEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, acceleration, accelerationIncludingGravity, rotationRate, interval)
}

// InitEvent fn
// js:"initEvent"
func (*DeviceMotionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*DeviceMotionEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*DeviceMotionEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*DeviceMotionEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Acceleration prop
// js:"acceleration"
func (*DeviceMotionEvent) Acceleration() (acceleration *deviceacceleration.DeviceAcceleration) {
	macro.Rewrite("$_.acceleration")
	return acceleration
}

// AccelerationIncludingGravity prop
// js:"accelerationIncludingGravity"
func (*DeviceMotionEvent) AccelerationIncludingGravity() (accelerationIncludingGravity *deviceacceleration.DeviceAcceleration) {
	macro.Rewrite("$_.accelerationIncludingGravity")
	return accelerationIncludingGravity
}

// Interval prop
// js:"interval"
func (*DeviceMotionEvent) Interval() (interval float32) {
	macro.Rewrite("$_.interval")
	return interval
}

// RotationRate prop
// js:"rotationRate"
func (*DeviceMotionEvent) RotationRate() (rotationRate *devicerotationrate.DeviceRotationRate) {
	macro.Rewrite("$_.rotationRate")
	return rotationRate
}

// Bubbles prop
// js:"bubbles"
func (*DeviceMotionEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*DeviceMotionEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*DeviceMotionEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*DeviceMotionEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*DeviceMotionEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*DeviceMotionEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*DeviceMotionEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*DeviceMotionEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*DeviceMotionEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*DeviceMotionEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*DeviceMotionEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*DeviceMotionEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*DeviceMotionEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*DeviceMotionEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
