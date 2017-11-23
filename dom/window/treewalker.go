package window

import "github.com/matthewmueller/golly/js"

// TreeWalker struct
// js:"TreeWalker,omit"
type TreeWalker struct {
}

// FirstChild fn
func (*TreeWalker) FirstChild() (n Node) {
	js.Rewrite("$<.firstChild()")
	return n
}

// LastChild fn
func (*TreeWalker) LastChild() (n Node) {
	js.Rewrite("$<.lastChild()")
	return n
}

// NextNode fn
func (*TreeWalker) NextNode() (n Node) {
	js.Rewrite("$<.nextNode()")
	return n
}

// NextSibling fn
func (*TreeWalker) NextSibling() (n Node) {
	js.Rewrite("$<.nextSibling()")
	return n
}

// ParentNode fn
func (*TreeWalker) ParentNode() (n Node) {
	js.Rewrite("$<.parentNode()")
	return n
}

// PreviousNode fn
func (*TreeWalker) PreviousNode() (n Node) {
	js.Rewrite("$<.previousNode()")
	return n
}

// PreviousSibling fn
func (*TreeWalker) PreviousSibling() (n Node) {
	js.Rewrite("$<.previousSibling()")
	return n
}

// CurrentNode prop
func (*TreeWalker) CurrentNode() (currentNode Node) {
	js.Rewrite("$<.currentNode")
	return currentNode
}

// CurrentNode prop
func (*TreeWalker) SetCurrentNode(currentNode Node) {
	js.Rewrite("$<.currentNode = $1", currentNode)
}

// ExpandEntityReferences prop
func (*TreeWalker) ExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.expandEntityReferences")
	return expandEntityReferences
}

// Filter prop
func (*TreeWalker) Filter() (filter *NodeFilter) {
	js.Rewrite("$<.filter")
	return filter
}

// Root prop
func (*TreeWalker) Root() (root Node) {
	js.Rewrite("$<.root")
	return root
}

// WhatToShow prop
func (*TreeWalker) WhatToShow() (whatToShow uint) {
	js.Rewrite("$<.whatToShow")
	return whatToShow
}
