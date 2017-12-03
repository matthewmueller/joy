package window

import "github.com/matthewmueller/joy/macro"

var _ EventTarget = (*Screen)(nil)

// Screen struct
// js:"Screen,omit"
type Screen struct {
}

// MsLockOrientation fn
// js:"msLockOrientation"
func (*Screen) MsLockOrientation(orientations interface{}) (b bool) {
	macro.Rewrite("$_.msLockOrientation($1)", orientations)
	return b
}

// MsUnlockOrientation fn
// js:"msUnlockOrientation"
func (*Screen) MsUnlockOrientation() {
	macro.Rewrite("$_.msUnlockOrientation()")
}

// AddEventListener fn
// js:"addEventListener"
func (*Screen) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Screen) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Screen) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// AvailHeight prop
// js:"availHeight"
func (*Screen) AvailHeight() (availHeight uint) {
	macro.Rewrite("$_.availHeight")
	return availHeight
}

// AvailWidth prop
// js:"availWidth"
func (*Screen) AvailWidth() (availWidth uint) {
	macro.Rewrite("$_.availWidth")
	return availWidth
}

// BufferDepth prop
// js:"bufferDepth"
func (*Screen) BufferDepth() (bufferDepth int) {
	macro.Rewrite("$_.bufferDepth")
	return bufferDepth
}

// SetBufferDepth prop
// js:"bufferDepth"
func (*Screen) SetBufferDepth(bufferDepth int) {
	macro.Rewrite("$_.bufferDepth = $1", bufferDepth)
}

// ColorDepth prop
// js:"colorDepth"
func (*Screen) ColorDepth() (colorDepth uint) {
	macro.Rewrite("$_.colorDepth")
	return colorDepth
}

// DeviceXDPI prop
// js:"deviceXDPI"
func (*Screen) DeviceXDPI() (deviceXDPI int) {
	macro.Rewrite("$_.deviceXDPI")
	return deviceXDPI
}

// DeviceYDPI prop
// js:"deviceYDPI"
func (*Screen) DeviceYDPI() (deviceYDPI int) {
	macro.Rewrite("$_.deviceYDPI")
	return deviceYDPI
}

// FontSmoothingEnabled prop
// js:"fontSmoothingEnabled"
func (*Screen) FontSmoothingEnabled() (fontSmoothingEnabled bool) {
	macro.Rewrite("$_.fontSmoothingEnabled")
	return fontSmoothingEnabled
}

// Height prop
// js:"height"
func (*Screen) Height() (height uint) {
	macro.Rewrite("$_.height")
	return height
}

// LogicalXDPI prop
// js:"logicalXDPI"
func (*Screen) LogicalXDPI() (logicalXDPI int) {
	macro.Rewrite("$_.logicalXDPI")
	return logicalXDPI
}

// LogicalYDPI prop
// js:"logicalYDPI"
func (*Screen) LogicalYDPI() (logicalYDPI int) {
	macro.Rewrite("$_.logicalYDPI")
	return logicalYDPI
}

// MsOrientation prop
// js:"msOrientation"
func (*Screen) MsOrientation() (msOrientation string) {
	macro.Rewrite("$_.msOrientation")
	return msOrientation
}

// Onmsorientationchange prop
// js:"onmsorientationchange"
func (*Screen) Onmsorientationchange() (onmsorientationchange func(Event)) {
	macro.Rewrite("$_.onmsorientationchange")
	return onmsorientationchange
}

// SetOnmsorientationchange prop
// js:"onmsorientationchange"
func (*Screen) SetOnmsorientationchange(onmsorientationchange func(Event)) {
	macro.Rewrite("$_.onmsorientationchange = $1", onmsorientationchange)
}

// PixelDepth prop
// js:"pixelDepth"
func (*Screen) PixelDepth() (pixelDepth uint) {
	macro.Rewrite("$_.pixelDepth")
	return pixelDepth
}

// SystemXDPI prop
// js:"systemXDPI"
func (*Screen) SystemXDPI() (systemXDPI int) {
	macro.Rewrite("$_.systemXDPI")
	return systemXDPI
}

// SystemYDPI prop
// js:"systemYDPI"
func (*Screen) SystemYDPI() (systemYDPI int) {
	macro.Rewrite("$_.systemYDPI")
	return systemYDPI
}

// Width prop
// js:"width"
func (*Screen) Width() (width uint) {
	macro.Rewrite("$_.width")
	return width
}
