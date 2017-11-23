package svgsymbolelement

import (
	"github.com/matthewmueller/golly/dom2/svgfittoviewbox"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"SVGSymbolElement,omit"
type SVGSymbolElement struct {
	window.SVGElement
	svgfittoviewbox.SVGFitToViewBox
}
