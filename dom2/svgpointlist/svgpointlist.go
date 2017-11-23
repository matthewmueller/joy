package svgpointlist

import (
	"github.com/matthewmueller/golly/dom2/svgpoint"
	"github.com/matthewmueller/golly/js"
)

// SVGPointList struct
// js:"SVGPointList,omit"
type SVGPointList struct {
}

// AppendItem
func (*SVGPointList) AppendItem(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.AppendItem($1)", newItem)
	return s
}

// Clear
func (*SVGPointList) Clear() {
	js.Rewrite("$<.Clear()")
}

// GetItem
func (*SVGPointList) GetItem(index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.GetItem($1)", index)
	return s
}

// Initialize
func (*SVGPointList) Initialize(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.Initialize($1)", newItem)
	return s
}

// InsertItemBefore
func (*SVGPointList) InsertItemBefore(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.InsertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem
func (*SVGPointList) RemoveItem(index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.RemoveItem($1)", index)
	return s
}

// ReplaceItem
func (*SVGPointList) ReplaceItem(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.ReplaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems
func (*SVGPointList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.NumberOfItems")
	return numberOfItems
}
