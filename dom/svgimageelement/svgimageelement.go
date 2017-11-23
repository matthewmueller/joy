package svgimageelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGImageElement struct
// js:"SVGImageElement,omit"
type SVGImageElement struct {
	window.SVGGraphicsElement
}

// Href prop
func (*SVGImageElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

// Height prop
func (*SVGImageElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// PreserveAspectRatio prop
func (*SVGImageElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

// Width prop
func (*SVGImageElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGImageElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGImageElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
