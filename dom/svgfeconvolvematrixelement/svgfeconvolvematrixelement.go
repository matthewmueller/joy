package svgfeconvolvematrixelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedboolean"
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedinteger"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatednumberlist"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEConvolveMatrixElement struct
// js:"SVGFEConvolveMatrixElement,omit"
type SVGFEConvolveMatrixElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFEConvolveMatrixElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEConvolveMatrixElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEConvolveMatrixElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEConvolveMatrixElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEConvolveMatrixElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// Bias prop
func (*SVGFEConvolveMatrixElement) Bias() (bias *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.bias")
	return bias
}

// Divisor prop
func (*SVGFEConvolveMatrixElement) Divisor() (divisor *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.divisor")
	return divisor
}

// EdgeMode prop
func (*SVGFEConvolveMatrixElement) EdgeMode() (edgeMode *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.edgeMode")
	return edgeMode
}

// In1 prop
func (*SVGFEConvolveMatrixElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

// KernelMatrix prop
func (*SVGFEConvolveMatrixElement) KernelMatrix() (kernelMatrix *svganimatednumberlist.SVGAnimatedNumberList) {
	js.Rewrite("$<.kernelMatrix")
	return kernelMatrix
}

// KernelUnitLengthX prop
func (*SVGFEConvolveMatrixElement) KernelUnitLengthX() (kernelUnitLengthX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

// KernelUnitLengthY prop
func (*SVGFEConvolveMatrixElement) KernelUnitLengthY() (kernelUnitLengthY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

// OrderX prop
func (*SVGFEConvolveMatrixElement) OrderX() (orderX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.orderX")
	return orderX
}

// OrderY prop
func (*SVGFEConvolveMatrixElement) OrderY() (orderY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.orderY")
	return orderY
}

// PreserveAlpha prop
func (*SVGFEConvolveMatrixElement) PreserveAlpha() (preserveAlpha *svganimatedboolean.SVGAnimatedBoolean) {
	js.Rewrite("$<.preserveAlpha")
	return preserveAlpha
}

// TargetX prop
func (*SVGFEConvolveMatrixElement) TargetX() (targetX *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.targetX")
	return targetX
}

// TargetY prop
func (*SVGFEConvolveMatrixElement) TargetY() (targetY *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.targetY")
	return targetY
}
