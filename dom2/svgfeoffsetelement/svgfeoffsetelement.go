package svgfeoffsetelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEOffsetElement struct
// js:"SVGFEOffsetElement,omit"
type SVGFEOffsetElement struct {
	window.SVGElement
}

// Height
func (*SVGFEOffsetElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEOffsetElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEOffsetElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEOffsetElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEOffsetElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// Dx
func (*SVGFEOffsetElement) Dx() (dx *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Dx")
	return dx
}

// Dy
func (*SVGFEOffsetElement) Dy() (dy *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Dy")
	return dy
}

// In1
func (*SVGFEOffsetElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}
