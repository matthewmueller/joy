package svgtests

// js:"SVGTests,omit"
type SVGTests interface {

	// HasExtension
	HasExtension(extension string) (b bool)

	// RequiredExtensions
	RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList)

	// RequiredFeatures
	RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList)

	// SystemLanguage
	SystemLanguage() (systemLanguage *svgstringlist.SVGStringList)
}
