package svgsymbolelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfittoviewbox"
)

type SVGSymbolElement struct {
	*svgelement.SVGElement
	*svgfittoviewbox.SVGFitToViewBox
}
