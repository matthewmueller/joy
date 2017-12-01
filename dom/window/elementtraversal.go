package window

// ElementTraversal interface
// js:"ElementTraversal"
type ElementTraversal interface {

	// ChildElementCount prop
	// js:"childElementCount"
	// jsrewrite:"$_.childElementCount"
	ChildElementCount() (childElementCount uint)

	// FirstElementChild prop
	// js:"firstElementChild"
	// jsrewrite:"$_.firstElementChild"
	FirstElementChild() (firstElementChild Element)

	// LastElementChild prop
	// js:"lastElementChild"
	// jsrewrite:"$_.lastElementChild"
	LastElementChild() (lastElementChild Element)

	// NextElementSibling prop
	// js:"nextElementSibling"
	// jsrewrite:"$_.nextElementSibling"
	NextElementSibling() (nextElementSibling Element)

	// PreviousElementSibling prop
	// js:"previousElementSibling"
	// jsrewrite:"$_.previousElementSibling"
	PreviousElementSibling() (previousElementSibling Element)
}
