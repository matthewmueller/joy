package svgfeimageelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEImageElement struct
// js:"SVGFEImageElement,omit"
type SVGFEImageElement struct {
	window.SVGElement
}

// Height
func (*SVGFEImageElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFEImageElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFEImageElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFEImageElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFEImageElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// Href
func (*SVGFEImageElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Href")
	return href
}

// PreserveAspectRatio
func (*SVGFEImageElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.PreserveAspectRatio")
	return preserveAspectRatio
}
