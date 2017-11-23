package svgimageelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svgurireference"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGImageElement,omit"
type SVGImageElement struct {
	window.SVGGraphicsElement
	svgurireference.SVGURIReference
}

// Height
func (*SVGImageElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// PreserveAspectRatio
func (*SVGImageElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.PreserveAspectRatio")
	return preserveAspectRatio
}

// Width
func (*SVGImageElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGImageElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGImageElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
