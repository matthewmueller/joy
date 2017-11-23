package svgfedisplacementmapelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEDisplacementMapElement struct
// js:"SVGFEDisplacementMapElement,omit"
type SVGFEDisplacementMapElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFEDisplacementMapElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEDisplacementMapElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEDisplacementMapElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEDisplacementMapElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEDisplacementMapElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// In1 prop
func (*SVGFEDisplacementMapElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// In2 prop
func (*SVGFEDisplacementMapElement) In2() (in2 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

// Scale prop
func (*SVGFEDisplacementMapElement) Scale() (scale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.scale")
	return scale
}

// XChannelSelector prop
func (*SVGFEDisplacementMapElement) XChannelSelector() (xChannelSelector *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.xChannelSelector")
	return xChannelSelector
}

// YChannelSelector prop
func (*SVGFEDisplacementMapElement) YChannelSelector() (yChannelSelector *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.yChannelSelector")
	return yChannelSelector
}
