package htmlprogresselement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLProgressElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLProgressElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLProgressElement) GetMax() (max float32) {
	js.Rewrite("$<.max")
	return max
}

func (*HTMLProgressElement) SetMax(max float32) {
	js.Rewrite("$<.max = $1", max)
}

func (*HTMLProgressElement) GetPosition() (position float32) {
	js.Rewrite("$<.position")
	return position
}

func (*HTMLProgressElement) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLProgressElement) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
