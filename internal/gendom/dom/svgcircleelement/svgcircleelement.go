package svgcircleelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svggraphicselement"
	"github.com/matthewmueller/golly/js"
)

type SVGCircleElement struct {
	*svggraphicselement.SVGGraphicsElement
}

func (*SVGCircleElement) GetCx() (cx *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

func (*SVGCircleElement) GetCy() (cy *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

func (*SVGCircleElement) GetR() (r *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.r")
	return r
}
