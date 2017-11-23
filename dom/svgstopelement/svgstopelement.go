package svgstopelement

import (
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGStopElement struct
// js:"SVGStopElement,omit"
type SVGStopElement struct {
	window.SVGElement
}

// Offset prop
func (*SVGStopElement) Offset() (offset *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.offset")
	return offset
}
