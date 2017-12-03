package svglengthlist

import (
	"github.com/matthewmueller/joy/dom/svglength"
	"github.com/matthewmueller/joy/js"
)

// SVGLengthList struct
// js:"SVGLengthList,omit"
type SVGLengthList struct {
}

// AppendItem fn
// js:"appendItem"
func (*SVGLengthList) AppendItem(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$_.appendItem($1)", newItem)
	return s
}

// Clear fn
// js:"clear"
func (*SVGLengthList) Clear() {
	js.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*SVGLengthList) GetItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$_.getItem($1)", index)
	return s
}

// Initialize fn
// js:"initialize"
func (*SVGLengthList) Initialize(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$_.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
// js:"insertItemBefore"
func (*SVGLengthList) InsertItemBefore(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*SVGLengthList) RemoveItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$_.removeItem($1)", index)
	return s
}

// ReplaceItem fn
// js:"replaceItem"
func (*SVGLengthList) ReplaceItem(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
// js:"numberOfItems"
func (*SVGLengthList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$_.numberOfItems")
	return numberOfItems
}
