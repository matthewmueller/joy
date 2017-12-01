package window

// NodeSelector interface
// js:"NodeSelector"
type NodeSelector interface {

	// QuerySelector
	// js:"querySelector"
	// jsrewrite:"$_.querySelector($1)"
	QuerySelector(selectors string) (e Element)

	// QuerySelectorAll
	// js:"querySelectorAll"
	// jsrewrite:"$_.querySelectorAll($1)"
	QuerySelectorAll(selectors string) (n *NodeList)
}
