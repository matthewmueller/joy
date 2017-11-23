package svgfemergenodeelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFEMergeNodeElement struct
// js:"SVGFEMergeNodeElement,omit"
type SVGFEMergeNodeElement struct {
	window.SVGElement
}

// In1 prop
func (*SVGFEMergeNodeElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}
