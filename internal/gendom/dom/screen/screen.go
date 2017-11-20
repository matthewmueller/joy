package screen

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/js"
)

type Screen struct {
	*eventtarget.EventTarget
}

func (*Screen) MsLockOrientation(orientations interface{}) (b bool) {
	js.Rewrite("$<.msLockOrientation($1)", orientations)
	return b
}

func (*Screen) MsUnlockOrientation() {
	js.Rewrite("$<.msUnlockOrientation()")
}

func (*Screen) GetAvailHeight() (availHeight uint) {
	js.Rewrite("$<.availHeight")
	return availHeight
}

func (*Screen) GetAvailWidth() (availWidth uint) {
	js.Rewrite("$<.availWidth")
	return availWidth
}

func (*Screen) GetBufferDepth() (bufferDepth int) {
	js.Rewrite("$<.bufferDepth")
	return bufferDepth
}

func (*Screen) SetBufferDepth(bufferDepth int) {
	js.Rewrite("$<.bufferDepth = $1", bufferDepth)
}

func (*Screen) GetColorDepth() (colorDepth uint) {
	js.Rewrite("$<.colorDepth")
	return colorDepth
}

func (*Screen) GetDeviceXDPI() (deviceXDPI int) {
	js.Rewrite("$<.deviceXDPI")
	return deviceXDPI
}

func (*Screen) GetDeviceYDPI() (deviceYDPI int) {
	js.Rewrite("$<.deviceYDPI")
	return deviceYDPI
}

func (*Screen) GetFontSmoothingEnabled() (fontSmoothingEnabled bool) {
	js.Rewrite("$<.fontSmoothingEnabled")
	return fontSmoothingEnabled
}

func (*Screen) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*Screen) GetLogicalXDPI() (logicalXDPI int) {
	js.Rewrite("$<.logicalXDPI")
	return logicalXDPI
}

func (*Screen) GetLogicalYDPI() (logicalYDPI int) {
	js.Rewrite("$<.logicalYDPI")
	return logicalYDPI
}

func (*Screen) GetMsOrientation() (msOrientation string) {
	js.Rewrite("$<.msOrientation")
	return msOrientation
}

func (*Screen) GetOnmsorientationchange() (MSOrientationChange *event.Event) {
	js.Rewrite("$<.onmsorientationchange")
	return MSOrientationChange
}

func (*Screen) SetOnmsorientationchange(MSOrientationChange *event.Event) {
	js.Rewrite("$<.onmsorientationchange = $1", MSOrientationChange)
}

func (*Screen) GetPixelDepth() (pixelDepth uint) {
	js.Rewrite("$<.pixelDepth")
	return pixelDepth
}

func (*Screen) GetSystemXDPI() (systemXDPI int) {
	js.Rewrite("$<.systemXDPI")
	return systemXDPI
}

func (*Screen) GetSystemYDPI() (systemYDPI int) {
	js.Rewrite("$<.systemYDPI")
	return systemYDPI
}

func (*Screen) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}
