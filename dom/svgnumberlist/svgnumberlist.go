package svgnumberlist

import (
	"github.com/matthewmueller/golly/dom/svgnumber"
	"github.com/matthewmueller/golly/js"
)

// SVGNumberList struct
// js:"SVGNumberList,omit"
type SVGNumberList struct {
}

// AppendItem fn
func (*SVGNumberList) AppendItem(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

// Clear fn
func (*SVGNumberList) Clear() {
	js.Rewrite("$<.clear()")
}

// GetItem fn
func (*SVGNumberList) GetItem(index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

// Initialize fn
func (*SVGNumberList) Initialize(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
func (*SVGNumberList) InsertItemBefore(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
func (*SVGNumberList) RemoveItem(index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

// ReplaceItem fn
func (*SVGNumberList) ReplaceItem(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
func (*SVGNumberList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
