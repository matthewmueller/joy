package svgpatternelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom2/svgfittoviewbox"
	"github.com/matthewmueller/golly/dom2/svgtests"
	"github.com/matthewmueller/golly/dom2/svgunittypes"
	"github.com/matthewmueller/golly/dom2/svgurireference"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPatternElement,omit"
type SVGPatternElement struct {
	window.SVGElement
	svgtests.SVGTests
	svgunittypes.SVGUnitTypes
	svgfittoviewbox.SVGFitToViewBox
	svgurireference.SVGURIReference
}

// Height
func (*SVGPatternElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// PatternContentUnits
func (*SVGPatternElement) PatternContentUnits() (patternContentUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.PatternContentUnits")
	return patternContentUnits
}

// PatternTransform
func (*SVGPatternElement) PatternTransform() (patternTransform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$<.PatternTransform")
	return patternTransform
}

// PatternUnits
func (*SVGPatternElement) PatternUnits() (patternUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.PatternUnits")
	return patternUnits
}

// Width
func (*SVGPatternElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGPatternElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGPatternElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
