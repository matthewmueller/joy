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

// HasExtension fn
func (*SVGMaskElement) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.hasExtension($1)", extension)
	return b
}

// RequiredExtensions prop
func (*SVGMaskElement) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredExtensions")
	return requiredExtensions
}

// RequiredFeatures prop
func (*SVGMaskElement) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredFeatures")
	return requiredFeatures
}

// SystemLanguage prop
func (*SVGMaskElement) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$<.systemLanguage")
	return systemLanguage
}

// Height prop
func (*SVGMaskElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// MaskContentUnits prop
func (*SVGMaskElement) MaskContentUnits() (maskContentUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.maskContentUnits")
	return maskContentUnits
}

// MaskUnits prop
func (*SVGMaskElement) MaskUnits() (maskUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.maskUnits")
	return maskUnits
}

// Width prop
func (*SVGMaskElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGMaskElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGMaskElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
