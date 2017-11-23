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

// HasExtension fn
func (*SVGPatternElement) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.hasExtension($1)", extension)
	return b
}

// RequiredExtensions prop
func (*SVGPatternElement) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredExtensions")
	return requiredExtensions
}

// RequiredFeatures prop
func (*SVGPatternElement) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredFeatures")
	return requiredFeatures
}

// SystemLanguage prop
func (*SVGPatternElement) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$<.systemLanguage")
	return systemLanguage
}

// PreserveAspectRatio prop
func (*SVGPatternElement) PreserveAspectRatio() (preserveAspectRatio *svganimatedpreserveaspectratio.SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

// ViewBox prop
func (*SVGPatternElement) ViewBox() (viewBox *svganimatedrect.SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}

// Href prop
func (*SVGPatternElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

// Height prop
func (*SVGPatternElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// PatternContentUnits prop
func (*SVGPatternElement) PatternContentUnits() (patternContentUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.patternContentUnits")
	return patternContentUnits
}

// PatternTransform prop
func (*SVGPatternElement) PatternTransform() (patternTransform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$<.patternTransform")
	return patternTransform
}

// PatternUnits prop
func (*SVGPatternElement) PatternUnits() (patternUnits *svganimatedenumeration.SVGAnimatedEnumeration) {
	js.Rewrite("$<.patternUnits")
	return patternUnits
}

// Width prop
func (*SVGPatternElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGPatternElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGPatternElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
