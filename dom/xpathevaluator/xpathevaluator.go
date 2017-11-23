package xpathevaluator

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/dom/xpathnsresolver"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *XPathEvaluator {
	js.Rewrite("XPathEvaluator")
	return &XPathEvaluator{}
}

// XPathEvaluator struct
// js:"XPathEvaluator,omit"
type XPathEvaluator struct {
}

// CreateExpression fn
func (*XPathEvaluator) CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (w *window.XPathExpression) {
	js.Rewrite("$<.createExpression($1, $2)", expression, resolver)
	return w
}

// CreateNSResolver fn
func (*XPathEvaluator) CreateNSResolver(nodeResolver *window.Node) (x *xpathnsresolver.XPathNSResolver) {
	js.Rewrite("$<.createNSResolver($1)", nodeResolver)
	return x
}

// Evaluate fn
func (*XPathEvaluator) Evaluate(expression string, contextNode window.Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *window.XPathResult) (w *window.XPathResult) {
	js.Rewrite("$<.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}
