package svglineelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGLineElement struct
// js:"SVGLineElement,omit"
type SVGLineElement struct {
	window.SVGGraphicsElement
}

// X1 prop
func (*SVGLineElement) X1() (x1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x1")
	return x1
}

// X2 prop
func (*SVGLineElement) X2() (x2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x2")
	return x2
}

// Y1 prop
func (*SVGLineElement) Y1() (y1 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y1")
	return y1
}

// Y2 prop
func (*SVGLineElement) Y2() (y2 *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y2")
	return y2
}
