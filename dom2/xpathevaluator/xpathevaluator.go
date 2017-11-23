package xpathevaluator

import "github.com/matthewmueller/golly/js"

// js:"XPathEvaluator,omit"
type XPathEvaluator struct {
}

// CreateExpression
func (*XPathEvaluator) CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (w *window.XPathExpression) {
	js.Rewrite("$<.CreateExpression($1, $2)", expression, resolver)
	return w
}

// CreateNSResolver
func (*XPathEvaluator) CreateNSResolver(nodeResolver *window.Node) (x *xpathnsresolver.XPathNSResolver) {
	js.Rewrite("$<.CreateNSResolver($1)", nodeResolver)
	return x
}

// Evaluate
func (*XPathEvaluator) Evaluate(expression string, contextNode window.Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *window.XPathResult) (w *window.XPathResult) {
	js.Rewrite("$<.Evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return w
}
