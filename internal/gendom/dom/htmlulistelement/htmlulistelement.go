package htmlulistelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLUListElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLUListElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLUListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

func (*HTMLUListElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLUListElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
