package svgfediffuselightingelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEDiffuseLightingElement,omit"
type SVGFEDiffuseLightingElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// DiffuseConstant
func (*SVGFEDiffuseLightingElement) DiffuseConstant() (diffuseConstant *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.DiffuseConstant")
	return diffuseConstant
}

// In1
func (*SVGFEDiffuseLightingElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// KernelUnitLengthX
func (*SVGFEDiffuseLightingElement) KernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.KernelUnitLengthX")
	return kernelUnitLengthX
}

// KernelUnitLengthY
func (*SVGFEDiffuseLightingElement) KernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.KernelUnitLengthY")
	return kernelUnitLengthY
}

// SurfaceScale
func (*SVGFEDiffuseLightingElement) SurfaceScale() (surfaceScale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.SurfaceScale")
	return surfaceScale
}
