package svgpatternelement

import (
	"github.com/matthewmueller/golly/dom2/svganimatedenumeration"
	"github.com/matthewmueller/golly/dom2/svganimatedlength"
	"github.com/matthewmueller/golly/dom2/svganimatedpreserveaspectratio"
	"github.com/matthewmueller/golly/dom2/svganimatedrect"
	"github.com/matthewmueller/golly/dom2/svganimatedstring"
	"github.com/matthewmueller/golly/dom2/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom2/svgstringlist"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SVGPatternElement struct
// js:"SVGPatternElement,omit"
type SVGPatternElement struct {
	window.SVGElement
}

// HasExtension
func (*SVGPatternElement) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.HasExtension($1)", extension)
	return b
}

// RequiredExtensions
func (*SVGPatternElement) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$<.RequiredExtensions")
	return requiredExtensions
}

// RequiredFeatures
func (*SVGPatternElement) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$<.RequiredFeatures")
	return requiredFeatures
}

// SystemLanguage
func (*SVGPatternElement) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$<.SystemLanguage")
	return systemLanguage
}

// PreserveAspectRatio
func (*SVGPatternElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.PreserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox
func (*SVGPatternElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.ViewBox")
	return viewBox
}

// Href
func (*SVGPatternElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.Href")
	return href
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
