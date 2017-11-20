package svgnumberlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgnumber"
	"github.com/matthewmueller/golly/js"
)

type SVGNumberList struct {
}

func (*SVGNumberList) AppendItem(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGNumberList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGNumberList) GetItem(index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGNumberList) Initialize(newItem *svgnumber.SVGNumber) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGNumberList) InsertItemBefore(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGNumberList) RemoveItem(index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGNumberList) ReplaceItem(newItem *svgnumber.SVGNumber, index uint) (s *svgnumber.SVGNumber) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGNumberList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
