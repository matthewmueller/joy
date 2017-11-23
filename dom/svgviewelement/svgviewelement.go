package svgviewelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom/svganimatedrect"
	"github.com/matthewmueller/golly/dom/svgstringlist"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGViewElement struct
// js:"SVGViewElement,omit"
type SVGViewElement struct {
	window.SVGElement
}

// ZoomAndPan prop
func (*SVGViewElement) ZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$<.zoomAndPan")
	return zoomAndPan
}

// PreserveAspectRatio prop
func (*SVGViewElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
func (*SVGViewElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}

// ViewTarget prop
func (*SVGViewElement) ViewTarget() (viewTarget *svgstringlist.SVGStringList) {
	js.Rewrite("$<.viewTarget")
	return viewTarget
}
