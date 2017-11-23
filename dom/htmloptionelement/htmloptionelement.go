package htmloptionelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLOptionElement struct
// js:"HTMLOptionElement,omit"
type HTMLOptionElement struct {
	window.HTMLElement
}

// DefaultSelected prop Sets or retrieves the status of an option.
func (*HTMLOptionElement) DefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.defaultSelected")
	return defaultSelected
}

// DefaultSelected prop Sets or retrieves the status of an option.
func (*HTMLOptionElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.defaultSelected = $1", defaultSelected)
}

// Disabled prop
func (*HTMLOptionElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLOptionElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLOptionElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// Index prop Sets or retrieves the ordinal position of an option in a list box.
func (*HTMLOptionElement) Index() (index int) {
	js.Rewrite("$<.index")
	return index
}

// Label prop Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptionElement) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// Label prop Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptionElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

// Selected prop Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptionElement) Selected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

// Selected prop Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptionElement) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

// Text prop Sets or retrieves the text string specified by the option tag.
func (*HTMLOptionElement) Text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// Text prop Sets or retrieves the text string specified by the option tag.
func (*HTMLOptionElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

// Value prop Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptionElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptionElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
