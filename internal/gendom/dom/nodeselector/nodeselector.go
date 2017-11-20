package nodeselector

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodelist"
	"github.com/matthewmueller/golly/js"
)

type NodeSelector struct {
}

func (*NodeSelector) QuerySelector(selectors string) (e *element.Element) {
	js.Rewrite("$<.querySelector($1)", selectors)
	return e
}

func (*NodeSelector) QuerySelectorAll(selectors string) (n *nodelist.NodeList) {
	js.Rewrite("$<.querySelectorAll($1)", selectors)
	return n
}
