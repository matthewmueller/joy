package window

// Event interface
// js:"Event"
type Event interface {

	// InitEvent
	// js:"initEvent"
	// jsrewrite:"$_.initEvent($1, $2, $3)"
	InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool)

	// PreventDefault
	// js:"preventDefault"
	// jsrewrite:"$_.preventDefault()"
	PreventDefault()

	// StopImmediatePropagation
	// js:"stopImmediatePropagation"
	// jsrewrite:"$_.stopImmediatePropagation()"
	StopImmediatePropagation()

	// StopPropagation
	// js:"stopPropagation"
	// jsrewrite:"$_.stopPropagation()"
	StopPropagation()

	// Bubbles prop
	// js:"bubbles"
	// jsrewrite:"$_.bubbles"
	Bubbles() (bubbles bool)

	// Cancelable prop
	// js:"cancelable"
	// jsrewrite:"$_.cancelable"
	Cancelable() (cancelable bool)

	// CancelBubble prop
	// js:"cancelBubble"
	// jsrewrite:"$_.cancelBubble"
	CancelBubble() (cancelBubble bool)

	// SetCancelBubble prop
	// js:"cancelBubble"
	// jsrewrite:"$_.cancelBubble = $1"
	SetCancelBubble(cancelBubble bool)

	// CurrentTarget prop
	// js:"currentTarget"
	// jsrewrite:"$_.currentTarget"
	CurrentTarget() (currentTarget EventTarget)

	// DefaultPrevented prop
	// js:"defaultPrevented"
	// jsrewrite:"$_.defaultPrevented"
	DefaultPrevented() (defaultPrevented bool)

	// EventPhase prop
	// js:"eventPhase"
	// jsrewrite:"$_.eventPhase"
	EventPhase() (eventPhase uint8)

	// IsTrusted prop
	// js:"isTrusted"
	// jsrewrite:"$_.isTrusted"
	IsTrusted() (isTrusted bool)

	// ReturnValue prop
	// js:"returnValue"
	// jsrewrite:"$_.returnValue"
	ReturnValue() (returnValue bool)

	// SetReturnValue prop
	// js:"returnValue"
	// jsrewrite:"$_.returnValue = $1"
	SetReturnValue(returnValue bool)

	// SrcElement prop
	// js:"srcElement"
	// jsrewrite:"$_.srcElement"
	SrcElement() (srcElement Element)

	// Target prop
	// js:"target"
	// jsrewrite:"$_.target"
	Target() (target EventTarget)

	// TimeStamp prop
	// js:"timeStamp"
	// jsrewrite:"$_.timeStamp"
	TimeStamp() (timeStamp uint64)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind string)
}
