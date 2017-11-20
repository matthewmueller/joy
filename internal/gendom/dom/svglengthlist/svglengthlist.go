package svglengthlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svglength"
	"github.com/matthewmueller/golly/js"
)

type SVGLengthList struct {
}

func (*SVGLengthList) AppendItem(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGLengthList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGLengthList) GetItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGLengthList) Initialize(newItem *svglength.SVGLength) (s *svglength.SVGLength) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGLengthList) InsertItemBefore(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGLengthList) RemoveItem(index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGLengthList) ReplaceItem(newItem *svglength.SVGLength, index uint) (s *svglength.SVGLength) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGLengthList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
