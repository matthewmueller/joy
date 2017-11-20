package htmloptionscollection

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/js"
)

type HTMLOptionsCollection struct {
	*htmlcollection.HTMLCollection
}

func (*HTMLOptionsCollection) Add(element interface{}, before interface{}) {
	js.Rewrite("$<.add($1, $2)", element, before)
}

func (*HTMLOptionsCollection) Remove(index int) {
	js.Rewrite("$<.remove($1)", index)
}

func (*HTMLOptionsCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*HTMLOptionsCollection) SetLength(length uint) {
	js.Rewrite("$<.length = $1", length)
}

func (*HTMLOptionsCollection) GetSelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}

func (*HTMLOptionsCollection) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.selectedIndex = $1", selectedIndex)
}
