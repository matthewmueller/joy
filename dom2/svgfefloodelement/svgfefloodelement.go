package svgfefloodelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEFloodElement struct
// js:"SVGFEFloodElement,omit"
type SVGFEFloodElement struct {
	window.SVGElement
}

// Height
func (*SVGFEFloodElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEFloodElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEFloodElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEFloodElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEFloodElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
