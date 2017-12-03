package xpathevaluator

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/xpathnsresolver"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New() *XPathEvaluator {
	macro.Rewrite("XPathEvaluator")
	return &XPathEvaluator{}
}

// XPathEvaluator struct
// js:"XPathEvaluator,omit"
type XPathEvaluator struct {
}

// CreateExpression fn
// js:"createExpression"
func (*XPathEvaluator) CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (w *window.XPathExpression) {
	macro.Rewrite("$_.createExpression($1, $2)", expression, resolver)
	return w
}

// CreateNSResolver fn
// js:"createNSResolver"
func (*XPathEvaluator) CreateNSResolver(nodeResolver *window.Node) (x *xpathnsresolver.XPathNSResolver) {
	macro.Rewrite("$_.createNSResolver($1)", nodeResolver)
	return x
}

// Evaluate fn
// js:"evaluate"
func (*XPathEvaluator) Evaluate(expression string, contextNode window.Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *window.XPathResult) (w *window.XPathResult) {
	macro.Rewrite("$_.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}
