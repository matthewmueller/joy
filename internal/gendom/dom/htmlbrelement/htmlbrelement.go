package htmlbrelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLBRElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLBRElement) GetClear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

func (*HTMLBRElement) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}
