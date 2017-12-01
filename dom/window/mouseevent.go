package window

// MouseEvent interface
// js:"MouseEvent"
type MouseEvent interface {
	UIEvent

	// GetModifierState
	// js:"getModifierState"
	// jsrewrite:"$_.getModifierState($1)"
	GetModifierState(keyArg string) (b bool)

	// InitMouseEvent
	// js:"initMouseEvent"
	// jsrewrite:"$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)"
	InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget)

	// AltKey prop
	// js:"altKey"
	// jsrewrite:"$_.altKey"
	AltKey() (altKey bool)

	// Button prop
	// js:"button"
	// jsrewrite:"$_.button"
	Button() (button uint8)

	// Buttons prop
	// js:"buttons"
	// jsrewrite:"$_.buttons"
	Buttons() (buttons uint8)

	// ClientX prop
	// js:"clientX"
	// jsrewrite:"$_.clientX"
	ClientX() (clientX int)

	// ClientY prop
	// js:"clientY"
	// jsrewrite:"$_.clientY"
	ClientY() (clientY int)

	// CtrlKey prop
	// js:"ctrlKey"
	// jsrewrite:"$_.ctrlKey"
	CtrlKey() (ctrlKey bool)

	// FromElement prop
	// js:"fromElement"
	// jsrewrite:"$_.fromElement"
	FromElement() (fromElement Element)

	// LayerX prop
	// js:"layerX"
	// jsrewrite:"$_.layerX"
	LayerX() (layerX int)

	// LayerY prop
	// js:"layerY"
	// jsrewrite:"$_.layerY"
	LayerY() (layerY int)

	// MetaKey prop
	// js:"metaKey"
	// jsrewrite:"$_.metaKey"
	MetaKey() (metaKey bool)

	// MovementX prop
	// js:"movementX"
	// jsrewrite:"$_.movementX"
	MovementX() (movementX int)

	// MovementY prop
	// js:"movementY"
	// jsrewrite:"$_.movementY"
	MovementY() (movementY int)

	// OffsetX prop
	// js:"offsetX"
	// jsrewrite:"$_.offsetX"
	OffsetX() (offsetX int)

	// OffsetY prop
	// js:"offsetY"
	// jsrewrite:"$_.offsetY"
	OffsetY() (offsetY int)

	// PageX prop
	// js:"pageX"
	// jsrewrite:"$_.pageX"
	PageX() (pageX int)

	// PageY prop
	// js:"pageY"
	// jsrewrite:"$_.pageY"
	PageY() (pageY int)

	// RelatedTarget prop
	// js:"relatedTarget"
	// jsrewrite:"$_.relatedTarget"
	RelatedTarget() (relatedTarget EventTarget)

	// ScreenX prop
	// js:"screenX"
	// jsrewrite:"$_.screenX"
	ScreenX() (screenX int)

	// ScreenY prop
	// js:"screenY"
	// jsrewrite:"$_.screenY"
	ScreenY() (screenY int)

	// ShiftKey prop
	// js:"shiftKey"
	// jsrewrite:"$_.shiftKey"
	ShiftKey() (shiftKey bool)

	// ToElement prop
	// js:"toElement"
	// jsrewrite:"$_.toElement"
	ToElement() (toElement Element)

	// Which prop
	// js:"which"
	// jsrewrite:"$_.which"
	Which() (which uint8)

	// X prop
	// js:"x"
	// jsrewrite:"$_.x"
	X() (x int)

	// Y prop
	// js:"y"
	// jsrewrite:"$_.y"
	Y() (y int)
}
