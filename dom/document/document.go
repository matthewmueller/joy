package document

import (
	"github.com/matthewmueller/golly/dom/element"
	"github.com/matthewmueller/golly/dom/node"
)

// Document interface
// js:"Document"
type Document interface {
	node.Node

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e element.Element)

	// DocumentElement prop Gets a reference to the root node of the document.
	// js:"documentElement"
	DocumentElement() (documentElement element.Element)
}
