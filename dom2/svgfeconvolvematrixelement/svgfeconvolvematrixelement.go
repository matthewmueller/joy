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

// Height
func (*SVGFEConvolveMatrixElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEConvolveMatrixElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEConvolveMatrixElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEConvolveMatrixElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEConvolveMatrixElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
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
