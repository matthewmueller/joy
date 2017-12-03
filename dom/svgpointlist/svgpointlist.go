package svgpointlist

import (
	"github.com/matthewmueller/joy/dom/svgpoint"
	"github.com/matthewmueller/joy/macro"
)

// SVGPointList struct
// js:"SVGPointList,omit"
type SVGPointList struct {
}

// AppendItem fn
// js:"appendItem"
func (*SVGPointList) AppendItem(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

// Clear fn
// js:"clear"
func (*SVGPointList) Clear() {
	macro.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*SVGPointList) GetItem(index uint) (s *svgpoint.SVGPoint) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

// Initialize fn
// js:"initialize"
func (*SVGPointList) Initialize(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
// js:"insertItemBefore"
func (*SVGPointList) InsertItemBefore(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*SVGPointList) RemoveItem(index uint) (s *svgpoint.SVGPoint) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

// ReplaceItem fn
// js:"replaceItem"
func (*SVGPointList) ReplaceItem(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
// js:"numberOfItems"
func (*SVGPointList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}
