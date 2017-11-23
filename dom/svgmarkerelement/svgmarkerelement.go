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

// SetOrientToAngle fn
func (*SVGMarkerElement) SetOrientToAngle(angle *svgangle.SVGAngle) {
	js.Rewrite("$<.setOrientToAngle($1)", angle)
}

// SetOrientToAuto fn
func (*SVGMarkerElement) SetOrientToAuto() {
	js.Rewrite("$<.setOrientToAuto()")
}

// PreserveAspectRatio prop
func (*SVGMarkerElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
func (*SVGMarkerElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}

// MarkerHeight prop
func (*SVGMarkerElement) MarkerHeight() (markerHeight *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.markerHeight")
	return markerHeight
}

// MarkerUnits prop
func (*SVGMarkerElement) MarkerUnits() (markerUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.markerUnits")
	return markerUnits
}

// MarkerWidth prop
func (*SVGMarkerElement) MarkerWidth() (markerWidth *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.markerWidth")
	return markerWidth
}

// OrientAngle prop
func (*SVGMarkerElement) OrientAngle() (orientAngle *svganimatedangle.SVGAnimatedAngle) {
	js.Rewrite("$<.orientAngle")
	return orientAngle
}

// OrientType prop
func (*SVGMarkerElement) OrientType() (orientType *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.orientType")
	return orientType
}

// RefX prop
func (*SVGMarkerElement) RefX() (refX *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.refX")
	return refX
}

// RefY prop
func (*SVGMarkerElement) RefY() (refY *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.refY")
	return refY
}
