package svgfeturbulenceelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom/svganimatedinteger"
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatednumber"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFETurbulenceElement struct
// js:"SVGFETurbulenceElement,omit"
type SVGFETurbulenceElement struct {
	window.SVGElement
}

// Height prop
func (*SVGFETurbulenceElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// Result prop
func (*SVGFETurbulenceElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

// Width prop
func (*SVGFETurbulenceElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGFETurbulenceElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGFETurbulenceElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

// BaseFrequencyX prop
func (*SVGFETurbulenceElement) BaseFrequencyX() (baseFrequencyX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.baseFrequencyX")
	return baseFrequencyX
}

// BaseFrequencyY prop
func (*SVGFETurbulenceElement) BaseFrequencyY() (baseFrequencyY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.baseFrequencyY")
	return baseFrequencyY
}

// NumOctaves prop
func (*SVGFETurbulenceElement) NumOctaves() (numOctaves *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.numOctaves")
	return numOctaves
}

// Seed prop
func (*SVGFETurbulenceElement) Seed() (seed *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.seed")
	return seed
}

// StitchTiles prop
func (*SVGFETurbulenceElement) StitchTiles() (stitchTiles *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.stitchTiles")
	return stitchTiles
}

// Type prop
func (*SVGFETurbulenceElement) Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}
