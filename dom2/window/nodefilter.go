package window

import "github.com/matthewmueller/golly/js"

// NodeFilter struct
// js:"NodeFilter,omit"
type NodeFilter struct {
}

// AcceptNode
func (*NodeFilter) AcceptNode(n Node) (i int8) {
	js.Rewrite("$<.AcceptNode($1)", n)
	return i
}
