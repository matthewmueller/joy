package svgfecolormatrixelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEColorMatrixElement struct
// js:"SVGFEColorMatrixElement,omit"
type SVGFEColorMatrixElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFEColorMatrixElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEColorMatrixElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEColorMatrixElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEColorMatrixElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEColorMatrixElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// In1 prop
func (*SVGFEColorMatrixElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// Type prop
func (*SVGFEColorMatrixElement) Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}

// Values prop
func (*SVGFEColorMatrixElement) Values() (values *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.values")
	return values
}
