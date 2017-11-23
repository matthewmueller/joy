package xsltprocessor

import (
	"github.com/matthewmueller/golly/dom2/window"
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

// ClearParameters
func (*XSLTProcessor) ClearParameters() {
	js.Rewrite("$<.ClearParameters()")
}

// GetParameter
func (*XSLTProcessor) GetParameter(namespaceURI string, localName string) (i interface{}) {
	js.Rewrite("$<.GetParameter($1, $2)", namespaceURI, localName)
	return i
}

// ImportStylesheet
func (*XSLTProcessor) ImportStylesheet(style window.Node) {
	js.Rewrite("$<.ImportStylesheet($1)", style)
}

// RemoveParameter
func (*XSLTProcessor) RemoveParameter(namespaceURI string, localName string) {
	js.Rewrite("$<.RemoveParameter($1, $2)", namespaceURI, localName)
}

// Reset
func (*XSLTProcessor) Reset() {
	js.Rewrite("$<.Reset()")
}

// SetParameter
func (*XSLTProcessor) SetParameter(namespaceURI string, localName string, value interface{}) {
	js.Rewrite("$<.SetParameter($1, $2, $3)", namespaceURI, localName, value)
}

// TransformToDocument
func (*XSLTProcessor) TransformToDocument(source window.Node) (w window.Document) {
	js.Rewrite("$<.TransformToDocument($1)", source)
	return w
}

// TransformToFragment
func (*XSLTProcessor) TransformToFragment(source window.Node, document window.Document) (w *window.DocumentFragment) {
	js.Rewrite("$<.TransformToFragment($1, $2)", source, document)
	return w
}
