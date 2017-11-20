package htmloptionelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLOptionElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLOptionElement) GetDefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.defaultSelected")
	return defaultSelected
}

func (*HTMLOptionElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.defaultSelected = $1", defaultSelected)
}

func (*HTMLOptionElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLOptionElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLOptionElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLOptionElement) GetIndex() (index int) {
	js.Rewrite("$<.index")
	return index
}

func (*HTMLOptionElement) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*HTMLOptionElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

func (*HTMLOptionElement) GetSelected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

func (*HTMLOptionElement) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

func (*HTMLOptionElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLOptionElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLOptionElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLOptionElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
