package svgfecolormatrixelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEColorMatrixElement struct
// js:"SVGFEColorMatrixElement,omit"
type SVGFEColorMatrixElement struct {
	window.SVGElement
}

// Height
func (*SVGFEColorMatrixElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEColorMatrixElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEColorMatrixElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEColorMatrixElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEColorMatrixElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
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
