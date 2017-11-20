package svgpathseglist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

type SVGPathSegList struct {
}

func (*SVGPathSegList) AppendItem(newItem *svgpathseg.SVGPathSeg) (s *svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGPathSegList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGPathSegList) GetItem(index uint) (s *svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGPathSegList) Initialize(newItem *svgpathseg.SVGPathSeg) (s *svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGPathSegList) InsertItemBefore(newItem *svgpathseg.SVGPathSeg, index uint) (s *svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGPathSegList) RemoveItem(index uint) (s *svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGPathSegList) ReplaceItem(newItem *svgpathseg.SVGPathSeg, index uint) (s *svgpathseg.SVGPathSeg) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGPathSegList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
