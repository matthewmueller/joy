package svgfediffuselightingelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEDiffuseLightingElement struct
// js:"SVGFEDiffuseLightingElement,omit"
type SVGFEDiffuseLightingElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFEDiffuseLightingElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEDiffuseLightingElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEDiffuseLightingElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEDiffuseLightingElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEDiffuseLightingElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// DiffuseConstant prop
func (*SVGFEDiffuseLightingElement) DiffuseConstant() (diffuseConstant *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.diffuseConstant")
	return diffuseConstant
}

// In1 prop
func (*SVGFEDiffuseLightingElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// KernelUnitLengthX prop
func (*SVGFEDiffuseLightingElement) KernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

// KernelUnitLengthY prop
func (*SVGFEDiffuseLightingElement) KernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

// SurfaceScale prop
func (*SVGFEDiffuseLightingElement) SurfaceScale() (surfaceScale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.surfaceScale")
	return surfaceScale
}
