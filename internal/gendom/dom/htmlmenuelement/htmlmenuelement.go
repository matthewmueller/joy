package htmlmenuelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLMenuElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLMenuElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLMenuElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

func (*HTMLMenuElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLMenuElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
