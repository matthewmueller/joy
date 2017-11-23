package htmloptionscollection

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLOptionsCollection struct
// js:"HTMLOptionsCollection,omit"
type HTMLOptionsCollection struct {
	window.HTMLCollection
}

// Add fn
func (*HTMLOptionsCollection) Add(element interface{}, before *interface{}) {
	js.Rewrite("$<.add($1, $2)", element, before)
}

// Remove fn
func (*HTMLOptionsCollection) Remove(index int) {
	js.Rewrite("$<.remove($1)", index)
}

// Length prop
func (*HTMLOptionsCollection) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Length prop
func (*HTMLOptionsCollection) SetLength(length uint) {
	js.Rewrite("$<.length = $1", length)
}

// SelectedIndex prop
func (*HTMLOptionsCollection) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}

// SelectedIndex prop
func (*HTMLOptionsCollection) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.selectedIndex = $1", selectedIndex)
}
