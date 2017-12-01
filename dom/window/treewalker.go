package window

import "github.com/matthewmueller/golly/js"

// TreeWalker struct
// js:"TreeWalker,omit"
type TreeWalker struct {
}

// FirstChild fn
// js:"firstChild"
func (*TreeWalker) FirstChild() (n Node) {
	js.Rewrite("$_.firstChild()")
	return n
}

// LastChild fn
// js:"lastChild"
func (*TreeWalker) LastChild() (n Node) {
	js.Rewrite("$_.lastChild()")
	return n
}

// NextNode fn
// js:"nextNode"
func (*TreeWalker) NextNode() (n Node) {
	js.Rewrite("$_.nextNode()")
	return n
}

// NextSibling fn
// js:"nextSibling"
func (*TreeWalker) NextSibling() (n Node) {
	js.Rewrite("$_.nextSibling()")
	return n
}

// ParentNode fn
// js:"parentNode"
func (*TreeWalker) ParentNode() (n Node) {
	js.Rewrite("$_.parentNode()")
	return n
}

// PreviousNode fn
// js:"previousNode"
func (*TreeWalker) PreviousNode() (n Node) {
	js.Rewrite("$_.previousNode()")
	return n
}

// PreviousSibling fn
// js:"previousSibling"
func (*TreeWalker) PreviousSibling() (n Node) {
	js.Rewrite("$_.previousSibling()")
	return n
}

// CurrentNode prop
// js:"currentNode"
func (*TreeWalker) CurrentNode() (currentNode Node) {
	js.Rewrite("$_.currentNode")
	return currentNode
}

// SetCurrentNode prop
// js:"currentNode"
func (*TreeWalker) SetCurrentNode(currentNode Node) {
	js.Rewrite("$_.currentNode = $1", currentNode)
}

// ExpandEntityReferences prop
// js:"expandEntityReferences"
func (*TreeWalker) ExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$_.expandEntityReferences")
	return expandEntityReferences
}

// Filter prop
// js:"filter"
func (*TreeWalker) Filter() (filter *NodeFilter) {
	js.Rewrite("$_.filter")
	return filter
}

// Root prop
// js:"root"
func (*TreeWalker) Root() (root Node) {
	js.Rewrite("$_.root")
	return root
}

// WhatToShow prop
// js:"whatToShow"
func (*TreeWalker) WhatToShow() (whatToShow uint) {
	js.Rewrite("$_.whatToShow")
	return whatToShow
}
