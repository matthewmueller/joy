package window

import "github.com/matthewmueller/joy/macro"

// NodeFilter struct
// js:"NodeFilter,omit"
type NodeFilter struct {
}

// AcceptNode fn
// js:"acceptNode"
func (*NodeFilter) AcceptNode(n Node) (i int8) {
	macro.Rewrite("$_.acceptNode($1)", n)
	return i
}
