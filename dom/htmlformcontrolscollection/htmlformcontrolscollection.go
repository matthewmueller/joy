package htmlformcontrolscollection

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.HTMLCollection = (*HTMLFormControlsCollection)(nil)

// HTMLFormControlsCollection struct
// js:"HTMLFormControlsCollection,omit"
type HTMLFormControlsCollection struct {
}

// NamedItem fn
// js:"namedItem"
func (*HTMLFormControlsCollection) NamedItem(name string) (w window.Element) {
	macro.Rewrite("$_.namedItem($1)", name)
	return w
}

// Item fn Retrieves an object from various collections.
// js:"item"
func (*HTMLFormControlsCollection) Item(index uint) (w window.Element) {
	macro.Rewrite("$_.item($1)", index)
	return w
}

// Length prop Sets or retrieves the number of objects in a collection.
// js:"length"
func (*HTMLFormControlsCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
