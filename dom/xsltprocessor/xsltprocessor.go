package xsltprocessor

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New() *XSLTProcessor {
	macro.Rewrite("XSLTProcessor")
	return &XSLTProcessor{}
}

// XSLTProcessor struct
// js:"XSLTProcessor,omit"
type XSLTProcessor struct {
}

// ClearParameters fn
// js:"clearParameters"
func (*XSLTProcessor) ClearParameters() {
	macro.Rewrite("$_.clearParameters()")
}

// GetParameter fn
// js:"getParameter"
func (*XSLTProcessor) GetParameter(namespaceURI string, localName string) (i interface{}) {
	macro.Rewrite("$_.getParameter($1, $2)", namespaceURI, localName)
	return i
}

// ImportStylesheet fn
// js:"importStylesheet"
func (*XSLTProcessor) ImportStylesheet(style window.Node) {
	macro.Rewrite("$_.importStylesheet($1)", style)
}

// RemoveParameter fn
// js:"removeParameter"
func (*XSLTProcessor) RemoveParameter(namespaceURI string, localName string) {
	macro.Rewrite("$_.removeParameter($1, $2)", namespaceURI, localName)
}

// Reset fn
// js:"reset"
func (*XSLTProcessor) Reset() {
	macro.Rewrite("$_.reset()")
}

// SetParameter fn
// js:"setParameter"
func (*XSLTProcessor) SetParameter(namespaceURI string, localName string, value interface{}) {
	macro.Rewrite("$_.setParameter($1, $2, $3)", namespaceURI, localName, value)
}

// TransformToDocument fn
// js:"transformToDocument"
func (*XSLTProcessor) TransformToDocument(source window.Node) (w window.Document) {
	macro.Rewrite("$_.transformToDocument($1)", source)
	return w
}

// TransformToFragment fn
// js:"transformToFragment"
func (*XSLTProcessor) TransformToFragment(source window.Node, document window.Document) (w *window.DocumentFragment) {
	macro.Rewrite("$_.transformToFragment($1, $2)", source, document)
	return w
}
