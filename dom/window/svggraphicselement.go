package window

import (
	"github.com/matthewmueller/golly/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom/svgmatrix"
	"github.com/matthewmueller/golly/dom/svgrect"
	"github.com/matthewmueller/golly/dom/svgstringlist"
)

// js:"SVGGraphicsElement,omit"
type SVGGraphicsElement interface {
	SVGElement

	// HasExtension
	// js:"hasExtension"
	HasExtension(extension string) (b bool)

	// GetBBox
	// js:"getBBox"
	GetBBox() (s *svgrect.SVGRect)

	// GetCTM
	// js:"getCTM"
	GetCTM() (s *svgmatrix.SVGMatrix)

	// GetScreenCTM
	// js:"getScreenCTM"
	GetScreenCTM() (s *svgmatrix.SVGMatrix)

	// GetTransformToElement
	// js:"getTransformToElement"
	GetTransformToElement(element SVGElement) (s *svgmatrix.SVGMatrix)

	// RequiredExtensions prop
	// js:"requiredExtensions"
	RequiredExtensions() (requiredExtensions *svgstringlist.SVGStringList)

	// RequiredFeatures prop
	// js:"requiredFeatures"
	RequiredFeatures() (requiredFeatures *svgstringlist.SVGStringList)

	// SystemLanguage prop
	// js:"systemLanguage"
	SystemLanguage() (systemLanguage *svgstringlist.SVGStringList)

	// FarthestViewportElement prop
	// js:"farthestViewportElement"
	FarthestViewportElement() (farthestViewportElement SVGElement)

	// NearestViewportElement prop
	// js:"nearestViewportElement"
	NearestViewportElement() (nearestViewportElement SVGElement)

	// Transform prop
	// js:"transform"
	Transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList)
}
