package window

import "github.com/matthewmueller/golly/js"

// ElementTraversal struct
// js:"ElementTraversal,omit"
type ElementTraversal struct {
}

// ChildElementCount prop
func (*ElementTraversal) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$<.childElementCount")
	return childElementCount
}

// FirstElementChild prop
func (*ElementTraversal) FirstElementChild() (firstElementChild Element) {
	js.Rewrite("$<.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
func (*ElementTraversal) LastElementChild() (lastElementChild Element) {
	js.Rewrite("$<.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
func (*ElementTraversal) NextElementSibling() (nextElementSibling Element) {
	js.Rewrite("$<.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
func (*ElementTraversal) PreviousElementSibling() (previousElementSibling Element) {
	js.Rewrite("$<.previousElementSibling")
	return previousElementSibling
}
