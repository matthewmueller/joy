package htmllielement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLLIElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLLIElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLLIElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLLIElement) GetValue() (value int) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLLIElement) SetValue(value int) {
	js.Rewrite("$<.value = $1", value)
}
