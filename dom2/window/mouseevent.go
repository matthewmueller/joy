package window

// js:"MouseEvent,omit"
type MouseEvent interface {
	UIEvent

	// GetModifierState
	GetModifierState(keyArg string) (b bool)

	// InitMouseEvent
	InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget)

	// AltKey
	AltKey() (altKey bool)

	// Button
	Button() (button uint8)

	// Buttons
	Buttons() (buttons uint8)

	// ClientX
	ClientX() (clientX int)

	// ClientY
	ClientY() (clientY int)

	// CtrlKey
	CtrlKey() (ctrlKey bool)

	// FromElement
	FromElement() (fromElement Element)

	// LayerX
	LayerX() (layerX int)

	// LayerY
	LayerY() (layerY int)

	// MetaKey
	MetaKey() (metaKey bool)

	// MovementX
	MovementX() (movementX int)

	// MovementY
	MovementY() (movementY int)

	// OffsetX
	OffsetX() (offsetX int)

	// OffsetY
	OffsetY() (offsetY int)

	// PageX
	PageX() (pageX int)

	// PageY
	PageY() (pageY int)

	// RelatedTarget
	RelatedTarget() (relatedTarget EventTarget)

	// ScreenX
	ScreenX() (screenX int)

	// ScreenY
	ScreenY() (screenY int)

	// ShiftKey
	ShiftKey() (shiftKey bool)

	// ToElement
	ToElement() (toElement Element)

	// Which
	Which() (which uint8)

	// X
	X() (x int)

	// Y
	Y() (y int)
}
