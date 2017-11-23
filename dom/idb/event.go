package idb

// js:"Event,omit"
type Event interface {

	// InitEvent
	InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool)

	// PreventDefault
	PreventDefault()

	// StopImmediatePropagation
	StopImmediatePropagation()

	// StopPropagation
	StopPropagation()

	// Bubbles
	Bubbles() (bubbles bool)

	// Cancelable
	Cancelable() (cancelable bool)

	// CancelBubble
	CancelBubble() (cancelBubble bool)

	// CancelBubble
	SetCancelBubble(cancelBubble bool)

	// CurrentTarget
	CurrentTarget() (currentTarget EventTarget)

	// DefaultPrevented
	DefaultPrevented() (defaultPrevented bool)

	// EventPhase
	EventPhase() (eventPhase uint8)

	// IsTrusted
	IsTrusted() (isTrusted bool)

	// ReturnValue
	ReturnValue() (returnValue bool)

	// ReturnValue
	SetReturnValue(returnValue bool)

	// SrcElement
	SrcElement() (srcElement Element)

	// Target
	Target() (target EventTarget)

	// TimeStamp
	TimeStamp() (timeStamp uint64)

	// Type
	Type() (kind string)
}
