package window

import "github.com/matthewmueller/joy/macro"

// DOMImplementation struct
// js:"DOMImplementation,omit"
type DOMImplementation struct {
}

// CreateDocument fn
// js:"createDocument"
func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *DocumentType) (d Document) {
	macro.Rewrite("$_.createDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

// CreateDocumentType fn
// js:"createDocumentType"
func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) (d *DocumentType) {
	macro.Rewrite("$_.createDocumentType($1, $2, $3)", qualifiedName, publicId, systemId)
	return d
}

// CreateHTMLDocument fn
// js:"createHTMLDocument"
func (*DOMImplementation) CreateHTMLDocument(title string) (d Document) {
	macro.Rewrite("$_.createHTMLDocument($1)", title)
	return d
}

// HasFeature fn
// js:"hasFeature"
func (*DOMImplementation) HasFeature() (b bool) {
	macro.Rewrite("$_.hasFeature()")
	return b
}
