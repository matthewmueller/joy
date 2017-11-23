package svgtextpathelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/svgtextcontentelement"
	"github.com/matthewmueller/golly/js"
)

// SVGTextPathElement struct
// js:"SVGTextPathElement,omit"
type SVGTextPathElement struct {
	svgtextcontentelement.SVGTextContentElement
}

// Href
func (*SVGTextPathElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Href")
	return href
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
