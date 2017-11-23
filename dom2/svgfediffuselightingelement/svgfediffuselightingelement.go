package svgfediffuselightingelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEDiffuseLightingElement struct
// js:"SVGFEDiffuseLightingElement,omit"
type SVGFEDiffuseLightingElement struct {
	window.SVGElement
}

// Height
func (*SVGFEDiffuseLightingElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEDiffuseLightingElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEDiffuseLightingElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEDiffuseLightingElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEDiffuseLightingElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
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
