package svgstopelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGStopElement,omit"
type SVGStopElement struct {
	window.SVGElement
}

// Offset
func (*SVGStopElement) Offset() (offset *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Offset")
	return offset
}
