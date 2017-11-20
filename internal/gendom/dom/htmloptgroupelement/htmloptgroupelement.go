package htmloptgroupelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLOptGroupElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLOptGroupElement) GetDefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.defaultSelected")
	return defaultSelected
}

func (*HTMLOptGroupElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.defaultSelected = $1", defaultSelected)
}

func (*HTMLOptGroupElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLOptGroupElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLOptGroupElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLOptGroupElement) GetIndex() (index int) {
	js.Rewrite("$<.index")
	return index
}

func (*HTMLOptGroupElement) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*HTMLOptGroupElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

func (*HTMLOptGroupElement) GetSelected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

func (*HTMLOptGroupElement) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

func (*HTMLOptGroupElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLOptGroupElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLOptGroupElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
