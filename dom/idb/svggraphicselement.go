package idb

import (
	"github.com/matthewmueller/golly/dom2/svganimatedtransformlist"
	"github.com/matthewmueller/golly/dom2/svgmatrix"
	"github.com/matthewmueller/golly/dom2/svgrect"
)

// js:"SVGGraphicsElement,omit"
type SVGGraphicsElement interface {
	SVGElement

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
