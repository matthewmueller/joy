package svgfeoffsetelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEOffsetElement struct
// js:"SVGFEOffsetElement,omit"
type SVGFEOffsetElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFEOffsetElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFEOffsetElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFEOffsetElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFEOffsetElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFEOffsetElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// Dx prop
func (*SVGFEOffsetElement) Dx() (dx *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.dx")
	return dx
}

// Dy prop
func (*SVGFEOffsetElement) Dy() (dy *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.dy")
	return dy
}

// In1 prop
func (*SVGFEOffsetElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}
