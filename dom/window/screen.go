package window

import "github.com/matthewmueller/golly/js"

// Screen struct
// js:"Screen,omit"
type Screen struct {
	EventTarget
}

// MsLockOrientation fn
func (*Screen) MsLockOrientation(orientations interface{}) (b bool) {
	js.Rewrite("$<.msLockOrientation($1)", orientations)
	return b
}

// MsUnlockOrientation fn
func (*Screen) MsUnlockOrientation() {
	js.Rewrite("$<.msUnlockOrientation()")
}

// AvailHeight prop
func (*Screen) AvailHeight() (availHeight uint) {
	js.Rewrite("$<.availHeight")
	return availHeight
}

// AvailWidth prop
func (*Screen) AvailWidth() (availWidth uint) {
	js.Rewrite("$<.availWidth")
	return availWidth
}

// BufferDepth prop
func (*Screen) BufferDepth() (bufferDepth int) {
	js.Rewrite("$<.bufferDepth")
	return bufferDepth
}

// BufferDepth prop
func (*Screen) SetBufferDepth(bufferDepth int) {
	js.Rewrite("$<.bufferDepth = $1", bufferDepth)
}

// ColorDepth prop
func (*Screen) ColorDepth() (colorDepth uint) {
	js.Rewrite("$<.colorDepth")
	return colorDepth
}

// DeviceXDPI prop
func (*Screen) DeviceXDPI() (deviceXDPI int) {
	js.Rewrite("$<.deviceXDPI")
	return deviceXDPI
}

// DeviceYDPI prop
func (*Screen) DeviceYDPI() (deviceYDPI int) {
	js.Rewrite("$<.deviceYDPI")
	return deviceYDPI
}

// FontSmoothingEnabled prop
func (*Screen) FontSmoothingEnabled() (fontSmoothingEnabled bool) {
	js.Rewrite("$<.fontSmoothingEnabled")
	return fontSmoothingEnabled
}

// Height prop
func (*Screen) Height() (height uint) {
	js.Rewrite("$<.height")
	return height
}

// LogicalXDPI prop
func (*Screen) LogicalXDPI() (logicalXDPI int) {
	js.Rewrite("$<.logicalXDPI")
	return logicalXDPI
}

// LogicalYDPI prop
func (*Screen) LogicalYDPI() (logicalYDPI int) {
	js.Rewrite("$<.logicalYDPI")
	return logicalYDPI
}

// MsOrientation prop
func (*Screen) MsOrientation() (msOrientation string) {
	js.Rewrite("$<.msOrientation")
	return msOrientation
}

// Onmsorientationchange prop
func (*Screen) Onmsorientationchange() (onmsorientationchange func(Event)) {
	js.Rewrite("$<.onmsorientationchange")
	return onmsorientationchange
}

// Onmsorientationchange prop
func (*Screen) SetOnmsorientationchange(onmsorientationchange func(Event)) {
	js.Rewrite("$<.onmsorientationchange = $1", onmsorientationchange)
}

// PixelDepth prop
func (*Screen) PixelDepth() (pixelDepth uint) {
	js.Rewrite("$<.pixelDepth")
	return pixelDepth
}

// SystemXDPI prop
func (*Screen) SystemXDPI() (systemXDPI int) {
	js.Rewrite("$<.systemXDPI")
	return systemXDPI
}

// SystemYDPI prop
func (*Screen) SystemYDPI() (systemYDPI int) {
	js.Rewrite("$<.systemYDPI")
	return systemYDPI
}

// Width prop
func (*Screen) Width() (width uint) {
	js.Rewrite("$<.width")
	return width
}
