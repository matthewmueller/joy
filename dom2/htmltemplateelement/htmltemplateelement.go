package htmltemplateelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTemplateElement struct
// js:"HTMLTemplateElement,omit"
type HTMLTemplateElement struct {
	window.HTMLElement
}

// Content
func (*HTMLTemplateElement) Content() (content *window.DocumentFragment) {
	js.Rewrite("$<.Content")
	return content
}
