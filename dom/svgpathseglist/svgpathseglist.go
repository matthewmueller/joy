package svgpathseglist

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

// SVGPathSegList struct
// js:"SVGPathSegList,omit"
type SVGPathSegList struct {
}

// AppendItem fn
// js:"appendItem"
func (*SVGPathSegList) AppendItem(newItem svgpathseg.SVGPathSeg) (s svgpathseg.SVGPathSeg) {
	macro.Rewrite("$_.appendItem($1)", newItem)
	return s
}

// Clear fn
// js:"clear"
func (*SVGPathSegList) Clear() {
	macro.Rewrite("$_.clear()")
}

// GetItem fn
// js:"getItem"
func (*SVGPathSegList) GetItem(index uint) (s svgpathseg.SVGPathSeg) {
	macro.Rewrite("$_.getItem($1)", index)
	return s
}

// Initialize fn
// js:"initialize"
func (*SVGPathSegList) Initialize(newItem svgpathseg.SVGPathSeg) (s svgpathseg.SVGPathSeg) {
	macro.Rewrite("$_.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
// js:"insertItemBefore"
func (*SVGPathSegList) InsertItemBefore(newItem svgpathseg.SVGPathSeg, index uint) (s svgpathseg.SVGPathSeg) {
	macro.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*SVGPathSegList) RemoveItem(index uint) (s svgpathseg.SVGPathSeg) {
	macro.Rewrite("$_.removeItem($1)", index)
	return s
}

// ReplaceItem fn
// js:"replaceItem"
func (*SVGPathSegList) ReplaceItem(newItem svgpathseg.SVGPathSeg, index uint) (s svgpathseg.SVGPathSeg) {
	macro.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
// js:"numberOfItems"
func (*SVGPathSegList) NumberOfItems() (numberOfItems uint) {
	macro.Rewrite("$_.numberOfItems")
	return numberOfItems
}
