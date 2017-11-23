package window

import "github.com/matthewmueller/golly/js"

// js:"Screen,omit"
type Screen struct {
	EventTarget
}

// MsLockOrientation
func (*Screen) MsLockOrientation(orientations interface{}) (b bool) {
	js.Rewrite("$<.MsLockOrientation($1)", orientations)
	return b
}

// MsUnlockOrientation
func (*Screen) MsUnlockOrientation() {
	js.Rewrite("$<.MsUnlockOrientation()")
}

// AvailHeight
func (*Screen) AvailHeight() (availHeight uint) {
	js.Rewrite("$<.AvailHeight")
	return availHeight
}

// AvailWidth
func (*Screen) AvailWidth() (availWidth uint) {
	js.Rewrite("$<.AvailWidth")
	return availWidth
}

// BufferDepth
func (*Screen) BufferDepth() (bufferDepth int) {
	js.Rewrite("$<.BufferDepth")
	return bufferDepth
}

// BufferDepth
func (*Screen) SetBufferDepth(bufferDepth int) {
	js.Rewrite("$<.BufferDepth = $1", bufferDepth)
}

// ColorDepth
func (*Screen) ColorDepth() (colorDepth uint) {
	js.Rewrite("$<.ColorDepth")
	return colorDepth
}

// DeviceXDPI
func (*Screen) DeviceXDPI() (deviceXDPI int) {
	js.Rewrite("$<.DeviceXDPI")
	return deviceXDPI
}

// DeviceYDPI
func (*Screen) DeviceYDPI() (deviceYDPI int) {
	js.Rewrite("$<.DeviceYDPI")
	return deviceYDPI
}

// FontSmoothingEnabled
func (*Screen) FontSmoothingEnabled() (fontSmoothingEnabled bool) {
	js.Rewrite("$<.FontSmoothingEnabled")
	return fontSmoothingEnabled
}

// Height
func (*Screen) Height() (height uint) {
	js.Rewrite("$<.Height")
	return height
}

// LogicalXDPI
func (*Screen) LogicalXDPI() (logicalXDPI int) {
	js.Rewrite("$<.LogicalXDPI")
	return logicalXDPI
}

// LogicalYDPI
func (*Screen) LogicalYDPI() (logicalYDPI int) {
	js.Rewrite("$<.LogicalYDPI")
	return logicalYDPI
}

// MsOrientation
func (*Screen) MsOrientation() (msOrientation string) {
	js.Rewrite("$<.MsOrientation")
	return msOrientation
}

// Onmsorientationchange
func (*Screen) Onmsorientationchange() (onmsorientationchange func(Event)) {
	js.Rewrite("$<.Onmsorientationchange")
	return onmsorientationchange
}

// Onmsorientationchange
func (*Screen) SetOnmsorientationchange(onmsorientationchange func(Event)) {
	js.Rewrite("$<.Onmsorientationchange = $1", onmsorientationchange)
}

// PixelDepth
func (*Screen) PixelDepth() (pixelDepth uint) {
	js.Rewrite("$<.PixelDepth")
	return pixelDepth
}

// SystemXDPI
func (*Screen) SystemXDPI() (systemXDPI int) {
	js.Rewrite("$<.SystemXDPI")
	return systemXDPI
}

// SystemYDPI
func (*Screen) SystemYDPI() (systemYDPI int) {
	js.Rewrite("$<.SystemYDPI")
	return systemYDPI
}

// Width
func (*Screen) Width() (width uint) {
	js.Rewrite("$<.Width")
	return width
}
