package svgfemorphologyelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEMorphologyElement,omit"
type SVGFEMorphologyElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
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
