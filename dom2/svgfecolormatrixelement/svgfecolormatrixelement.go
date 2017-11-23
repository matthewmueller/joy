package svgfecolormatrixelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEColorMatrixElement,omit"
type SVGFEColorMatrixElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// In1
func (*SVGFEColorMatrixElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// Type
func (*SVGFEColorMatrixElement) Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Type")
	return kind
}

// Values
func (*SVGFEColorMatrixElement) Values() (values *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.Values")
	return values
}
