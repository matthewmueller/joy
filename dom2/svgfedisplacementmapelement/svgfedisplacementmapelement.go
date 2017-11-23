package svgfedisplacementmapelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEDisplacementMapElement,omit"
type SVGFEDisplacementMapElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// In1
func (*SVGFEDisplacementMapElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// In2
func (*SVGFEDisplacementMapElement) In2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In2")
	return in2
}

// Scale
func (*SVGFEDisplacementMapElement) Scale() (scale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Scale")
	return scale
}

// XChannelSelector
func (*SVGFEDisplacementMapElement) XChannelSelector() (xChannelSelector *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.XChannelSelector")
	return xChannelSelector
}

// YChannelSelector
func (*SVGFEDisplacementMapElement) YChannelSelector() (yChannelSelector *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.YChannelSelector")
	return yChannelSelector
}
