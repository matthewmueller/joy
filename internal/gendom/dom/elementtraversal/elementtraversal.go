package elementtraversal

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/js"
)

type ElementTraversal struct {
}

func (*ElementTraversal) GetChildElementCount() (childElementCount uint) {
	js.Rewrite("$<.childElementCount")
	return childElementCount
}

func (*ElementTraversal) GetFirstElementChild() (firstElementChild *element.Element) {
	js.Rewrite("$<.firstElementChild")
	return firstElementChild
}

func (*ElementTraversal) GetLastElementChild() (lastElementChild *element.Element) {
	js.Rewrite("$<.lastElementChild")
	return lastElementChild
}

func (*ElementTraversal) GetNextElementSibling() (nextElementSibling *element.Element) {
	js.Rewrite("$<.nextElementSibling")
	return nextElementSibling
}

func (*ElementTraversal) GetPreviousElementSibling() (previousElementSibling *element.Element) {
	js.Rewrite("$<.previousElementSibling")
	return previousElementSibling
}
