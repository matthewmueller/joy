package window

import "github.com/matthewmueller/golly/js"

// TreeWalker struct
// js:"TreeWalker,omit"
type TreeWalker struct {
}

// FirstChild
func (*TreeWalker) FirstChild() (n Node) {
	js.Rewrite("$<.FirstChild()")
	return n
}

// LastChild
func (*TreeWalker) LastChild() (n Node) {
	js.Rewrite("$<.LastChild()")
	return n
}

// NextNode
func (*TreeWalker) NextNode() (n Node) {
	js.Rewrite("$<.NextNode()")
	return n
}

// NextSibling
func (*TreeWalker) NextSibling() (n Node) {
	js.Rewrite("$<.NextSibling()")
	return n
}

// ParentNode
func (*TreeWalker) ParentNode() (n Node) {
	js.Rewrite("$<.ParentNode()")
	return n
}

// PreviousNode
func (*TreeWalker) PreviousNode() (n Node) {
	js.Rewrite("$<.PreviousNode()")
	return n
}

// PreviousSibling
func (*TreeWalker) PreviousSibling() (n Node) {
	js.Rewrite("$<.PreviousSibling()")
	return n
}

// CurrentNode
func (*TreeWalker) CurrentNode() (currentNode Node) {
	js.Rewrite("$<.CurrentNode")
	return currentNode
}

// CurrentNode
func (*TreeWalker) SetCurrentNode(currentNode Node) {
	js.Rewrite("$<.CurrentNode = $1", currentNode)
}

// ExpandEntityReferences
func (*TreeWalker) ExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.ExpandEntityReferences")
	return expandEntityReferences
}

// Filter
func (*TreeWalker) Filter() (filter *NodeFilter) {
	js.Rewrite("$<.Filter")
	return filter
}

// Root
func (*TreeWalker) Root() (root Node) {
	js.Rewrite("$<.Root")
	return root
}

// WhatToShow
func (*TreeWalker) WhatToShow() (whatToShow uint) {
	js.Rewrite("$<.WhatToShow")
	return whatToShow
}
