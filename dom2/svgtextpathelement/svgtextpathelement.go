package svgtextpathelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svgurireference"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGTextPathElement,omit"
type SVGTextPathElement struct {
	svgtextcontentelement.SVGTextContentElement
	svgurireference.SVGURIReference
}

// Method
func (*SVGTextPathElement) Method() (method *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Method")
	return method
}

// Spacing
func (*SVGTextPathElement) Spacing() (spacing *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Spacing")
	return spacing
}

// StartOffset
func (*SVGTextPathElement) StartOffset() (startOffset *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.StartOffset")
	return startOffset
}
