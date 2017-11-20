package getsvgdocument

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/js"
)

type GetSVGDocument struct {
}

func (*GetSVGDocument) GetSVGDocument() (d *document.Document) {
	js.Rewrite("$<.getSVGDocument()")
	return d
}
