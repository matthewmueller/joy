package svgfedistantlightelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/js"
)

type SVGFEDistantLightElement struct {
	*svgelement.SVGElement
}

func (*SVGFEDistantLightElement) GetAzimuth() (azimuth *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.azimuth")
	return azimuth
}

func (*SVGFEDistantLightElement) GetElevation() (elevation *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.elevation")
	return elevation
}
