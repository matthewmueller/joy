package htmloptgroupelement

import (
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLOptGroupElement struct
// js:"HTMLOptGroupElement,omit"
type HTMLOptGroupElement struct {
	window.HTMLElement
}

// DefaultSelected prop Sets or retrieves the status of an option.
func (*HTMLOptGroupElement) DefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.defaultSelected")
	return defaultSelected
}

// DefaultSelected prop Sets or retrieves the status of an option.
func (*HTMLOptGroupElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.defaultSelected = $1", defaultSelected)
}

// Disabled prop
func (*HTMLOptGroupElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLOptGroupElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLOptGroupElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// Index prop Sets or retrieves the ordinal position of an option in a list box.
func (*HTMLOptGroupElement) Index() (index int) {
	js.Rewrite("$<.index")
	return index
}

// Label prop Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptGroupElement) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// Label prop Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptGroupElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

// Selected prop Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptGroupElement) Selected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

// Selected prop Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptGroupElement) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

// Text prop Sets or retrieves the text string specified by the option tag.
func (*HTMLOptGroupElement) Text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// Value prop Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptGroupElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptGroupElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
