package svgstringlist

import "github.com/matthewmueller/golly/js"

// js:"SVGStringList,omit"
type SVGStringList struct {
}

// AppendItem
func (*SVGStringList) AppendItem(newItem string) (s string) {
	js.Rewrite("$<.AppendItem($1)", newItem)
	return s
}

// Clear
func (*SVGStringList) Clear() {
	js.Rewrite("$<.Clear()")
}

// GetItem
func (*SVGStringList) GetItem(index uint) (s string) {
	js.Rewrite("$<.GetItem($1)", index)
	return s
}

// Initialize
func (*SVGStringList) Initialize(newItem string) (s string) {
	js.Rewrite("$<.Initialize($1)", newItem)
	return s
}

// InsertItemBefore
func (*SVGStringList) InsertItemBefore(newItem string, index uint) (s string) {
	js.Rewrite("$<.InsertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem
func (*SVGStringList) RemoveItem(index uint) (s string) {
	js.Rewrite("$<.RemoveItem($1)", index)
	return s
}

// ReplaceItem
func (*SVGStringList) ReplaceItem(newItem string, index uint) (s string) {
	js.Rewrite("$<.ReplaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems
func (*SVGStringList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.NumberOfItems")
	return numberOfItems
}
