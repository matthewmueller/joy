package window

import "github.com/matthewmueller/golly/dom/domstringmap"

// SVGElement interface
// js:"SVGElement"
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

	// SetXmlbase prop
	// js:"xmlbase"
	SetXmlbase(xmlbase string)
}
