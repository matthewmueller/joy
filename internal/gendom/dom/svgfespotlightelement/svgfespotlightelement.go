package svgfespotlightelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGFESpotLightElement struct {
	*svgelement.SVGElement
}

func (*SVGFESpotLightElement) GetLimitingConeAngle() (limitingConeAngle *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.limitingConeAngle")
	return limitingConeAngle
}

func (*SVGFESpotLightElement) GetPointsAtX() (pointsAtX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtX")
	return pointsAtX
}

func (*SVGFESpotLightElement) GetPointsAtY() (pointsAtY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtY")
	return pointsAtY
}

func (*SVGFESpotLightElement) GetPointsAtZ() (pointsAtZ *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtZ")
	return pointsAtZ
}

func (*SVGFESpotLightElement) GetSpecularExponent() (specularExponent *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.specularExponent")
	return specularExponent
}

func (*SVGFESpotLightElement) GetX() (x *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFESpotLightElement) GetY() (y *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGFESpotLightElement) GetZ() (z *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.z")
	return z
}
