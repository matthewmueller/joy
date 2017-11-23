package svgpathseglist

import (
	"github.com/matthewmueller/golly/dom2/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

// js:"SVGPathSegList,omit"
type SVGPathSegList struct {
}

// AppendItem
func (*SVGPathSegList) AppendItem(newItem svgpathseg.SVGPathSeg) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.AppendItem($1)", newItem)
	return s
}

// Clear
func (*SVGPathSegList) Clear() {
	js.Rewrite("$<.Clear()")
}

// GetItem
func (*SVGPathSegList) GetItem(index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.GetItem($1)", index)
	return s
}

// Initialize
func (*SVGPathSegList) Initialize(newItem svgpathseg.SVGPathSeg) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.Initialize($1)", newItem)
	return s
}

// InsertItemBefore
func (*SVGPathSegList) InsertItemBefore(newItem svgpathseg.SVGPathSeg, index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.InsertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem
func (*SVGPathSegList) RemoveItem(index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.RemoveItem($1)", index)
	return s
}

// ReplaceItem
func (*SVGPathSegList) ReplaceItem(newItem svgpathseg.SVGPathSeg, index uint) (s svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.ReplaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems
func (*SVGPathSegList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.NumberOfItems")
	return numberOfItems
}
