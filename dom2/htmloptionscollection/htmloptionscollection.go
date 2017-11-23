package htmloptionscollection

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLOptionsCollection,omit"
type HTMLOptionsCollection struct {
	window.HTMLCollection
}

// Add
func (*HTMLOptionsCollection) Add(element interface{}, before *interface{}) {
	js.Rewrite("$<.Add($1, $2)", element, before)
}

// Remove
func (*HTMLOptionsCollection) Remove(index int) {
	js.Rewrite("$<.Remove($1)", index)
}

// Length
func (*HTMLOptionsCollection) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Length
func (*HTMLOptionsCollection) SetLength(length uint) {
	js.Rewrite("$<.Length = $1", length)
}

// SelectedIndex
func (*HTMLOptionsCollection) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.SelectedIndex")
	return selectedIndex
}

// SelectedIndex
func (*HTMLOptionsCollection) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.SelectedIndex = $1", selectedIndex)
}
