package window

import (
	"github.com/matthewmueller/golly/dom2/svgpoint"
	"github.com/matthewmueller/golly/dom2/svgrect"
	"github.com/matthewmueller/golly/js"
)

// SVGZoomEvent struct
// js:"SVGZoomEvent,omit"
type SVGZoomEvent struct {
	UIEvent
}

// NewScale prop
func (*SVGZoomEvent) NewScale() (newScale float32) {
	js.Rewrite("$<.newScale")
	return newScale
}

// NewTranslate prop
func (*SVGZoomEvent) NewTranslate() (newTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.newTranslate")
	return newTranslate
}

// PreviousScale prop
func (*SVGZoomEvent) PreviousScale() (previousScale float32) {
	js.Rewrite("$<.previousScale")
	return previousScale
}

// PreviousTranslate prop
func (*SVGZoomEvent) PreviousTranslate() (previousTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.previousTranslate")
	return previousTranslate
}

// ZoomRectScreen prop
func (*SVGZoomEvent) ZoomRectScreen() (zoomRectScreen *svgrect.SVGRect) {
	js.Rewrite("$<.zoomRectScreen")
	return zoomRectScreen
}
