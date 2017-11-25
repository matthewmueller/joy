package window

import "github.com/matthewmueller/golly/dom/domstringmap"

// js:"SVGElement,omit"
type SVGElement interface {
	Element

	// Dataset prop
	// js:"dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// OwnerSVGElement prop
	// js:"ownerSVGElement"
	OwnerSVGElement() (ownerSVGElement *SVGSVGElement)

	// ViewportElement prop
	// js:"viewportElement"
	ViewportElement() (viewportElement SVGElement)

	// Xmlbase prop
	// js:"xmlbase"
	Xmlbase() (xmlbase string)

	// Xmlbase prop
	// js:"setxmlbase"
	SetXmlbase(xmlbase string)
}
