package htmltemplateelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/documentfragment"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLTemplateElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLTemplateElement) GetContent() (content *documentfragment.DocumentFragment) {
	js.Rewrite("$<.content")
	return content
}
