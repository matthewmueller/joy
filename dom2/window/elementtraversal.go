package window

// js:"ElementTraversal,omit"
type ElementTraversal interface {

	// ChildElementCount
	ChildElementCount() (childElementCount uint)

	// FirstElementChild
	FirstElementChild() (firstElementChild Element)

	// LastElementChild
	LastElementChild() (lastElementChild Element)

	// NextElementSibling
	NextElementSibling() (nextElementSibling Element)

	// PreviousElementSibling
	PreviousElementSibling() (previousElementSibling Element)
}
