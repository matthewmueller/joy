package window

import (
	"github.com/matthewmueller/golly/dom2/svgmatrix"
	"github.com/matthewmueller/golly/dom2/svgrect"
	"github.com/matthewmueller/golly/dom2/svgtests"
)

// js:"SVGGraphicsElement,omit"
type SVGGraphicsElement interface {
	SVGElement
	svgtests.SVGTests

	// GetBBox
	GetBBox() (s *svgrect.SVGRect)

	// GetCTM
	GetCTM() (s *svgmatrix.SVGMatrix)

	// GetScreenCTM
	GetScreenCTM() (s *svgmatrix.SVGMatrix)

	// GetTransformToElement
	GetTransformToElement(element SVGElement) (s *svgmatrix.SVGMatrix)

	// FarthestViewportElement
	FarthestViewportElement() (farthestViewportElement SVGElement)

	// NearestViewportElement
	NearestViewportElement() (nearestViewportElement SVGElement)

	// Transform
	Transform() (transform *svganimatedtransformlist.SVGAnimatedTransformList)
}
