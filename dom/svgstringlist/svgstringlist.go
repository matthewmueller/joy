package svgstringlist

import "github.com/matthewmueller/golly/js"

// SVGStringList struct
// js:"SVGStringList,omit"
type SVGStringList struct {
}

// AppendItem fn
// js:"appendItem"
func (*SVGStringList) AppendItem(newItem string) (s string) {
	js.Rewrite("$_.appendItem($1)", newItem)
	return s
}

// Clear fn
// js:"clear"
func (*SVGStringList) Clear() {
	js.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*SVGStringList) GetItem(index uint) (s string) {
	js.Rewrite("$_.getItem($1)", index)
	return s
}

// Initialize fn
// js:"initialize"
func (*SVGStringList) Initialize(newItem string) (s string) {
	js.Rewrite("$_.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
// js:"insertItemBefore"
func (*SVGStringList) InsertItemBefore(newItem string, index uint) (s string) {
	js.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*SVGStringList) RemoveItem(index uint) (s string) {
	js.Rewrite("$_.removeItem($1)", index)
	return s
}

// ReplaceItem fn
// js:"replaceItem"
func (*SVGStringList) ReplaceItem(newItem string, index uint) (s string) {
	js.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
// js:"numberOfItems"
func (*SVGStringList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$_.numberOfItems")
	return numberOfItems
}
