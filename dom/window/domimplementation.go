package window

import "github.com/matthewmueller/joy/js"

// DOMImplementation struct
// js:"DOMImplementation,omit"
type DOMImplementation struct {
}

// CreateDocument fn
// js:"createDocument"
func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *DocumentType) (d Document) {
	js.Rewrite("$_.createDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

// CreateDocumentType fn
// js:"createDocumentType"
func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) (d *DocumentType) {
	js.Rewrite("$_.createDocumentType($1, $2, $3)", qualifiedName, publicId, systemId)
	return d
}

// CreateHTMLDocument fn
// js:"createHTMLDocument"
func (*DOMImplementation) CreateHTMLDocument(title string) (d Document) {
	js.Rewrite("$_.createHTMLDocument($1)", title)
	return d
}

// HasFeature fn
// js:"hasFeature"
func (*DOMImplementation) HasFeature() (b bool) {
	js.Rewrite("$_.hasFeature()")
	return b
}
