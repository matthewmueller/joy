package svgfecomponenttransferelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEComponentTransferElement struct
// js:"SVGFEComponentTransferElement,omit"
type SVGFEComponentTransferElement struct {
	window.SVGElement
}

// Height
func (*SVGFEComponentTransferElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEComponentTransferElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEComponentTransferElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEComponentTransferElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEComponentTransferElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// In1
func (*SVGFEComponentTransferElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}
