package window

import "github.com/matthewmueller/joy/js"

// NodeFilter struct
// js:"NodeFilter,omit"
type NodeFilter struct {
}

// AcceptNode fn
// js:"acceptNode"
func (*NodeFilter) AcceptNode(n Node) (i int8) {
	js.Rewrite("$_.acceptNode($1)", n)
	return i
}
