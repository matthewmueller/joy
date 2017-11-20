package svggraphicselement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgmatrix"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgrect"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgtests"
	"github.com/matthewmueller/golly/js"
)

type SVGGraphicsElement struct {
	*svgelement.SVGElement
	*svgtests.SVGTests
}

func (*SVGGraphicsElement) GetBBox() (s *svgrect.SVGRect) {
	js.Rewrite("$<.getBBox()")
	return s
}

func (*SVGGraphicsElement) GetCTM() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.getCTM()")
	return s
}

func (*SVGGraphicsElement) GetScreenCTM() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.getScreenCTM()")
	return s
}

func (*SVGGraphicsElement) GetTransformToElement(element *svgelement.SVGElement) (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.getTransformToElement($1)", element)
	return s
}

func (*SVGGraphicsElement) GetFarthestViewportElement() (farthestViewportElement *svgelement.SVGElement) {
	js.Rewrite("$<.farthestViewportElement")
	return farthestViewportElement
}

func (*SVGGraphicsElement) GetNearestViewportElement() (nearestViewportElement *svgelement.SVGElement) {
	js.Rewrite("$<.nearestViewportElement")
	return nearestViewportElement
}

func (*SVGGraphicsElement) GetTransform() (transform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$<.transform")
	return transform
}
