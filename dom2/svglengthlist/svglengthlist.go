package svglengthlist

import (
	"github.com/matthewmueller/golly/dom2/svglength"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGLengthList,omit"
type SVGLengthList struct {
}

// AppendItem
func (*SVGLengthList) AppendItem(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$<.AppendItem($1)", newItem)
	return s
}

// Clear
func (*SVGLengthList) Clear() {
	js.Rewrite("$<.Clear()")
}

// GetItem
func (*SVGLengthList) GetItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.GetItem($1)", index)
	return s
}

// Initialize
func (*SVGLengthList) Initialize(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$<.Initialize($1)", newItem)
	return s
}

// InsertItemBefore
func (*SVGLengthList) InsertItemBefore(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.InsertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem
func (*SVGLengthList) RemoveItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.RemoveItem($1)", index)
	return s
}

// ReplaceItem
func (*SVGLengthList) ReplaceItem(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.ReplaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems
func (*SVGLengthList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.NumberOfItems")
	return numberOfItems
}
