package xsltprocessor

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/documentfragment"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type XSLTProcessor struct {
}

func (*XSLTProcessor) ClearParameters() {
	js.Rewrite("$<.clearParameters()")
}

func (*XSLTProcessor) GetParameter(namespaceURI string, localName string) (i interface{}) {
	js.Rewrite("$<.getParameter($1, $2)", namespaceURI, localName)
	return i
}

func (*XSLTProcessor) ImportStylesheet(style *node.Node) {
	js.Rewrite("$<.importStylesheet($1)", style)
}

func (*XSLTProcessor) RemoveParameter(namespaceURI string, localName string) {
	js.Rewrite("$<.removeParameter($1, $2)", namespaceURI, localName)
}

func (*XSLTProcessor) Reset() {
	js.Rewrite("$<.reset()")
}

func (*XSLTProcessor) SetParameter(namespaceURI string, localName string, value interface{}) {
	js.Rewrite("$<.setParameter($1, $2, $3)", namespaceURI, localName, value)
}

func (*XSLTProcessor) TransformToDocument(source *node.Node) (d *document.Document) {
	js.Rewrite("$<.transformToDocument($1)", source)
	return d
}

func (*XSLTProcessor) TransformToFragment(source *node.Node, document *document.Document) (d *documentfragment.DocumentFragment) {
	js.Rewrite("$<.transformToFragment($1, $2)", source, document)
	return d
}
