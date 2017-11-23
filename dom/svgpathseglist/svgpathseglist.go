package svgpathseglist

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// SVGPathSegList struct
// js:"SVGPathSegList,omit"
type SVGPathSegList struct {
}

// AppendItem fn
func (*SVGPathSegList) AppendItem(newItem svgpathseg.SVGPathSeg) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

// Clear fn
func (*SVGPathSegList) Clear() {
	js.Rewrite("$<.clear()")
}

// GetItem fn
func (*SVGPathSegList) GetItem(index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

// Initialize fn
func (*SVGPathSegList) Initialize(newItem svgpathseg.SVGPathSeg) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
func (*SVGPathSegList) InsertItemBefore(newItem svgpathseg.SVGPathSeg, index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
func (*SVGPathSegList) RemoveItem(index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

// ReplaceItem fn
func (*SVGPathSegList) ReplaceItem(newItem svgpathseg.SVGPathSeg, index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
func (*SVGPathSegList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
