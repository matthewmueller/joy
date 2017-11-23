package svgfeoffsetelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEOffsetElement,omit"
type SVGFEOffsetElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// Dx
func (*SVGFEOffsetElement) Dx() (dx *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Dx")
	return dx
}

// Dy
func (*SVGFEOffsetElement) Dy() (dy *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Dy")
	return dy
}

// In1
func (*SVGFEOffsetElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}
