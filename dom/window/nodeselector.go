package window

import "github.com/matthewmueller/golly/js"

// NodeSelector struct
// js:"NodeSelector,omit"
type NodeSelector struct {
}

// QuerySelector fn
func (*NodeSelector) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$<.querySelector($1)", selectors)
	return e
}

// QuerySelectorAll fn
func (*NodeSelector) QuerySelectorAll(selectors string) (n *NodeList) {
	js.Rewrite("$<.querySelectorAll($1)", selectors)
	return n
}
