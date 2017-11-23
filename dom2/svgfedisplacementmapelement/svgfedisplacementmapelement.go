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

// Height
func (*SVGFEDisplacementMapElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEDisplacementMapElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEDisplacementMapElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEDisplacementMapElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEDisplacementMapElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
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
