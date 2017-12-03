package svgtests

import "github.com/matthewmueller/joy/dom/svgstringlist"

// SVGTests interface
// js:"SVGTests"
type SVGTests interface {

	// HasExtension
	// js:"hasExtension"
	// jsrewrite:"$_.hasExtension($1)"
	HasExtension(extension string) (b bool)

	// RequiredExtensions prop
	// js:"requiredExtensions"
	// jsrewrite:"$_.requiredExtensions"
	RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList)

	// RequiredFeatures prop
	// js:"requiredFeatures"
	// jsrewrite:"$_.requiredFeatures"
	RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList)

	// SystemLanguage prop
	// js:"systemLanguage"
	// jsrewrite:"$_.systemLanguage"
	SystemLanguage() (systemLanguage *svgstringlist.SVGStringList)
}
