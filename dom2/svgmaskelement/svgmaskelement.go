package svgmaskelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGMaskElement,omit"
type SVGMaskElement struct {
	window.SVGElement
	svgtests.SVGTests
	svgunittypes.SVGUnitTypes
}

// Height
func (*SVGMaskElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// MaskContentUnits
func (*SVGMaskElement) MaskContentUnits() (maskContentUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.MaskContentUnits")
	return maskContentUnits
}

// MaskUnits
func (*SVGMaskElement) MaskUnits() (maskUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.MaskUnits")
	return maskUnits
}

// Width
func (*SVGMaskElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGMaskElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGMaskElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
