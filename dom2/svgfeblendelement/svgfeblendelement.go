package svgfeblendelement

import "github.com/matthewmueller/golly/js"

// js:"SVGFEBlendElement,omit"
type SVGFEBlendElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// In1
func (*SVGFEBlendElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// In2
func (*SVGFEBlendElement) In2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In2")
	return in2
}

// Mode
func (*SVGFEBlendElement) Mode() (mode *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Mode")
	return mode
}
