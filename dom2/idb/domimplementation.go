package idb

import "github.com/matthewmueller/golly/js"

// DOMImplementation struct
// js:"DOMImplementation,omit"
type DOMImplementation struct {
}

// CreateDocument
func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *DocumentType) (d Document) {
	js.Rewrite("$<.CreateDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

// CreateDocumentType
func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) (d *DocumentType) {
	js.Rewrite("$<.CreateDocumentType($1, $2, $3)", qualifiedName, publicId, systemId)
	return d
}

// CreateHTMLDocument
func (*DOMImplementation) CreateHTMLDocument(title string) (d Document) {
	js.Rewrite("$<.CreateHTMLDocument($1)", title)
	return d
}

// HasFeature
func (*DOMImplementation) HasFeature() (b bool) {
	js.Rewrite("$<.HasFeature()")
	return b
}
