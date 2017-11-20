package svgtests

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgstringlist"
	"github.com/matthewmueller/golly/js"
)

type SVGTests struct {
}

func (*SVGTests) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.hasExtension($1)", extension)
	return b
}

func (*SVGTests) GetRequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredExtensions")
	return requiredExtensions
}

func (*SVGTests) GetRequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList) {
	js.Rewrite("$<.requiredFeatures")
	return requiredFeatures
}

func (*SVGTests) GetSystemLanguage() (systemLanguage *svgstringlist.SVGStringList) {
	js.Rewrite("$<.systemLanguage")
	return systemLanguage
}
