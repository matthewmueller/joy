package svgmarkerelement

import (
	"github.com/matthewmueller/golly/dom2/svgangle"
	"github.com/matthewmueller/golly/dom2/svganimatedangle"
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedrect"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGMarkerElement struct
// js:"SVGMarkerElement,omit"
type SVGMarkerElement struct {
	window.SVGElement
}

// SetOrientToAngle
func (*SVGMarkerElement) SetOrientToAngle(angle *svgangle.SVGAngle) {
	js.Rewrite("$<.SetOrientToAngle($1)", angle)
}

// SetOrientToAuto
func (*SVGMarkerElement) SetOrientToAuto() {
	js.Rewrite("$<.SetOrientToAuto()")
}

// PreserveAspectRatio
func (*SVGMarkerElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.PreserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox
func (*SVGMarkerElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.ViewBox")
	return viewBox
}

// MarkerHeight
func (*SVGMarkerElement) MarkerHeight() (markerHeight *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.MarkerHeight")
	return markerHeight
}

// MarkerUnits
func (*SVGMarkerElement) MarkerUnits() (markerUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.MarkerUnits")
	return markerUnits
}

// MarkerWidth
func (*SVGMarkerElement) MarkerWidth() (markerWidth *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.MarkerWidth")
	return markerWidth
}

// OrientAngle
func (*SVGMarkerElement) OrientAngle() (orientAngle *svganimatedangle.SVGAnimatedAngle) {
	js.Rewrite("$<.OrientAngle")
	return orientAngle
}

// OrientType
func (*SVGMarkerElement) OrientType() (orientType *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.OrientType")
	return orientType
}

// RefX
func (*SVGMarkerElement) RefX() (refX *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.RefX")
	return refX
}

// RefY
func (*SVGMarkerElement) RefY() (refY *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.RefY")
	return refY
}
