package svgviewelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedrect"
	"github.com/matthewmueller/golly/dom2/svgstringlist"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGViewElement struct
// js:"SVGViewElement,omit"
type SVGViewElement struct {
	window.SVGElement
}

// ZoomAndPan
func (*SVGViewElement) ZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$<.ZoomAndPan")
	return zoomAndPan
}

// PreserveAspectRatio
func (*SVGViewElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.PreserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox
func (*SVGViewElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.ViewBox")
	return viewBox
}

// ViewTarget
func (*SVGViewElement) ViewTarget() (viewTarget *svgstringlist.SVGStringList) {
	js.Rewrite("$<.ViewTarget")
	return viewTarget
}
