package window

import "github.com/matthewmueller/joy/js"

// NodeIterator struct
// js:"NodeIterator,omit"
type NodeIterator struct {
}

// Detach fn
// js:"detach"
func (*NodeIterator) Detach() {
	js.Rewrite("$_.detach()")
}

// NextNode fn
// js:"nextNode"
func (*NodeIterator) NextNode() (n Node) {
	js.Rewrite("$_.nextNode()")
	return n
}

// PreviousNode fn
// js:"previousNode"
func (*NodeIterator) PreviousNode() (n Node) {
	js.Rewrite("$_.previousNode()")
	return n
}

// ExpandEntityReferences prop
// js:"expandEntityReferences"
func (*NodeIterator) ExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$_.expandEntityReferences")
	return expandEntityReferences
}

// Filter prop
// js:"filter"
func (*NodeIterator) Filter() (filter *NodeFilter) {
	js.Rewrite("$_.filter")
	return filter
}

// Root prop
// js:"root"
func (*NodeIterator) Root() (root Node) {
	js.Rewrite("$_.root")
	return root
}

// WhatToShow prop
// js:"whatToShow"
func (*NodeIterator) WhatToShow() (whatToShow uint) {
	js.Rewrite("$_.whatToShow")
	return whatToShow
}
