package xpathexpression

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/xpathresult"
	"github.com/matthewmueller/golly/js"
)

type XPathExpression struct {
}

func (*XPathExpression) Evaluate(contextNode *node.Node, kind uint8, result *xpathresult.XPathResult) (x *XPathExpression) {
	js.Rewrite("$<.evaluate($1, $2, $3)", contextNode, kind, result)
	return x
}
