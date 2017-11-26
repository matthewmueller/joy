package window

// DocumentFragment struct
// js:"DocumentFragment,omit"
type DocumentFragment struct {
	Node
}

// // QuerySelector fn
// func (*DocumentFragment) QuerySelector(selectors string) (e Element) {
// 	js.Rewrite("$<.querySelector($1)", selectors)
// 	return e
// }

// // QuerySelectorAll fn
// func (*DocumentFragment) QuerySelectorAll(selectors string) (n *NodeList) {
// 	js.Rewrite("$<.querySelectorAll($1)", selectors)
// 	return n
// }
