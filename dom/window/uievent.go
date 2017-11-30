package window

// UIEvent interface
// js:"UIEvent"
type UIEvent interface {
	Event

	// InitUIEvent
	// js:"initUIEvent"
	InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int)

	// Detail prop
	// js:"detail"
	Detail() (detail int)

	// View prop
	// js:"view"
	View() (view *Window)
}
