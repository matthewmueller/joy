package svgtests

import "github.com/matthewmueller/golly/dom/svgstringlist"

// SVGTests interface
// js:"SVGTests"
type SVGTests interface {

	// HasExtension
	// js:"hasExtension"
	HasExtension(extension string) (b bool)

	// RequiredExtensions prop
	// js:"requiredExtensions"
	RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList)

	// RequiredFeatures prop
	// js:"requiredFeatures"
	RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList)

	// SystemLanguage prop
	// js:"systemLanguage"
	SystemLanguage() (systemLanguage *svgstringlist.SVGStringList)
}
