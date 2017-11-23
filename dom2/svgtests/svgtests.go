package svgtests

import (
	"github.com/matthewmueller/golly/dom2/svgstringlist"
	"github.com/matthewmueller/golly/js"
)

// SVGTests struct
// js:"SVGTests,omit"
type SVGTests struct {
}

// HasExtension
func (*SVGTests) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.HasExtension($1)", extension)
	return b
}

// RequiredExtensions
func (*SVGTests) RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$<.RequiredExtensions")
	return requiredExtensions
}

// RequiredFeatures
func (*SVGTests) RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$<.RequiredFeatures")
	return requiredFeatures
}

// SystemLanguage
func (*SVGTests) SystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$<.SystemLanguage")
	return systemLanguage
}
