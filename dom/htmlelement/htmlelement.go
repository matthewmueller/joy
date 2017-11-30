package htmlelement

import "github.com/matthewmueller/golly/dom/element"

// HTMLElement interface
// js:"HTMLElement"
type HTMLElement interface {
	element.Element

	// Blur
	// js:"blur"
	Blur()

	// Focus
	// js:"focus"
	Focus()
}
