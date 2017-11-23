package xsltprocessor

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *XSLTProcessor {
	js.Rewrite("XSLTProcessor")
	return &XSLTProcessor{}
}

// XSLTProcessor struct
// js:"XSLTProcessor,omit"
type XSLTProcessor struct {
}

// ClearParameters fn
func (*XSLTProcessor) ClearParameters() {
	js.Rewrite("$<.clearParameters()")
}

// GetParameter fn
func (*XSLTProcessor) GetParameter(namespaceURI string, localName string) (i interface{}) {
	js.Rewrite("$<.getParameter($1, $2)", namespaceURI, localName)
	return i
}

// ImportStylesheet fn
func (*XSLTProcessor) ImportStylesheet(style window.Node) {
	js.Rewrite("$<.importStylesheet($1)", style)
}

// RemoveParameter fn
func (*XSLTProcessor) RemoveParameter(namespaceURI string, localName string) {
	js.Rewrite("$<.removeParameter($1, $2)", namespaceURI, localName)
}

// Reset fn
func (*XSLTProcessor) Reset() {
	js.Rewrite("$<.reset()")
}

// SetParameter fn
func (*XSLTProcessor) SetParameter(namespaceURI string, localName string, value interface{}) {
	js.Rewrite("$<.setParameter($1, $2, $3)", namespaceURI, localName, value)
}

// TransformToDocument fn
func (*XSLTProcessor) TransformToDocument(source window.Node) (w window.Document) {
	js.Rewrite("$<.transformToDocument($1)", source)
	return w
}

// TransformToFragment fn
func (*XSLTProcessor) TransformToFragment(source window.Node, document window.Document) (w *window.DocumentFragment) {
	js.Rewrite("$<.transformToFragment($1, $2)", source, document)
	return w
}
