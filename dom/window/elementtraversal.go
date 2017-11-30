package window

// ElementTraversal interface
// js:"ElementTraversal"
type ElementTraversal interface {

	// ChildElementCount prop
	// js:"childElementCount"
	ChildElementCount() (childElementCount uint)

	// FirstElementChild prop
	// js:"firstElementChild"
	FirstElementChild() (firstElementChild Element)

	// LastElementChild prop
	// js:"lastElementChild"
	LastElementChild() (lastElementChild Element)

	// NextElementSibling prop
	// js:"nextElementSibling"
	NextElementSibling() (nextElementSibling Element)

	// PreviousElementSibling prop
	// js:"previousElementSibling"
	PreviousElementSibling() (previousElementSibling Element)
}
