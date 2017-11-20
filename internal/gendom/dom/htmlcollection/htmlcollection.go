package htmlcollection

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/js"
)

type HTMLCollection struct {
}

func (*HTMLCollection) Item(index uint) (e *element.Element) {
	js.Rewrite("$<.item($1)", index)
	return e
}

func (*HTMLCollection) NamedItem(name string) (e *element.Element) {
	js.Rewrite("$<.namedItem($1)", name)
	return e
}

func (*HTMLCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
