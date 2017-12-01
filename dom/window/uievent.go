package window

// UIEvent interface
// js:"UIEvent"
type UIEvent interface {
	Event

	// InitUIEvent
	// js:"initUIEvent"
	// jsrewrite:"$_.initUIEvent($1, $2, $3, $4, $5)"
	InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int)

	// Detail prop
	// js:"detail"
	// jsrewrite:"$_.detail"
	Detail() (detail int)

	// View prop
	// js:"view"
	// jsrewrite:"$_.view"
	View() (view *Window)
}
