package window

import (
	"github.com/matthewmueller/joy/dom/datatransfer"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/macro"
)

var _ MouseEvent = (*DragEvent)(nil)
var _ UIEvent = (*DragEvent)(nil)
var _ Event = (*DragEvent)(nil)

// DragEvent struct
// js:"DragEvent,omit"
type DragEvent struct {
}

// InitDragEvent fn
// js:"initDragEvent"
func (*DragEvent) InitDragEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget, dataTransferArg *datatransfer.DataTransfer) {
	macro.Rewrite("$_.initDragEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, dataTransferArg)
}

// MsConvertURL fn
// js:"msConvertURL"
func (*DragEvent) MsConvertURL(file *file.File, targetType string, targetURL *string) {
	macro.Rewrite("$_.msConvertURL($1, $2, $3)", file, targetType, targetURL)
}

// GetModifierState fn
// js:"getModifierState"
func (*DragEvent) GetModifierState(keyArg string) (b bool) {
	macro.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

// InitMouseEvent fn
// js:"initMouseEvent"
func (*DragEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget) {
	macro.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*DragEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*DragEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*DragEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*DragEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*DragEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// DataTransfer prop
// js:"dataTransfer"
func (*DragEvent) DataTransfer() (dataTransfer *datatransfer.DataTransfer) {
	macro.Rewrite("$_.dataTransfer")
	return dataTransfer
}

// AltKey prop
// js:"altKey"
func (*DragEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

// Button prop
// js:"button"
func (*DragEvent) Button() (button uint8) {
	macro.Rewrite("$_.button")
	return button
}

// Buttons prop
// js:"buttons"
func (*DragEvent) Buttons() (buttons uint8) {
	macro.Rewrite("$_.buttons")
	return buttons
}

// ClientX prop
// js:"clientX"
func (*DragEvent) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

// ClientY prop
// js:"clientY"
func (*DragEvent) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

// CtrlKey prop
// js:"ctrlKey"
func (*DragEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

// FromElement prop
// js:"fromElement"
func (*DragEvent) FromElement() (fromElement Element) {
	macro.Rewrite("$_.fromElement")
	return fromElement
}

// LayerX prop
// js:"layerX"
func (*DragEvent) LayerX() (layerX int) {
	macro.Rewrite("$_.layerX")
	return layerX
}

// LayerY prop
// js:"layerY"
func (*DragEvent) LayerY() (layerY int) {
	macro.Rewrite("$_.layerY")
	return layerY
}

// MetaKey prop
// js:"metaKey"
func (*DragEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

// MovementX prop
// js:"movementX"
func (*DragEvent) MovementX() (movementX int) {
	macro.Rewrite("$_.movementX")
	return movementX
}

// MovementY prop
// js:"movementY"
func (*DragEvent) MovementY() (movementY int) {
	macro.Rewrite("$_.movementY")
	return movementY
}

// OffsetX prop
// js:"offsetX"
func (*DragEvent) OffsetX() (offsetX int) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

// OffsetY prop
// js:"offsetY"
func (*DragEvent) OffsetY() (offsetY int) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

// PageX prop
// js:"pageX"
func (*DragEvent) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

// PageY prop
// js:"pageY"
func (*DragEvent) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

// RelatedTarget prop
// js:"relatedTarget"
func (*DragEvent) RelatedTarget() (relatedTarget EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

// ScreenX prop
// js:"screenX"
func (*DragEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*DragEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

// ShiftKey prop
// js:"shiftKey"
func (*DragEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

// ToElement prop
// js:"toElement"
func (*DragEvent) ToElement() (toElement Element) {
	macro.Rewrite("$_.toElement")
	return toElement
}

// Which prop
// js:"which"
func (*DragEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

// X prop
// js:"x"
func (*DragEvent) X() (x int) {
	macro.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*DragEvent) Y() (y int) {
	macro.Rewrite("$_.y")
	return y
}

// Detail prop
// js:"detail"
func (*DragEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*DragEvent) View() (view *Window) {
	macro.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*DragEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*DragEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*DragEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*DragEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*DragEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*DragEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*DragEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*DragEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*DragEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*DragEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*DragEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*DragEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*DragEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*DragEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
