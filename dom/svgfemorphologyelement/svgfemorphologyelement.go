package svgfemorphologyelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEMorphologyElement struct
// js:"SVGFEMorphologyElement,omit"
type SVGFEMorphologyElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFEMorphologyElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEMorphologyElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEMorphologyElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEMorphologyElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEMorphologyElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// In1 prop
func (*SVGFEMorphologyElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// Operator prop
func (*SVGFEMorphologyElement) Operator() (operator *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.operator")
	return operator
}

// RadiusX prop
func (*SVGFEMorphologyElement) RadiusX() (radiusX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.radiusX")
	return radiusX
}

// RadiusY prop
func (*SVGFEMorphologyElement) RadiusY() (radiusY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.radiusY")
	return radiusY
}
