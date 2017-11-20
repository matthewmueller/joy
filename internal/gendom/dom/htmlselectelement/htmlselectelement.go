package htmlselectelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmloptionscollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/validitystate"
	"github.com/matthewmueller/golly/js"
)

type HTMLSelectElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLSelectElement) Add(element *htmlelement.HTMLElement, before interface{}) {
	js.Rewrite("$<.add($1, $2)", element, before)
}

func (*HTMLSelectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLSelectElement) Item(name interface{}, index interface{}) (i interface{}) {
	js.Rewrite("$<.item($1, $2)", name, index)
	return i
}

func (*HTMLSelectElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

func (*HTMLSelectElement) Remove(index int) {
	js.Rewrite("$<.remove($1)", index)
}

func (*HTMLSelectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLSelectElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLSelectElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLSelectElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLSelectElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLSelectElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLSelectElement) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*HTMLSelectElement) SetLength(length uint) {
	js.Rewrite("$<.length = $1", length)
}

func (*HTMLSelectElement) GetMultiple() (multiple bool) {
	js.Rewrite("$<.multiple")
	return multiple
}

func (*HTMLSelectElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.multiple = $1", multiple)
}

func (*HTMLSelectElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLSelectElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLSelectElement) GetOptions() (options *htmloptionscollection.HTMLOptionsCollection) {
	js.Rewrite("$<.options")
	return options
}

func (*HTMLSelectElement) GetRequired() (required bool) {
	js.Rewrite("$<.required")
	return required
}

func (*HTMLSelectElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

func (*HTMLSelectElement) GetSelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}

func (*HTMLSelectElement) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.selectedIndex = $1", selectedIndex)
}

func (*HTMLSelectElement) GetSelectedOptions() (selectedOptions *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.selectedOptions")
	return selectedOptions
}

func (*HTMLSelectElement) GetSize() (size uint) {
	js.Rewrite("$<.size")
	return size
}

func (*HTMLSelectElement) SetSize(size uint) {
	js.Rewrite("$<.size = $1", size)
}

func (*HTMLSelectElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLSelectElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLSelectElement) GetValidity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLSelectElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLSelectElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLSelectElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
