package svgfemorphologyelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEMorphologyElement struct
// js:"SVGFEMorphologyElement,omit"
type SVGFEMorphologyElement struct {
	window.SVGElement
}

// Height
func (*SVGFEMorphologyElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEMorphologyElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEMorphologyElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEMorphologyElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEMorphologyElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// In1
func (*SVGFEMorphologyElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// Operator
func (*SVGFEMorphologyElement) Operator() (operator *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Operator")
	return operator
}

// RadiusX
func (*SVGFEMorphologyElement) RadiusX() (radiusX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.RadiusX")
	return radiusX
}

// RadiusY
func (*SVGFEMorphologyElement) RadiusY() (radiusY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.RadiusY")
	return radiusY
}
