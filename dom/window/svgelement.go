package window

import (
	"github.com/matthewmueller/golly/dom/domstringmap"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGElement,omit"
type SVGElement interface {
	Element

	// Dataset prop
	// js:"dataset,rewrite=dataset"
	Dataset() (dataset *domstringmap.DOMStringMap)

	// OwnerSVGElement prop
	// js:"ownerSVGElement,rewrite=ownersvgelement"
	OwnerSVGElement() (ownerSVGElement *SVGSVGElement)

	// ViewportElement prop
	// js:"viewportElement,rewrite=viewportelement"
	ViewportElement() (viewportElement SVGElement)

	// Xmlbase prop
	// js:"xmlbase,rewrite=xmlbase"
	Xmlbase() (xmlbase string)

	// Xmlbase prop
	// js:"setxmlbase,rewrite=setxmlbase"
	SetXmlbase(xmlbase string)
}

// dataset prop
func dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$<.dataset")
	return dataset
}

// ownersvgelement prop
func ownersvgelement() (ownerSVGElement *SVGSVGElement) {
	js.Rewrite("$<.ownerSVGElement")
	return ownerSVGElement
}

// viewportelement prop
func viewportelement() (viewportElement SVGElement) {
	js.Rewrite("$<.viewportElement")
	return viewportElement
}

// xmlbase prop
func xmlbase() (xmlbase string) {
	js.Rewrite("$<.xmlbase")
	return xmlbase
}

// setxmlbase prop
func setxmlbase(xmlbase string) {
	js.Rewrite("$<.xmlbase = xmlbase")
}
