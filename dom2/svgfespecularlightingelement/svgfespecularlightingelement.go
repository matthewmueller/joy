package svgfespecularlightingelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFESpecularLightingElement struct
// js:"SVGFESpecularLightingElement,omit"
type SVGFESpecularLightingElement struct {
	window.SVGElement
}

// Height
func (*SVGFESpecularLightingElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFESpecularLightingElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFESpecularLightingElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFESpecularLightingElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFESpecularLightingElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// In1
func (*SVGFESpecularLightingElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// KernelUnitLengthX
func (*SVGFESpecularLightingElement) KernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.KernelUnitLengthX")
	return kernelUnitLengthX
}

// KernelUnitLengthY
func (*SVGFESpecularLightingElement) KernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.KernelUnitLengthY")
	return kernelUnitLengthY
}

// SpecularConstant
func (*SVGFESpecularLightingElement) SpecularConstant() (specularConstant *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.SpecularConstant")
	return specularConstant
}

// SpecularExponent
func (*SVGFESpecularLightingElement) SpecularExponent() (specularExponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.SpecularExponent")
	return specularExponent
}

// SurfaceScale
func (*SVGFESpecularLightingElement) SurfaceScale() (surfaceScale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.SurfaceScale")
	return surfaceScale
}
