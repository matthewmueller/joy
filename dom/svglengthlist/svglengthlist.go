package svglengthlist

import (
	"github.com/matthewmueller/golly/dom/svglength"
	"github.com/matthewmueller/golly/js"
)

// SVGLengthList struct
// js:"SVGLengthList,omit"
type SVGLengthList struct {
}

// AppendItem fn
func (*SVGLengthList) AppendItem(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

// Clear fn
func (*SVGLengthList) Clear() {
	js.Rewrite("$<.clear()")
}

// GetItem fn
func (*SVGLengthList) GetItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

// Initialize fn
func (*SVGLengthList) Initialize(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
func (*SVGLengthList) InsertItemBefore(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
func (*SVGLengthList) RemoveItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

// ReplaceItem fn
func (*SVGLengthList) ReplaceItem(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
func (*SVGLengthList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
