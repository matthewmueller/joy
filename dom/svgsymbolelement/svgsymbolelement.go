package svgsymbolelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedrect"
	"github.com/matthewmueller/golly/dom2/window"
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
