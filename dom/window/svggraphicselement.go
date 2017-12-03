package window

import (
	"github.com/matthewmueller/joy/dom/svganimatedtransformlist"
	"github.com/matthewmueller/joy/dom/svgmatrix"
	"github.com/matthewmueller/joy/dom/svgrect"
	"github.com/matthewmueller/joy/dom/svgstringlist"
)

// SVGGraphicsElement interface
// js:"SVGGraphicsElement"
type SVGGraphicsElement interface {
	SVGElement

	// HasExtension
	// js:"hasExtension"
	// jsrewrite:"$_.hasExtension($1)"
	HasExtension(extension string) (b bool)

	// GetBBox
	// js:"getBBox"
	// jsrewrite:"$_.getBBox()"
	GetBBox() (s *svgrect.SVGRect)

	// GetCTM
	// js:"getCTM"
	// jsrewrite:"$_.getCTM()"
	GetCTM() (s *svgmatrix.SVGMatrix)

	// GetScreenCTM
	// js:"getScreenCTM"
	// jsrewrite:"$_.getScreenCTM()"
	GetScreenCTM() (s *svgmatrix.SVGMatrix)

	// GetTransformToElement
	// js:"getTransformToElement"
	// jsrewrite:"$_.getTransformToElement($1)"
	GetTransformToElement(element SVGElement) (s *svgmatrix.SVGMatrix)

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

	// FarthestViewportElement prop
	// js:"farthestViewportElement"
	// jsrewrite:"$_.farthestViewportElement"
	FarthestViewportElement() (farthestViewportElement SVGElement)

	// NearestViewportElement prop
	// js:"nearestViewportElement"
	// jsrewrite:"$_.nearestViewportElement"
	NearestViewportElement() (nearestViewportElement SVGElement)

	// Transform prop
	// js:"transform"
	// jsrewrite:"$_.transform"
	Transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList)
}
