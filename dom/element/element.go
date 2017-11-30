package element

import "github.com/matthewmueller/golly/dom/node"

// Element interface
// js:"Element"
type Element interface {
	node.Node

	// QuerySelector
	// js:"querySelector"
	QuerySelector(selectors string) (e Element)

	// ClassName prop
	// js:"className"
	ClassName() (className string)

	// SetClassName prop
	// js:"className"
	SetClassName(className string)

	// ID prop
	// js:"id"
	ID() (id string)

	// SetID prop
	// js:"id"
	SetID(id string)

	// InnerHTML prop
	// js:"innerHTML"
	InnerHTML() (innerHTML string)

	// SetInnerHTML prop
	// js:"innerHTML"
	SetInnerHTML(innerHTML string)

	// OuterHTML prop
	// js:"outerHTML"
	OuterHTML() (outerHTML string)

	// SetOuterHTML prop
	// js:"outerHTML"
	SetOuterHTML(outerHTML string)

	// ScrollHeight prop
	// js:"scrollHeight"
	ScrollHeight() (scrollHeight int)

	// ScrollLeft prop
	// js:"scrollLeft"
	ScrollLeft() (scrollLeft int)

	// SetScrollLeft prop
	// js:"scrollLeft"
	SetScrollLeft(scrollLeft int)

	// ScrollTop prop
	// js:"scrollTop"
	ScrollTop() (scrollTop int)

	// SetScrollTop prop
	// js:"scrollTop"
	SetScrollTop(scrollTop int)
}
