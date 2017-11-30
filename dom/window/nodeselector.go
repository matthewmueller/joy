package window

// NodeSelector interface
// js:"NodeSelector"
type NodeSelector interface {

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll"
	QuerySelectorAll(selectors string) (n *NodeList)
}
