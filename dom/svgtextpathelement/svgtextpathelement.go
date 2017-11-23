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

// Href prop
func (*SVGTextPathElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

// Method prop
func (*SVGTextPathElement) Method() (method *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.method")
	return method
}

// Spacing prop
func (*SVGTextPathElement) Spacing() (spacing *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.spacing")
	return spacing
}

// StartOffset prop
func (*SVGTextPathElement) StartOffset() (startOffset *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.startOffset")
	return startOffset
}
