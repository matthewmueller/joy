package htmlquoteelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLQuoteElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLQuoteElement) GetCite() (cite string) {
	js.Rewrite("$<.cite")
	return cite
}

func (*HTMLQuoteElement) SetCite(cite string) {
	js.Rewrite("$<.cite = $1", cite)
}
