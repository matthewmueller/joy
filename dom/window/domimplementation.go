package window

import "github.com/matthewmueller/golly/js"

// DOMImplementation struct
// js:"DOMImplementation,omit"
type DOMImplementation struct {
}

// CreateDocument fn
func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *DocumentType) (d Document) {
	js.Rewrite("$<.createDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

// CreateDocumentType fn
func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) (d *DocumentType) {
	js.Rewrite("$<.createDocumentType($1, $2, $3)", qualifiedName, publicId, systemId)
	return d
}

// CreateHTMLDocument fn
func (*DOMImplementation) CreateHTMLDocument(title string) (d Document) {
	js.Rewrite("$<.createHTMLDocument($1)", title)
	return d
}

// HasFeature fn
func (*DOMImplementation) HasFeature() (b bool) {
	js.Rewrite("$<.hasFeature()")
	return b
}
