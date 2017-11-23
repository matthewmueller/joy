package window

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
