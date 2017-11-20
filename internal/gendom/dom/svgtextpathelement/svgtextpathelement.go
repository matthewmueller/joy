package svgtextpathelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtextcontentelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgurireference"
	"github.com/matthewmueller/golly/js"
)

type SVGTextPathElement struct {
	*svgtextcontentelement.SVGTextContentElement
	*svgurireference.SVGURIReference
}

func (*SVGTextPathElement) GetMethod() (method *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.method")
	return method
}

func (*SVGTextPathElement) GetSpacing() (spacing *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.spacing")
	return spacing
}

func (*SVGTextPathElement) GetStartOffset() (startOffset *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.startOffset")
	return startOffset
}
