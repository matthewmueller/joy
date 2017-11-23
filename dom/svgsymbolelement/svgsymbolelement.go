package svgsymbolelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom/svganimatedrect"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGSymbolElement struct
// js:"SVGSymbolElement,omit"
type SVGSymbolElement struct {
	window.SVGElement
}

// PreserveAspectRatio prop
func (*SVGSymbolElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
func (*SVGSymbolElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}
