package svgfespecularlightingelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFESpecularLightingElement struct
// js:"SVGFESpecularLightingElement,omit"
type SVGFESpecularLightingElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFESpecularLightingElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFESpecularLightingElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFESpecularLightingElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFESpecularLightingElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFESpecularLightingElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// In1 prop
func (*SVGFESpecularLightingElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// KernelUnitLengthX prop
func (*SVGFESpecularLightingElement) KernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

// KernelUnitLengthY prop
func (*SVGFESpecularLightingElement) KernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

// SpecularConstant prop
func (*SVGFESpecularLightingElement) SpecularConstant() (specularConstant *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.specularConstant")
	return specularConstant
}

// SpecularExponent prop
func (*SVGFESpecularLightingElement) SpecularExponent() (specularExponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.specularExponent")
	return specularExponent
}

// SurfaceScale prop
func (*SVGFESpecularLightingElement) SurfaceScale() (surfaceScale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.surfaceScale")
	return surfaceScale
}
