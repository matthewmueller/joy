package treewalker

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/nodefilter"
	"github.com/matthewmueller/golly/js"
)

type TreeWalker struct {
}

func (*TreeWalker) FirstChild() (n *node.Node) {
	js.Rewrite("$<.firstChild()")
	return n
}

func (*TreeWalker) LastChild() (n *node.Node) {
	js.Rewrite("$<.lastChild()")
	return n
}

func (*TreeWalker) NextNode() (n *node.Node) {
	js.Rewrite("$<.nextNode()")
	return n
}

func (*TreeWalker) NextSibling() (n *node.Node) {
	js.Rewrite("$<.nextSibling()")
	return n
}

func (*TreeWalker) ParentNode() (n *node.Node) {
	js.Rewrite("$<.parentNode()")
	return n
}

func (*TreeWalker) PreviousNode() (n *node.Node) {
	js.Rewrite("$<.previousNode()")
	return n
}

func (*TreeWalker) PreviousSibling() (n *node.Node) {
	js.Rewrite("$<.previousSibling()")
	return n
}

func (*TreeWalker) GetCurrentNode() (currentNode *node.Node) {
	js.Rewrite("$<.currentNode")
	return currentNode
}

func (*TreeWalker) SetCurrentNode(currentNode *node.Node) {
	js.Rewrite("$<.currentNode = $1", currentNode)
}

func (*TreeWalker) GetExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.expandEntityReferences")
	return expandEntityReferences
}

func (*TreeWalker) GetFilter() (filter *nodefilter.NodeFilter) {
	js.Rewrite("$<.filter")
	return filter
}

func (*TreeWalker) GetRoot() (root *node.Node) {
	js.Rewrite("$<.root")
	return root
}

func (*TreeWalker) GetWhatToShow() (whatToShow uint) {
	js.Rewrite("$<.whatToShow")
	return whatToShow
}
