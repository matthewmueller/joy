package window

import "github.com/matthewmueller/golly/js"

// ElementTraversal struct
// js:"ElementTraversal,omit"
type ElementTraversal struct {
}

// ChildElementCount
func (*ElementTraversal) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$<.ChildElementCount")
	return childElementCount
}

// FirstElementChild
func (*ElementTraversal) FirstElementChild() (firstElementChild Element) {
	js.Rewrite("$<.FirstElementChild")
	return firstElementChild
}

// LastElementChild
func (*ElementTraversal) LastElementChild() (lastElementChild Element) {
	js.Rewrite("$<.LastElementChild")
	return lastElementChild
}

// NextElementSibling
func (*ElementTraversal) NextElementSibling() (nextElementSibling Element) {
	js.Rewrite("$<.NextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling
func (*ElementTraversal) PreviousElementSibling() (previousElementSibling Element) {
	js.Rewrite("$<.PreviousElementSibling")
	return previousElementSibling
}
