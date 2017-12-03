package window

import "github.com/matthewmueller/joy/macro"

// NodeIterator struct
// js:"NodeIterator,omit"
type NodeIterator struct {
}

// Detach fn
// js:"detach"
func (*NodeIterator) Detach() {
	macro.Rewrite("$_.detach()")
}

// NextNode fn
// js:"nextNode"
func (*NodeIterator) NextNode() (n Node) {
	macro.Rewrite("$_.nextNode()")
	return n
}

// PreviousNode fn
// js:"previousNode"
func (*NodeIterator) PreviousNode() (n Node) {
	macro.Rewrite("$_.previousNode()")
	return n
}

// ExpandEntityReferences prop
// js:"expandEntityReferences"
func (*NodeIterator) ExpandEntityReferences() (expandEntityReferences bool) {
	macro.Rewrite("$_.expandEntityReferences")
	return expandEntityReferences
}

// Filter prop
// js:"filter"
func (*NodeIterator) Filter() (filter *NodeFilter) {
	macro.Rewrite("$_.filter")
	return filter
}

// Root prop
// js:"root"
func (*NodeIterator) Root() (root Node) {
	macro.Rewrite("$_.root")
	return root
}

// WhatToShow prop
// js:"whatToShow"
func (*NodeIterator) WhatToShow() (whatToShow uint) {
	macro.Rewrite("$_.whatToShow")
	return whatToShow
}
