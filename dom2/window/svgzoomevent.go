package window

import (
	"github.com/matthewmueller/golly/dom2/svgpoint"
	"github.com/matthewmueller/golly/dom2/svgrect"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGZoomEvent,omit"
type SVGZoomEvent struct {
	UIEvent
}

// NewScale
func (*SVGZoomEvent) NewScale() (newScale float32) {
	js.Rewrite("$<.NewScale")
	return newScale
}

// NewTranslate
func (*SVGZoomEvent) NewTranslate() (newTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.NewTranslate")
	return newTranslate
}

// PreviousScale
func (*SVGZoomEvent) PreviousScale() (previousScale float32) {
	js.Rewrite("$<.PreviousScale")
	return previousScale
}

// PreviousTranslate
func (*SVGZoomEvent) PreviousTranslate() (previousTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.PreviousTranslate")
	return previousTranslate
}

// ZoomRectScreen
func (*SVGZoomEvent) ZoomRectScreen() (zoomRectScreen *svgrect.SVGRect) {
	js.Rewrite("$<.ZoomRectScreen")
	return zoomRectScreen
}
