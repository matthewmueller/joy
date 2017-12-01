package getsvgdocument

import "github.com/matthewmueller/golly/dom/window"

// GetSVGDocument interface
// js:"GetSVGDocument"
type GetSVGDocument interface {

	// GetSVGDocument
	// js:"getSVGDocument"
	// jsrewrite:"$_.getSVGDocument()"
	GetSVGDocument() (w window.Document)
}
