package window

import "github.com/matthewmueller/golly/js"

// NodeIterator struct
// js:"NodeIterator,omit"
type NodeIterator struct {
}

// Detach fn
func (*NodeIterator) Detach() {
	js.Rewrite("$<.detach()")
}

// NextNode fn
func (*NodeIterator) NextNode() (n Node) {
	js.Rewrite("$<.nextNode()")
	return n
}

// PreviousNode fn
func (*NodeIterator) PreviousNode() (n Node) {
	js.Rewrite("$<.previousNode()")
	return n
}

// ExpandEntityReferences prop
func (*NodeIterator) ExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.expandEntityReferences")
	return expandEntityReferences
}

// Filter prop
func (*NodeIterator) Filter() (filter *NodeFilter) {
	js.Rewrite("$<.filter")
	return filter
}

// Root prop
func (*NodeIterator) Root() (root Node) {
	js.Rewrite("$<.root")
	return root
}

// WhatToShow prop
func (*NodeIterator) WhatToShow() (whatToShow uint) {
	js.Rewrite("$<.whatToShow")
	return whatToShow
}
