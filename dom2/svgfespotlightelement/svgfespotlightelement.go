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

// LimitingConeAngle
func (*SVGFESpotLightElement) LimitingConeAngle() (limitingConeAngle *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.LimitingConeAngle")
	return limitingConeAngle
}

// PointsAtX
func (*SVGFESpotLightElement) PointsAtX() (pointsAtX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.PointsAtX")
	return pointsAtX
}

// PointsAtY
func (*SVGFESpotLightElement) PointsAtY() (pointsAtY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.PointsAtY")
	return pointsAtY
}

// PointsAtZ
func (*SVGFESpotLightElement) PointsAtZ() (pointsAtZ *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.PointsAtZ")
	return pointsAtZ
}

// SpecularExponent
func (*SVGFESpotLightElement) SpecularExponent() (specularExponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.SpecularExponent")
	return specularExponent
}

// X
func (*SVGFESpotLightElement) X() (x *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFESpotLightElement) Y() (y *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Y")
	return y
}

// Z
func (*SVGFESpotLightElement) Z() (z *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Z")
	return z
}
