package svgfeturbulenceelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedinteger"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatednumber"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGFETurbulenceElement struct
// js:"SVGFETurbulenceElement,omit"
type SVGFETurbulenceElement struct {
	window.SVGElement
}

// Height
func (*SVGFETurbulenceElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// Result
func (*SVGFETurbulenceElement) Result() (result *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Result")
	return result
}

// Width
func (*SVGFETurbulenceElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGFETurbulenceElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGFETurbulenceElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}

// BaseFrequencyX
func (*SVGFETurbulenceElement) BaseFrequencyX() (baseFrequencyX *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.BaseFrequencyX")
	return baseFrequencyX
}

// BaseFrequencyY
func (*SVGFETurbulenceElement) BaseFrequencyY() (baseFrequencyY *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.BaseFrequencyY")
	return baseFrequencyY
}

// NumOctaves
func (*SVGFETurbulenceElement) NumOctaves() (numOctaves *svganimatedinteger.SVGAnimatedInteger) {
	js.Rewrite("$<.NumOctaves")
	return numOctaves
}

// Seed
func (*SVGFETurbulenceElement) Seed() (seed *svganimatednumber.SVGAnimatedNumber) {
	js.Rewrite("$<.Seed")
	return seed
}

// StitchTiles
func (*SVGFETurbulenceElement) StitchTiles() (stitchTiles *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.StitchTiles")
	return stitchTiles
}

// Type
func (*SVGFETurbulenceElement) Type() (kind *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.Type")
	return kind
}
