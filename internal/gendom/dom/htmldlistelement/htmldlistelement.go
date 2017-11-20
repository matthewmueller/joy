package htmldlistelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLDListElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLDListElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLDListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}
