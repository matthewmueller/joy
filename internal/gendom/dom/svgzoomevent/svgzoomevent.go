package svgzoomevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpoint"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgrect"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/js"
)

type SVGZoomEvent struct {
	*uievent.UIEvent
}

func (*SVGZoomEvent) GetNewScale() (newScale float32) {
	js.Rewrite("$<.newScale")
	return newScale
}

func (*SVGZoomEvent) GetNewTranslate() (newTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.newTranslate")
	return newTranslate
}

func (*SVGZoomEvent) GetPreviousScale() (previousScale float32) {
	js.Rewrite("$<.previousScale")
	return previousScale
}

func (*SVGZoomEvent) GetPreviousTranslate() (previousTranslate *svgpoint.SVGPoint) {
	js.Rewrite("$<.previousTranslate")
	return previousTranslate
}

func (*SVGZoomEvent) GetZoomRectScreen() (zoomRectScreen *svgrect.SVGRect) {
	js.Rewrite("$<.zoomRectScreen")
	return zoomRectScreen
}
