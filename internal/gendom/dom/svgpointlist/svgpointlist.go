package svgpointlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgpoint"
	"github.com/matthewmueller/golly/js"
)

type SVGPointList struct {
}

func (*SVGPointList) AppendItem(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGPointList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGPointList) GetItem(index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGPointList) Initialize(newItem *svgpoint.SVGPoint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGPointList) InsertItemBefore(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGPointList) RemoveItem(index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGPointList) ReplaceItem(newItem *svgpoint.SVGPoint, index uint) (s *svgpoint.SVGPoint) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGPointList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
