package window

import "github.com/matthewmueller/golly/js"

// js:"NodeIterator,omit"
type NodeIterator struct {
}

// Detach
func (*NodeIterator) Detach() {
	js.Rewrite("$<.Detach()")
}

// NextNode
func (*NodeIterator) NextNode() (n Node) {
	js.Rewrite("$<.NextNode()")
	return n
}

// PreviousNode
func (*NodeIterator) PreviousNode() (n Node) {
	js.Rewrite("$<.PreviousNode()")
	return n
}

// ExpandEntityReferences
func (*NodeIterator) ExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.ExpandEntityReferences")
	return expandEntityReferences
}

// Filter
func (*NodeIterator) Filter() (filter *NodeFilter) {
	js.Rewrite("$<.Filter")
	return filter
}

// Root
func (*NodeIterator) Root() (root Node) {
	js.Rewrite("$<.Root")
	return root
}

// WhatToShow
func (*NodeIterator) WhatToShow() (whatToShow uint) {
	js.Rewrite("$<.WhatToShow")
	return whatToShow
}
