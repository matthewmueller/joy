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
