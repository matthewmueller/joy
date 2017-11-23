package childnode

import "github.com/matthewmueller/golly/js"

// ChildNode struct
// js:"ChildNode,omit"
type ChildNode struct {
}

// Remove
func (*ChildNode) Remove() {
	js.Rewrite("$<.Remove()")
}
