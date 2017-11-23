package svgstringlist

import "github.com/matthewmueller/golly/js"

// SVGStringList struct
// js:"SVGStringList,omit"
type SVGStringList struct {
}

// AppendItem fn
func (*SVGStringList) AppendItem(newItem string) (s string) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

// Clear fn
func (*SVGStringList) Clear() {
	js.Rewrite("$<.clear()")
}

// GetItem fn
func (*SVGStringList) GetItem(index uint) (s string) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

// Initialize fn
func (*SVGStringList) Initialize(newItem string) (s string) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
func (*SVGStringList) InsertItemBefore(newItem string, index uint) (s string) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
func (*SVGStringList) RemoveItem(index uint) (s string) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

// ReplaceItem fn
func (*SVGStringList) ReplaceItem(newItem string, index uint) (s string) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
func (*SVGStringList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
