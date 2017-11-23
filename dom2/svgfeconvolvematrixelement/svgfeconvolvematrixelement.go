package svgfeconvolvematrixelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedboolean"
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedinteger"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEConvolveMatrixElement,omit"
type SVGFEConvolveMatrixElement struct {
	window.SVGElement
	svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

// Bias
func (*SVGFEConvolveMatrixElement) Bias() (bias *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Bias")
	return bias
}

// Divisor
func (*SVGFEConvolveMatrixElement) Divisor() (divisor *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Divisor")
	return divisor
}

// EdgeMode
func (*SVGFEConvolveMatrixElement) EdgeMode() (edgeMode *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.EdgeMode")
	return edgeMode
}

// In1
func (*SVGFEConvolveMatrixElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}

// KernelMatrix
func (*SVGFEConvolveMatrixElement) KernelMatrix() (kernelMatrix *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.KernelMatrix")
	return kernelMatrix
}

// KernelUnitLengthX
func (*SVGFEConvolveMatrixElement) KernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.KernelUnitLengthX")
	return kernelUnitLengthX
}

// KernelUnitLengthY
func (*SVGFEConvolveMatrixElement) KernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.KernelUnitLengthY")
	return kernelUnitLengthY
}

// OrderX
func (*SVGFEConvolveMatrixElement) OrderX() (orderX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.OrderX")
	return orderX
}

// OrderY
func (*SVGFEConvolveMatrixElement) OrderY() (orderY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.OrderY")
	return orderY
}

// PreserveAlpha
func (*SVGFEConvolveMatrixElement) PreserveAlpha() (preserveAlpha *svganimatedboolean.SVGAnimatedBoolean) {
	js.Rewrite("$<.PreserveAlpha")
	return preserveAlpha
}

// TargetX
func (*SVGFEConvolveMatrixElement) TargetX() (targetX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.TargetX")
	return targetX
}

// TargetY
func (*SVGFEConvolveMatrixElement) TargetY() (targetY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.TargetY")
	return targetY
}
