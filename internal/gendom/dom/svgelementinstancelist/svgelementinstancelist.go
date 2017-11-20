package svgelementinstancelist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/svgelementinstance"
	"github.com/matthewmueller/golly/js"
)

type SVGElementInstanceList struct {
}

func (*SVGElementInstanceList) Item(index uint) (s *svgelementinstance.SVGElementInstance) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*SVGElementInstanceList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
