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

// DefaultSelected Sets or retrieves the status of an option.
func (*HTMLOptionElement) DefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.DefaultSelected")
	return defaultSelected
}

// DefaultSelected Sets or retrieves the status of an option.
func (*HTMLOptionElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.DefaultSelected = $1", defaultSelected)
}

// Disabled
func (*HTMLOptionElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLOptionElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLOptionElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// Index Sets or retrieves the ordinal position of an option in a list box.
func (*HTMLOptionElement) Index() (index int) {
	js.Rewrite("$<.Index")
	return index
}

// Label Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptionElement) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// Label Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptionElement) SetLabel(label string) {
	js.Rewrite("$<.Label = $1", label)
}

// Selected Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptionElement) Selected() (selected bool) {
	js.Rewrite("$<.Selected")
	return selected
}

// Selected Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptionElement) SetSelected(selected bool) {
	js.Rewrite("$<.Selected = $1", selected)
}

// Text Sets or retrieves the text string specified by the option tag.
func (*HTMLOptionElement) Text() (text string) {
	js.Rewrite("$<.Text")
	return text
}

// Text Sets or retrieves the text string specified by the option tag.
func (*HTMLOptionElement) SetText(text string) {
	js.Rewrite("$<.Text = $1", text)
}

// Value Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptionElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptionElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}
