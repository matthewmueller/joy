package svgmaskelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svgstringlist"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGMaskElement struct
// js:"SVGMaskElement,omit"
type SVGMaskElement struct {
	window.SVGElement
}

// HasExtension
func (*SVGMaskElement) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.HasExtension($1)", extension)
	return b
}

// RequiredExtensions
func (*SVGMaskElement) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$<.RequiredExtensions")
	return requiredExtensions
}

// RequiredFeatures
func (*SVGMaskElement) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$<.RequiredFeatures")
	return requiredFeatures
}

// SystemLanguage
func (*SVGMaskElement) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$<.SystemLanguage")
	return systemLanguage
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
