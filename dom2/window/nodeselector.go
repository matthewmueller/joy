package window

// js:"NodeSelector,omit"
type NodeSelector interface {

	// QuerySelector
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	QuerySelectorAll(selectors string) (n *NodeList)
}
