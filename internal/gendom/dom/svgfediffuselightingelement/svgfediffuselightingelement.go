package svgfediffuselightingelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFEDiffuseLightingElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEDiffuseLightingElement) GetDiffuseConstant() (diffuseConstant *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.diffuseConstant")
	return diffuseConstant
}

func (*SVGFEDiffuseLightingElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEDiffuseLightingElement) GetKernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

func (*SVGFEDiffuseLightingElement) GetKernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

func (*SVGFEDiffuseLightingElement) GetSurfaceScale() (surfaceScale *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.surfaceScale")
	return surfaceScale
}
