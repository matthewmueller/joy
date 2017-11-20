package htmldataelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLDataElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLDataElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLDataElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
