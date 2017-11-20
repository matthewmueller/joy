package xpathevaluator

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/xpathexpression"
	"github.com/matthewmueller/golly/internal/gendom/dom/xpathnsresolver"
	"github.com/matthewmueller/golly/internal/gendom/dom/xpathresult"
	"github.com/matthewmueller/golly/js"
)

type XPathEvaluator struct {
}

func (*XPathEvaluator) CreateExpression(expression string, resolver *xpathnsresolver.XPathNSResolver) (x *xpathexpression.XPathExpression) {
	js.Rewrite("$<.createExpression($1, $2)", expression, resolver)
	return x
}

func (*XPathEvaluator) CreateNSResolver(nodeResolver *node.Node) (x *xpathnsresolver.XPathNSResolver) {
	js.Rewrite("$<.createNSResolver($1)", nodeResolver)
	return x
}

func (*XPathEvaluator) Evaluate(expression string, contextNode *node.Node, resolver *xpathnsresolver.XPathNSResolver, kind uint8, result *xpathresult.XPathResult) (x *xpathresult.XPathResult) {
	js.Rewrite("$<.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return x
}
