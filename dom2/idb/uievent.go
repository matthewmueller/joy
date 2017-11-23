package idb

// js:"UIEvent,omit"
type UIEvent interface {
	Event

	// InitUIEvent
	InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int)

	// Detail
	Detail() (detail int)

	// View
	View() (view *Window)
}
