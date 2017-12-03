package window

import "github.com/matthewmueller/joy/macro"

// TreeWalker struct
// js:"TreeWalker,omit"
type TreeWalker struct {
}

// FirstChild fn
// js:"firstChild"
func (*TreeWalker) FirstChild() (n Node) {
	macro.Rewrite("$_.firstChild()")
	return n
}

// LastChild fn
// js:"lastChild"
func (*TreeWalker) LastChild() (n Node) {
	macro.Rewrite("$_.lastChild()")
	return n
}

// NextNode fn
// js:"nextNode"
func (*TreeWalker) NextNode() (n Node) {
	macro.Rewrite("$_.nextNode()")
	return n
}

// NextSibling fn
// js:"nextSibling"
func (*TreeWalker) NextSibling() (n Node) {
	macro.Rewrite("$_.nextSibling()")
	return n
}

// ParentNode fn
// js:"parentNode"
func (*TreeWalker) ParentNode() (n Node) {
	macro.Rewrite("$_.parentNode()")
	return n
}

// PreviousNode fn
// js:"previousNode"
func (*TreeWalker) PreviousNode() (n Node) {
	macro.Rewrite("$_.previousNode()")
	return n
}

// PreviousSibling fn
// js:"previousSibling"
func (*TreeWalker) PreviousSibling() (n Node) {
	macro.Rewrite("$_.previousSibling()")
	return n
}

// CurrentNode prop
// js:"currentNode"
func (*TreeWalker) CurrentNode() (currentNode Node) {
	macro.Rewrite("$_.currentNode")
	return currentNode
}

// SetCurrentNode prop
// js:"currentNode"
func (*TreeWalker) SetCurrentNode(currentNode Node) {
	macro.Rewrite("$_.currentNode = $1", currentNode)
}

// ExpandEntityReferences prop
// js:"expandEntityReferences"
func (*TreeWalker) ExpandEntityReferences() (expandEntityReferences bool) {
	macro.Rewrite("$_.expandEntityReferences")
	return expandEntityReferences
}

// Filter prop
// js:"filter"
func (*TreeWalker) Filter() (filter *NodeFilter) {
	macro.Rewrite("$_.filter")
	return filter
}

// Root prop
// js:"root"
func (*TreeWalker) Root() (root Node) {
	macro.Rewrite("$_.root")
	return root
}

// WhatToShow prop
// js:"whatToShow"
func (*TreeWalker) WhatToShow() (whatToShow uint) {
	macro.Rewrite("$_.whatToShow")
	return whatToShow
}
