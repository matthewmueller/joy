package svgfeturbulenceelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatedinteger"
	"github.com/matthewmueller/golly/internal/gendom/dom/svganimatednumber"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/svgfilterprimitivestandardattributes"
	"github.com/matthewmueller/golly/js"
)

type SVGFETurbulenceElement struct {
	*svgelement.SVGElement
	*svgfilterprimitivestandardattributes.SVGFilterPrimitiveStandardAttributes
}

func (*SVGFETurbulenceElement) GetBaseFrequencyX() (baseFrequencyX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.baseFrequencyX")
	return baseFrequencyX
}

func (*SVGFETurbulenceElement) GetBaseFrequencyY() (baseFrequencyY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.baseFrequencyY")
	return baseFrequencyY
}

func (*SVGFETurbulenceElement) GetNumOctaves() (numOctaves *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.numOctaves")
	return numOctaves
}

func (*SVGFETurbulenceElement) GetSeed() (seed *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.seed")
	return seed
}

func (*SVGFETurbulenceElement) GetStitchTiles() (stitchTiles *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.stitchTiles")
	return stitchTiles
}

func (*SVGFETurbulenceElement) GetType() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}
