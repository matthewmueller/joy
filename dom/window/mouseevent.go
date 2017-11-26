package window

import "github.com/matthewmueller/golly/js"

// js:"MouseEvent,omit"
type MouseEvent interface {
	UIEvent

	// GetModifierState
	// js:"getModifierState,rewrite=getmodifierstate"
	GetModifierState(keyArg string) (b bool)

	// InitMouseEvent
	// js:"initMouseEvent,rewrite=initmouseevent"
	InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget)

	// AltKey prop
	// js:"altKey,rewrite=altkey"
	AltKey() (altKey bool)

	// Button prop
	// js:"button,rewrite=button"
	Button() (button uint8)

	// Buttons prop
	// js:"buttons,rewrite=buttons"
	Buttons() (buttons uint8)

	// ClientX prop
	// js:"clientX,rewrite=clientx"
	ClientX() (clientX int)

	// ClientY prop
	// js:"clientY,rewrite=clienty"
	ClientY() (clientY int)

	// CtrlKey prop
	// js:"ctrlKey,rewrite=ctrlkey"
	CtrlKey() (ctrlKey bool)

	// FromElement prop
	// js:"fromElement,rewrite=fromelement"
	FromElement() (fromElement Element)

	// LayerX prop
	// js:"layerX,rewrite=layerx"
	LayerX() (layerX int)

	// LayerY prop
	// js:"layerY,rewrite=layery"
	LayerY() (layerY int)

	// MetaKey prop
	// js:"metaKey,rewrite=metakey"
	MetaKey() (metaKey bool)

	// MovementX prop
	// js:"movementX,rewrite=movementx"
	MovementX() (movementX int)

	// MovementY prop
	// js:"movementY,rewrite=movementy"
	MovementY() (movementY int)

	// OffsetX prop
	// js:"offsetX,rewrite=offsetx"
	OffsetX() (offsetX int)

	// OffsetY prop
	// js:"offsetY,rewrite=offsety"
	OffsetY() (offsetY int)

	// PageX prop
	// js:"pageX,rewrite=pagex"
	PageX() (pageX int)

	// PageY prop
	// js:"pageY,rewrite=pagey"
	PageY() (pageY int)

	// RelatedTarget prop
	// js:"relatedTarget,rewrite=relatedtarget"
	RelatedTarget() (relatedTarget EventTarget)

	// ScreenX prop
	// js:"screenX,rewrite=screenx"
	ScreenX() (screenX int)

	// ScreenY prop
	// js:"screenY,rewrite=screeny"
	ScreenY() (screenY int)

	// ShiftKey prop
	// js:"shiftKey,rewrite=shiftkey"
	ShiftKey() (shiftKey bool)

	// ToElement prop
	// js:"toElement,rewrite=toelement"
	ToElement() (toElement Element)

	// Which prop
	// js:"which,rewrite=which"
	Which() (which uint8)

	// X prop
	// js:"x,rewrite=x"
	X() (x int)

	// Y prop
	// js:"y,rewrite=y"
	Y() (y int)
}

// getmodifierstate fn
func getmodifierstate(keyArg string) (b bool) {
	js.Rewrite("$<.getModifierState($1)", keyArg)
	return b
}

// initmouseevent fn
func initmouseevent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget) {
	js.Rewrite("$<.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

// altkey prop
func altkey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

// button prop
func button() (button uint8) {
	js.Rewrite("$<.button")
	return button
}

// buttons prop
func buttons() (buttons uint8) {
	js.Rewrite("$<.buttons")
	return buttons
}

// clientx prop
func clientx() (clientX int) {
	js.Rewrite("$<.clientX")
	return clientX
}

// clienty prop
func clienty() (clientY int) {
	js.Rewrite("$<.clientY")
	return clientY
}

// ctrlkey prop
func ctrlkey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

// fromelement prop
func fromelement() (fromElement Element) {
	js.Rewrite("$<.fromElement")
	return fromElement
}

// layerx prop
func layerx() (layerX int) {
	js.Rewrite("$<.layerX")
	return layerX
}

// layery prop
func layery() (layerY int) {
	js.Rewrite("$<.layerY")
	return layerY
}

// metakey prop
func metakey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

// movementx prop
func movementx() (movementX int) {
	js.Rewrite("$<.movementX")
	return movementX
}

// movementy prop
func movementy() (movementY int) {
	js.Rewrite("$<.movementY")
	return movementY
}

// offsetx prop
func offsetx() (offsetX int) {
	js.Rewrite("$<.offsetX")
	return offsetX
}

// offsety prop
func offsety() (offsetY int) {
	js.Rewrite("$<.offsetY")
	return offsetY
}

// pagex prop
func pagex() (pageX int) {
	js.Rewrite("$<.pageX")
	return pageX
}

// pagey prop
func pagey() (pageY int) {
	js.Rewrite("$<.pageY")
	return pageY
}

// relatedtarget prop
func relatedtarget() (relatedTarget EventTarget) {
	js.Rewrite("$<.relatedTarget")
	return relatedTarget
}

// screenx prop
func screenx() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

// screeny prop
func screeny() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

// shiftkey prop
func shiftkey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

// toelement prop
func toelement() (toElement Element) {
	js.Rewrite("$<.toElement")
	return toElement
}

// which prop
func which() (which uint8) {
	js.Rewrite("$<.which")
	return which
}

// x prop
func x() (x int) {
	js.Rewrite("$<.x")
	return x
}

// y prop
func y() (y int) {
	js.Rewrite("$<.y")
	return y
}
