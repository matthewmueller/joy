package window

import "github.com/matthewmueller/golly/dom2/domstringmap"

// js:"SVGElement,omit"
type SVGElement interface {
	Element

	// Dataset
	Dataset() (dataset *domstringmap.DOMStringMap)

	// OwnerSVGElement
	OwnerSVGElement() (ownerSVGElement *SVGSVGElement)

	// ViewportElement
	ViewportElement() (viewportElement SVGElement)

	// Xmlbase
	Xmlbase() (xmlbase string)

	// Xmlbase
	SetXmlbase(xmlbase string)
}
