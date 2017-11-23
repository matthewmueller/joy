package svgfemergenodeelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGFEMergeNodeElement,omit"
type SVGFEMergeNodeElement struct {
	window.SVGElement
}

// In1
func (*SVGFEMergeNodeElement) In1() (in1 *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.In1")
	return in1
}
