package svgpointlist

import (
	"github.com/matthewmueller/golly/dom2/svgpoint"
	"github.com/matthewmueller/golly/js"
)

// SVGPointList struct
// js:"SVGPointList,omit"
type SVGPointList struct {
}

// AppendItem fn
func (*SVGPointList) AppendItem(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

// Clear fn
func (*SVGPointList) Clear() {
	js.Rewrite("$<.clear()")
}

// GetItem fn
func (*SVGPointList) GetItem(index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

// Initialize fn
func (*SVGPointList) Initialize(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
func (*SVGPointList) InsertItemBefore(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
func (*SVGPointList) RemoveItem(index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

// ReplaceItem fn
func (*SVGPointList) ReplaceItem(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
func (*SVGPointList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
