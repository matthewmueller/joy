package svgforeignobjectelement

import "github.com/matthewmueller/golly/js"

// js:"SVGForeignObjectElement,omit"
type SVGForeignObjectElement struct {
	window.SVGGraphicsElement
}

// Height
func (*SVGForeignObjectElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Width
func (*SVGForeignObjectElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGForeignObjectElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGForeignObjectElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
