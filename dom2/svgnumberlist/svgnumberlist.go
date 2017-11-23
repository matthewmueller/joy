package svgnumberlist

import (
	"github.com/matthewmueller/golly/dom2/svgnumber"
	"github.com/matthewmueller/golly/js"
)

// SVGNumberList struct
// js:"SVGNumberList,omit"
type SVGNumberList struct {
}

// AppendItem
func (*SVGNumberList) AppendItem(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.AppendItem($1)", newItem)
	return s
}

// Clear
func (*SVGNumberList) Clear() {
	js.Rewrite("$<.Clear()")
}

// GetItem
func (*SVGNumberList) GetItem(index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.GetItem($1)", index)
	return s
}

// Initialize
func (*SVGNumberList) Initialize(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.Initialize($1)", newItem)
	return s
}

// InsertItemBefore
func (*SVGNumberList) InsertItemBefore(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.InsertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem
func (*SVGNumberList) RemoveItem(index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.RemoveItem($1)", index)
	return s
}

// ReplaceItem
func (*SVGNumberList) ReplaceItem(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.ReplaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems
func (*SVGNumberList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.NumberOfItems")
	return numberOfItems
}
