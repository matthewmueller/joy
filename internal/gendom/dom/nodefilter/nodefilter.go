package nodefilter

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type NodeFilter struct {
}

func (*NodeFilter) AcceptNode(n *node.Node) (i int8) {
	js.Rewrite("$<.acceptNode($1)", n)
	return i
}
