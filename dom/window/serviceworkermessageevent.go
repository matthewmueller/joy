package window

import "github.com/matthewmueller/joy/js"

var _ Event = (*ServiceWorkerMessageEvent)(nil)

// NewServiceWorkerMessageEvent fn
func NewServiceWorkerMessageEvent(kind string, eventinitdict *ServiceWorkerMessageEventInit) *ServiceWorkerMessageEvent {
	js.Rewrite("ServiceWorkerMessageEvent")
	return &ServiceWorkerMessageEvent{}
}

// ServiceWorkerMessageEvent struct
// js:"ServiceWorkerMessageEvent,omit"
type ServiceWorkerMessageEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*ServiceWorkerMessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*ServiceWorkerMessageEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*ServiceWorkerMessageEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*ServiceWorkerMessageEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Data prop
// js:"data"
func (*ServiceWorkerMessageEvent) Data() (data interface{}) {
	js.Rewrite("$_.data")
	return data
}

// LastEventID prop
// js:"lastEventId"
func (*ServiceWorkerMessageEvent) LastEventID() (lastEventId string) {
	js.Rewrite("$_.lastEventId")
	return lastEventId
}

// Origin prop
// js:"origin"
func (*ServiceWorkerMessageEvent) Origin() (origin string) {
	js.Rewrite("$_.origin")
	return origin
}

// Ports prop
// js:"ports"
func (*ServiceWorkerMessageEvent) Ports() (ports []*MessagePort) {
	js.Rewrite("$_.ports")
	return ports
}

// Source prop
// js:"source"
func (*ServiceWorkerMessageEvent) Source() (source interface{}) {
	js.Rewrite("$_.source")
	return source
}

// Bubbles prop
// js:"bubbles"
func (*ServiceWorkerMessageEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*ServiceWorkerMessageEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*ServiceWorkerMessageEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*ServiceWorkerMessageEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*ServiceWorkerMessageEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*ServiceWorkerMessageEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*ServiceWorkerMessageEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*ServiceWorkerMessageEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*ServiceWorkerMessageEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*ServiceWorkerMessageEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*ServiceWorkerMessageEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*ServiceWorkerMessageEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*ServiceWorkerMessageEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*ServiceWorkerMessageEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
