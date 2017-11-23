package svgrectelement

import "github.com/matthewmueller/golly/js"

// js:"SVGRectElement,omit"
type SVGRectElement struct {
	window.SVGGraphicsElement
}

// Height
func (*SVGRectElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Rx
func (*SVGRectElement) Rx() (rx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Rx")
	return rx
}

// Ry
func (*SVGRectElement) Ry() (ry *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Ry")
	return ry
}

// Width
func (*SVGRectElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGRectElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGRectElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
