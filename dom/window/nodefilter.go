package window

import "github.com/matthewmueller/golly/js"

// NodeFilter struct
// js:"NodeFilter,omit"
type NodeFilter struct {
}

// AcceptNode fn
func (*NodeFilter) AcceptNode(n Node) (i int8) {
	js.Rewrite("$<.acceptNode($1)", n)
	return i
}
