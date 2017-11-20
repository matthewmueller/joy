package svgfeconvolvematrixelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedboolean"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedinteger"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumberlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedstring"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFEConvolveMatrixElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEConvolveMatrixElement) GetBias() (bias *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.bias")
	return bias
}

func (*SVGFEConvolveMatrixElement) GetDivisor() (divisor *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.divisor")
	return divisor
}

func (*SVGFEConvolveMatrixElement) GetEdgeMode() (edgeMode *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.edgeMode")
	return edgeMode
}

func (*SVGFEConvolveMatrixElement) GetIn1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEConvolveMatrixElement) GetKernelMatrix() (kernelMatrix *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.kernelMatrix")
	return kernelMatrix
}

func (*SVGFEConvolveMatrixElement) GetKernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

func (*SVGFEConvolveMatrixElement) GetKernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

func (*SVGFEConvolveMatrixElement) GetOrderX() (orderX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.orderX")
	return orderX
}

func (*SVGFEConvolveMatrixElement) GetOrderY() (orderY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.orderY")
	return orderY
}

func (*SVGFEConvolveMatrixElement) GetPreserveAlpha() (preserveAlpha *svganimatedboolean.SVGAnimatedBoolean) {
	js.Rewrite("$<.preserveAlpha")
	return preserveAlpha
}

func (*SVGFEConvolveMatrixElement) GetTargetX() (targetX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.targetX")
	return targetX
}

func (*SVGFEConvolveMatrixElement) GetTargetY() (targetY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.targetY")
	return targetY
}
