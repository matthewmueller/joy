package getsvgdocument

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// GetSVGDocument struct
// js:"GetSVGDocument,omit"
type GetSVGDocument struct {
}

// GetSVGDocument fn
func (*GetSVGDocument) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.getSVGDocument()")
	return w
}
