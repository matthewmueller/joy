package idb

import "github.com/matthewmueller/golly/js"

// NodeSelector struct
// js:"NodeSelector,omit"
type NodeSelector struct {
}

// QuerySelector
func (*NodeSelector) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$<.QuerySelector($1)", selectors)
	return e
}

// QuerySelectorAll
func (*NodeSelector) QuerySelectorAll(selectors string) (n *NodeList) {
	js.Rewrite("$<.QuerySelectorAll($1)", selectors)
	return n
}
