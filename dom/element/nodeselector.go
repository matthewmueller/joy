package element

// NodeSelector interface
// js:"NodeSelector"
type NodeSelector interface {

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)
}
