package window

import "github.com/matthewmueller/joy/macro"

var _ Event = (*ServiceWorkerMessageEvent)(nil)

// NewServiceWorkerMessageEvent fn
func NewServiceWorkerMessageEvent(kind string, eventinitdict *ServiceWorkerMessageEventInit) *ServiceWorkerMessageEvent {
	macro.Rewrite("new ServiceWorkerMessageEvent($1, $2)", kind, eventinitdict)
	return &ServiceWorkerMessageEvent{}
}

// ServiceWorkerMessageEvent struct
// js:"ServiceWorkerMessageEvent,omit"
type ServiceWorkerMessageEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*ServiceWorkerMessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*ServiceWorkerMessageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*ServiceWorkerMessageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*ServiceWorkerMessageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Data prop
// js:"data"
func (*ServiceWorkerMessageEvent) Data() (data interface{}) {
	macro.Rewrite("$_.data")
	return data
}

// LastEventID prop
// js:"lastEventId"
func (*ServiceWorkerMessageEvent) LastEventID() (lastEventId string) {
	macro.Rewrite("$_.lastEventId")
	return lastEventId
}

// Origin prop
// js:"origin"
func (*ServiceWorkerMessageEvent) Origin() (origin string) {
	macro.Rewrite("$_.origin")
	return origin
}

// Ports prop
// js:"ports"
func (*ServiceWorkerMessageEvent) Ports() (ports []*MessagePort) {
	macro.Rewrite("$_.ports")
	return ports
}

// Source prop
// js:"source"
func (*ServiceWorkerMessageEvent) Source() (source interface{}) {
	macro.Rewrite("$_.source")
	return source
}

// Bubbles prop
// js:"bubbles"
func (*ServiceWorkerMessageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*ServiceWorkerMessageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*ServiceWorkerMessageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*ServiceWorkerMessageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*ServiceWorkerMessageEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*ServiceWorkerMessageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*ServiceWorkerMessageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*ServiceWorkerMessageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*ServiceWorkerMessageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*ServiceWorkerMessageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*ServiceWorkerMessageEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*ServiceWorkerMessageEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*ServiceWorkerMessageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*ServiceWorkerMessageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
