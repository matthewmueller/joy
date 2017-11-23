package window

import (
	"github.com/matthewmueller/golly/dom/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom/svgmatrix"
	"github.com/matthewmueller/golly/dom/svgrect"
)

// js:"SVGGraphicsElement,omit"
type SVGGraphicsElement interface {
	SVGElement

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
