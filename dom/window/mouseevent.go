package window

// MouseEvent interface
// js:"MouseEvent"
type MouseEvent interface {
	UIEvent

	// GetModifierState
	// js:"getModifierState"
	GetModifierState(keyArg string) (b bool)

	// InitMouseEvent
	// js:"initMouseEvent"
	InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget)

	// AltKey prop
	// js:"altKey"
	AltKey() (altKey bool)

	// Button prop
	// js:"button"
	Button() (button uint8)

	// Buttons prop
	// js:"buttons"
	Buttons() (buttons uint8)

	// ClientX prop
	// js:"clientX"
	ClientX() (clientX int)

	// ClientY prop
	// js:"clientY"
	ClientY() (clientY int)

	// CtrlKey prop
	// js:"ctrlKey"
	CtrlKey() (ctrlKey bool)

	// FromElement prop
	// js:"fromElement"
	FromElement() (fromElement Element)

	// LayerX prop
	// js:"layerX"
	LayerX() (layerX int)

	// LayerY prop
	// js:"layerY"
	LayerY() (layerY int)

	// MetaKey prop
	// js:"metaKey"
	MetaKey() (metaKey bool)

	// MovementX prop
	// js:"movementX"
	MovementX() (movementX int)

	// MovementY prop
	// js:"movementY"
	MovementY() (movementY int)

	// OffsetX prop
	// js:"offsetX"
	OffsetX() (offsetX int)

	// OffsetY prop
	// js:"offsetY"
	OffsetY() (offsetY int)

	// PageX prop
	// js:"pageX"
	PageX() (pageX int)

	// PageY prop
	// js:"pageY"
	PageY() (pageY int)

	// RelatedTarget prop
	// js:"relatedTarget"
	RelatedTarget() (relatedTarget EventTarget)

	// ScreenX prop
	// js:"screenX"
	ScreenX() (screenX int)

	// ScreenY prop
	// js:"screenY"
	ScreenY() (screenY int)

	// ShiftKey prop
	// js:"shiftKey"
	ShiftKey() (shiftKey bool)

	// ToElement prop
	// js:"toElement"
	ToElement() (toElement Element)

	// Which prop
	// js:"which"
	Which() (which uint8)

	// X prop
	// js:"x"
	X() (x int)

	// Y prop
	// js:"y"
	Y() (y int)
}
