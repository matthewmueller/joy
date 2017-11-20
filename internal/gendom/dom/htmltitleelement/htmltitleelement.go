package htmltitleelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLTitleElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLTitleElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLTitleElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}
