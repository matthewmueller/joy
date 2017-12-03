package htmlareascollection

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.HTMLCollection = (*HTMLAreasCollection)(nil)

// HTMLAreasCollection struct
// js:"HTMLAreasCollection,omit"
type HTMLAreasCollection struct {
}

// Item fn Retrieves an object from various collections.
// js:"item"
func (*HTMLAreasCollection) Item(index uint) (w window.Element) {
	macro.Rewrite("$_.item($1)", index)
	return w
}

// NamedItem fn Retrieves a select object or an object from an options collection.
// js:"namedItem"
func (*HTMLAreasCollection) NamedItem(name string) (w window.Element) {
	macro.Rewrite("$_.namedItem($1)", name)
	return w
}

// Length prop Sets or retrieves the number of objects in a collection.
// js:"length"
func (*HTMLAreasCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
