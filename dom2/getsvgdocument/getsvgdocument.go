package getsvgdocument

import "github.com/matthewmueller/golly/dom2/window"

// js:"GetSVGDocument,omit"
type GetSVGDocument interface {

	// GetSVGDocument
	GetSVGDocument() (w window.Document)
}
