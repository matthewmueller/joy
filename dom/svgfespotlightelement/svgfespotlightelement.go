package svgfespotlightelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFESpotLightElement struct
// js:"SVGFESpotLightElement,omit"
type SVGFESpotLightElement struct {
	window.SVGElement
}

// LimitingConeAngle prop
func (*SVGFESpotLightElement) LimitingConeAngle() (limitingConeAngle *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.limitingConeAngle")
	return limitingConeAngle
}

// PointsAtX prop
func (*SVGFESpotLightElement) PointsAtX() (pointsAtX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtX")
	return pointsAtX
}

// PointsAtY prop
func (*SVGFESpotLightElement) PointsAtY() (pointsAtY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtY")
	return pointsAtY
}

// PointsAtZ prop
func (*SVGFESpotLightElement) PointsAtZ() (pointsAtZ *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtZ")
	return pointsAtZ
}

// SpecularExponent prop
func (*SVGFESpotLightElement) SpecularExponent() (specularExponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.specularExponent")
	return specularExponent
}

// X prop
func (*SVGFESpotLightElement) X() (x *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFESpotLightElement) Y() (y *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.y")
	return y
}

// Z prop
func (*SVGFESpotLightElement) Z() (z *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.z")
	return z
}
