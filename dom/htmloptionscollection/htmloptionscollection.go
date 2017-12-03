package htmloptionscollection

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.HTMLCollection = (*HTMLOptionsCollection)(nil)

// HTMLOptionsCollection struct
// js:"HTMLOptionsCollection,omit"
type HTMLOptionsCollection struct {
}

// Add fn
// js:"add"
func (*HTMLOptionsCollection) Add(element interface{}, before *interface{}) {
	macro.Rewrite("$_.add($1, $2)", element, before)
}

// Remove fn
// js:"remove"
func (*HTMLOptionsCollection) Remove(index int) {
	macro.Rewrite("$_.remove($1)", index)
}

// Item fn Retrieves an object from various collections.
// js:"item"
func (*HTMLOptionsCollection) Item(index uint) (w window.Element) {
	macro.Rewrite("$_.item($1)", index)
	return w
}

// NamedItem fn Retrieves a select object or an object from an options collection.
// js:"namedItem"
func (*HTMLOptionsCollection) NamedItem(name string) (w window.Element) {
	macro.Rewrite("$_.namedItem($1)", name)
	return w
}

// Length prop
// js:"length"
func (*HTMLOptionsCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

// SetLength prop
// js:"length"
func (*HTMLOptionsCollection) SetLength(length uint) {
	macro.Rewrite("$_.length = $1", length)
}

// SelectedIndex prop
// js:"selectedIndex"
func (*HTMLOptionsCollection) SelectedIndex() (selectedIndex int) {
	macro.Rewrite("$_.selectedIndex")
	return selectedIndex
}

// SetSelectedIndex prop
// js:"selectedIndex"
func (*HTMLOptionsCollection) SetSelectedIndex(selectedIndex int) {
	macro.Rewrite("$_.selectedIndex = $1", selectedIndex)
}
