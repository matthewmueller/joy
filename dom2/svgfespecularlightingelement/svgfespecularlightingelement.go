package svgfespecularlightingelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFESpecularLightingElement,omit"
type SVGFESpecularLightingElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
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
