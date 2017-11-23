package svgtests

import (
	"github.com/matthewmueller/golly/dom2/svgstringlist"
	"github.com/matthewmueller/golly/js"
)

// SVGTests struct
// js:"SVGTests,omit"
type SVGTests struct {
}

// HasExtension fn
func (*SVGTests) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.hasExtension($1)", extension)
	return b
}

// RequiredExtensions prop
func (*SVGTests) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredExtensions")
	return requiredExtensions
}

// RequiredFeatures prop
func (*SVGTests) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredFeatures")
	return requiredFeatures
}

// SystemLanguage prop
func (*SVGTests) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$<.systemLanguage")
	return systemLanguage
}
