package domparser

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/js"
)

type DOMParser struct {
}

func (*DOMParser) ParseFromString(source string, mimeType string) (d *document.Document) {
	js.Rewrite("$<.parseFromString($1, $2)", source, mimeType)
	return d
}
