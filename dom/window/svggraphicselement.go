package window

import (
	"github.com/matthewmueller/golly/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom/svgmatrix"
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/dom/svgstringlist"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGGraphicsElement,omit"
type SVGGraphicsElement interface {
	SVGElement

	// HasExtension
	// js:"hasExtension,rewrite=hasextension"
	HasExtension(extension string) (b bool)

	// GetBBox
	// js:"getBBox,rewrite=getbbox"
	GetBBox() (s *svgrect.SVGRect)

	// GetCTM
	// js:"getCTM,rewrite=getctm"
	GetCTM() (s *svgmatrix.SVGMatrix)

	// GetScreenCTM
	// js:"getScreenCTM,rewrite=getscreenctm"
	GetScreenCTM() (s *svgmatrix.SVGMatrix)

	// GetTransformToElement
	// js:"getTransformToElement,rewrite=gettransformtoelement"
	GetTransformToElement(element SVGElement) (s *svgmatrix.SVGMatrix)

	// RequiredExtensions prop
	// js:"requiredExtensions,rewrite=requiredextensions"
	RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList)

	// RequiredFeatures prop
	// js:"requiredFeatures,rewrite=requiredfeatures"
	RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList)

	// SystemLanguage prop
	// js:"systemLanguage,rewrite=systemlanguage"
	SystemLanguage() (systemLanguage *svgstringlist.SVGStringList)

	// FarthestViewportElement prop
	// js:"farthestViewportElement,rewrite=farthestviewportelement"
	FarthestViewportElement() (farthestViewportElement SVGElement)

	// NearestViewportElement prop
	// js:"nearestViewportElement,rewrite=nearestviewportelement"
	NearestViewportElement() (nearestViewportElement SVGElement)

	// Transform prop
	// js:"transform,rewrite=transform"
	Transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList)
}

// getbbox fn
func getbbox() (s *svgrect.SVGRect) {
	js.Rewrite("$<.getBBox()")
	return s
}

// getctm fn
func getctm() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.getCTM()")
	return s
}

// getscreenctm fn
func getscreenctm() (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.getScreenCTM()")
	return s
}

// gettransformtoelement fn
func gettransformtoelement(element SVGElement) (s *svgmatrix.SVGMatrix) {
	js.Rewrite("$<.getTransformToElement($1)", element)
	return s
}

// farthestviewportelement prop
func farthestviewportelement() (farthestViewportElement SVGElement) {
	js.Rewrite("$<.farthestViewportElement")
	return farthestViewportElement
}

// nearestviewportelement prop
func nearestviewportelement() (nearestViewportElement SVGElement) {
	js.Rewrite("$<.nearestViewportElement")
	return nearestViewportElement
}

// transform prop
func transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList) {
	js.Rewrite("$<.transform")
	return transform
}
