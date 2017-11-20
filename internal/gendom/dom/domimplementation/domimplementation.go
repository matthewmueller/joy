package domimplementation

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/documenttype"
	"github.com/matthewmueller/golly/js"
)

type DOMImplementation struct {
}

func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *documenttype.DocumentType) (d *document.Document) {
	js.Rewrite("$<.createDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) (d *documenttype.DocumentType) {
	js.Rewrite("$<.createDocumentType($1, $2, $3)", qualifiedName, publicId, systemId)
	return d
}

func (*DOMImplementation) CreateHTMLDocument(title string) (d *document.Document) {
	js.Rewrite("$<.createHTMLDocument($1)", title)
	return d
}

func (*DOMImplementation) HasFeature() (b bool) {
	js.Rewrite("$<.hasFeature()")
	return b
}
