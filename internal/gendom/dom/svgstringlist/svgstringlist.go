package svgstringlist

import "github.com/matthewmueller/golly/js"

type SVGStringList struct {
}

func (*SVGStringList) AppendItem(newItem string) (s string) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGStringList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGStringList) GetItem(index uint) (s string) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGStringList) Initialize(newItem string) (s string) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGStringList) InsertItemBefore(newItem string, index uint) (s string) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGStringList) RemoveItem(index uint) (s string) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGStringList) ReplaceItem(newItem string, index uint) (s string) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGStringList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
