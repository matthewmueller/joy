package svgrectelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGRectElement struct
// js:"SVGRectElement,omit"
type SVGRectElement struct {
	window.SVGGraphicsElement
}

// Height prop
func (*SVGRectElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Rx prop
func (*SVGRectElement) Rx() (rx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.rx")
	return rx
}

// Ry prop
func (*SVGRectElement) Ry() (ry *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.ry")
	return ry
}

// Width prop
func (*SVGRectElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGRectElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGRectElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
