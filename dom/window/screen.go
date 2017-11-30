package window

import "github.com/matthewmueller/golly/js"

var _ EventTarget = (*Screen)(nil)

// Screen struct
// js:"Screen,omit"
type Screen struct {
}

// MsLockOrientation fn
// js:"msLockOrientation"
func (*Screen) MsLockOrientation(orientations interface{}) (b bool) {
	js.Rewrite("$_.msLockOrientation($1)", orientations)
	return b
}

// MsUnlockOrientation fn
// js:"msUnlockOrientation"
func (*Screen) MsUnlockOrientation() {
	js.Rewrite("$_.msUnlockOrientation()")
}

// AddEventListener fn
// js:"addEventListener"
func (*Screen) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Screen) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Screen) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// AvailHeight prop
// js:"availHeight"
func (*Screen) AvailHeight() (availHeight uint) {
	js.Rewrite("$_.availHeight")
	return availHeight
}

// AvailWidth prop
// js:"availWidth"
func (*Screen) AvailWidth() (availWidth uint) {
	js.Rewrite("$_.availWidth")
	return availWidth
}

// BufferDepth prop
// js:"bufferDepth"
func (*Screen) BufferDepth() (bufferDepth int) {
	js.Rewrite("$_.bufferDepth")
	return bufferDepth
}

// SetBufferDepth prop
// js:"bufferDepth"
func (*Screen) SetBufferDepth(bufferDepth int) {
	js.Rewrite("$_.bufferDepth = $1", bufferDepth)
}

// ColorDepth prop
// js:"colorDepth"
func (*Screen) ColorDepth() (colorDepth uint) {
	js.Rewrite("$_.colorDepth")
	return colorDepth
}

// DeviceXDPI prop
// js:"deviceXDPI"
func (*Screen) DeviceXDPI() (deviceXDPI int) {
	js.Rewrite("$_.deviceXDPI")
	return deviceXDPI
}

// DeviceYDPI prop
// js:"deviceYDPI"
func (*Screen) DeviceYDPI() (deviceYDPI int) {
	js.Rewrite("$_.deviceYDPI")
	return deviceYDPI
}

// FontSmoothingEnabled prop
// js:"fontSmoothingEnabled"
func (*Screen) FontSmoothingEnabled() (fontSmoothingEnabled bool) {
	js.Rewrite("$_.fontSmoothingEnabled")
	return fontSmoothingEnabled
}

// Height prop
// js:"height"
func (*Screen) Height() (height uint) {
	js.Rewrite("$_.height")
	return height
}

// LogicalXDPI prop
// js:"logicalXDPI"
func (*Screen) LogicalXDPI() (logicalXDPI int) {
	js.Rewrite("$_.logicalXDPI")
	return logicalXDPI
}

// LogicalYDPI prop
// js:"logicalYDPI"
func (*Screen) LogicalYDPI() (logicalYDPI int) {
	js.Rewrite("$_.logicalYDPI")
	return logicalYDPI
}

// MsOrientation prop
// js:"msOrientation"
func (*Screen) MsOrientation() (msOrientation string) {
	js.Rewrite("$_.msOrientation")
	return msOrientation
}

// Onmsorientationchange prop
// js:"onmsorientationchange"
func (*Screen) Onmsorientationchange() (onmsorientationchange func(Event)) {
	js.Rewrite("$_.onmsorientationchange")
	return onmsorientationchange
}

// SetOnmsorientationchange prop
// js:"onmsorientationchange"
func (*Screen) SetOnmsorientationchange(onmsorientationchange func(Event)) {
	js.Rewrite("$_.onmsorientationchange = $1", onmsorientationchange)
}

// PixelDepth prop
// js:"pixelDepth"
func (*Screen) PixelDepth() (pixelDepth uint) {
	js.Rewrite("$_.pixelDepth")
	return pixelDepth
}

// SystemXDPI prop
// js:"systemXDPI"
func (*Screen) SystemXDPI() (systemXDPI int) {
	js.Rewrite("$_.systemXDPI")
	return systemXDPI
}

// SystemYDPI prop
// js:"systemYDPI"
func (*Screen) SystemYDPI() (systemYDPI int) {
	js.Rewrite("$_.systemYDPI")
	return systemYDPI
}

// Width prop
// js:"width"
func (*Screen) Width() (width uint) {
	js.Rewrite("$_.width")
	return width
}
