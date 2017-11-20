package svgmarkerelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgangle"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedangle"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedlength"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfittoviewbox"
	"github.com/matthewmueller/golly/js"
)

type SVGMarkerElement struct {
	*svgelement.SVGElement
	*svgfittoviewbox.SVGFitToViewBox
}

func (*SVGMarkerElement) SetOrientToAngle(angle *svgangle.SVGAngle) {
	js.Rewrite("$<.setOrientToAngle($1)", angle)
}

func (*SVGMarkerElement) SetOrientToAuto() {
	js.Rewrite("$<.setOrientToAuto()")
}

func (*SVGMarkerElement) GetMarkerHeight() (markerHeight *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.markerHeight")
	return markerHeight
}

func (*SVGMarkerElement) GetMarkerUnits() (markerUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.markerUnits")
	return markerUnits
}

func (*SVGMarkerElement) GetMarkerWidth() (markerWidth *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.markerWidth")
	return markerWidth
}

func (*SVGMarkerElement) GetOrientAngle() (orientAngle *svganimatedangle.SVGAnimatedAngle) {
	js.Rewrite("$<.orientAngle")
	return orientAngle
}

func (*SVGMarkerElement) GetOrientType() (orientType *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.orientType")
	return orientType
}

func (*SVGMarkerElement) GetRefX() (refX *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.refX")
	return refX
}

func (*SVGMarkerElement) GetRefY() (refY *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.refY")
	return refY
}
