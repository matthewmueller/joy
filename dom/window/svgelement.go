package window

import "github.com/matthewmueller/golly/dom/domstringmap"

// SVGElement interface
// js:"SVGElement"
type SVGElement interface {
	Element

	// Dataset prop
	// js:"dataset"
	// jsrewrite:"$_.dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// OwnerSVGElement prop
	// js:"ownerSVGElement"
	// jsrewrite:"$_.ownerSVGElement"
	OwnerSVGElement() (ownerSVGElement *SVGSVGElement)

	// ViewportElement prop
	// js:"viewportElement"
	// jsrewrite:"$_.viewportElement"
	ViewportElement() (viewportElement SVGElement)

	// Xmlbase prop
	// js:"xmlbase"
	// jsrewrite:"$_.xmlbase"
	Xmlbase() (xmlbase string)

	// SetXmlbase prop
	// js:"xmlbase"
	// jsrewrite:"$_.xmlbase = $1"
	SetXmlbase(xmlbase string)
}
