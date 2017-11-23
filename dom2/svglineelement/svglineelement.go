package svglineelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGLineElement,omit"
type SVGLineElement struct {
	window.SVGGraphicsElement
}

// X1
func (*SVGLineElement) X1() (x1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X1")
	return x1
}

// X2
func (*SVGLineElement) X2() (x2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X2")
	return x2
}

// Y1
func (*SVGLineElement) Y1() (y1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y1")
	return y1
}

// Y2
func (*SVGLineElement) Y2() (y2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y2")
	return y2
}
