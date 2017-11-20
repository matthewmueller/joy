package htmlpreelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLPreElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLPreElement) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLPreElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}
