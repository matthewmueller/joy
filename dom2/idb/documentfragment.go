package idb

import "github.com/matthewmueller/golly/js"

// DocumentFragment struct
// js:"DocumentFragment,omit"
type DocumentFragment struct {
	Node
}

// QuerySelector
func (*DocumentFragment) QuerySelector(selectors string) (e Element) {
	js.Rewrite("$<.QuerySelector($1)", selectors)
	return e
}

// QuerySelectorAll
func (*DocumentFragment) QuerySelectorAll(selectors string) (n *NodeList) {
	js.Rewrite("$<.QuerySelectorAll($1)", selectors)
	return n
}
