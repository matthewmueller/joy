package svgnumberlist

import (
	"github.com/matthewmueller/joy/dom/svgnumber"
	"github.com/matthewmueller/joy/macro"
)

// SVGNumberList struct
// js:"SVGNumberList,omit"
type SVGNumberList struct {
}

// AppendItem fn
// js:"appendItem"
func (*SVGNumberList) AppendItem(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

// Clear fn
// js:"clear"
func (*SVGNumberList) Clear() {
	macro.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*SVGNumberList) GetItem(index uint) (s *svgnumber.SVGNumber) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

// Initialize fn
// js:"initialize"
func (*SVGNumberList) Initialize(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
// js:"insertItemBefore"
func (*SVGNumberList) InsertItemBefore(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*SVGNumberList) RemoveItem(index uint) (s *svgnumber.SVGNumber) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

// ReplaceItem fn
// js:"replaceItem"
func (*SVGNumberList) ReplaceItem(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
// js:"numberOfItems"
func (*SVGNumberList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}
