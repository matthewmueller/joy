package htmloptgroupelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLOptGroupElement,omit"
type HTMLOptGroupElement struct {
	window.HTMLElement
}

// DefaultSelected Sets or retrieves the status of an option.
func (*HTMLOptGroupElement) DefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.DefaultSelected")
	return defaultSelected
}

// DefaultSelected Sets or retrieves the status of an option.
func (*HTMLOptGroupElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.DefaultSelected = $1", defaultSelected)
}

// Disabled
func (*HTMLOptGroupElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLOptGroupElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLOptGroupElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// Index Sets or retrieves the ordinal position of an option in a list box.
func (*HTMLOptGroupElement) Index() (index int) {
	js.Rewrite("$<.Index")
	return index
}

// Label Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptGroupElement) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// Label Sets or retrieves a value that you can use to implement your own label functionality for the object.
func (*HTMLOptGroupElement) SetLabel(label string) {
	js.Rewrite("$<.Label = $1", label)
}

// Selected Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptGroupElement) Selected() (selected bool) {
	js.Rewrite("$<.Selected")
	return selected
}

// Selected Sets or retrieves whether the option in the list box is the default item.
func (*HTMLOptGroupElement) SetSelected(selected bool) {
	js.Rewrite("$<.Selected = $1", selected)
}

// Text Sets or retrieves the text string specified by the option tag.
func (*HTMLOptGroupElement) Text() (text string) {
	js.Rewrite("$<.Text")
	return text
}

// Value Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptGroupElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLOptGroupElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}
