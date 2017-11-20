package nodeiterator

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodefilter"
	"github.com/matthewmueller/golly/js"
)

type NodeIterator struct {
}

func (*NodeIterator) Detach() {
	js.Rewrite("$<.detach()")
}

func (*NodeIterator) NextNode() (n *node.Node) {
	js.Rewrite("$<.nextNode()")
	return n
}

func (*NodeIterator) PreviousNode() (n *node.Node) {
	js.Rewrite("$<.previousNode()")
	return n
}

func (*NodeIterator) GetExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.expandEntityReferences")
	return expandEntityReferences
}

func (*NodeIterator) GetFilter() (filter *nodefilter.NodeFilter) {
	js.Rewrite("$<.filter")
	return filter
}

func (*NodeIterator) GetRoot() (root *node.Node) {
	js.Rewrite("$<.root")
	return root
}

func (*NodeIterator) GetWhatToShow() (whatToShow uint) {
	js.Rewrite("$<.whatToShow")
	return whatToShow
}
