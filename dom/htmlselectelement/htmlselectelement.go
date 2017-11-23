package htmlselectelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/htmloptionscollection"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLSelectElement struct
// js:"HTMLSelectElement,omit"
type HTMLSelectElement struct {
	window.HTMLElement
}

// Add fn Adds an element to the areas, controlRange, or options collection.
//     * @param element Variant of type Number that specifies the index position in the collection where the element is placed. If no value is given, the method places the element at the end of the collection.
//     * @param before Variant of type Object that specifies an element to insert before, or null to append the object to the collection.
func (*HTMLSelectElement) Add(element window.HTMLElement, before *interface{}) {
	js.Rewrite("$<.add($1, $2)", element, before)
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLSelectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// Item fn Retrieves a select object or an object from an options collection.
//     * @param name Variant of type Number or String that specifies the object or collection to retrieve. If this parameter is an integer, it is the zero-based index of the object. If this parameter is a string, all objects with matching name or id properties are retrieved, and a collection is returned if more than one match is made.
//     * @param index Variant of type Number that specifies the zero-based index of the object to retrieve when a collection is returned.
func (*HTMLSelectElement) Item(name *interface{}, index *interface{}) (i interface{}) {
	js.Rewrite("$<.item($1, $2)", name, index)
	return i
}

// NamedItem fn Retrieves a select object or an object from an options collection.
//     * @param namedItem A String that specifies the name or id property of the object to retrieve. A collection is returned if more than one match is made.
func (*HTMLSelectElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

// Remove fn Removes an element from the collection.
//     * @param index Number that specifies the zero-based index of the element to remove from the collection.
func (*HTMLSelectElement) Remove(index *int) {
	js.Rewrite("$<.remove($1)", index)
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLSelectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLSelectElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLSelectElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

// Disabled prop
func (*HTMLSelectElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLSelectElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLSelectElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// Length prop Sets or retrieves the number of objects in a collection.
func (*HTMLSelectElement) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Length prop Sets or retrieves the number of objects in a collection.
func (*HTMLSelectElement) SetLength(length uint) {
	js.Rewrite("$<.length = $1", length)
}

// Multiple prop Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLSelectElement) Multiple() (multiple bool) {
	js.Rewrite("$<.multiple")
	return multiple
}

// Multiple prop Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLSelectElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.multiple = $1", multiple)
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLSelectElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLSelectElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Options prop
func (*HTMLSelectElement) Options() (options *htmloptionscollection.HTMLOptionsCollection) {
	js.Rewrite("$<.options")
	return options
}

// Required prop When present, marks an element that can't be submitted without a value.
func (*HTMLSelectElement) Required() (required bool) {
	js.Rewrite("$<.required")
	return required
}

// Required prop When present, marks an element that can't be submitted without a value.
func (*HTMLSelectElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

// SelectedIndex prop Sets or retrieves the index of the selected option in a select object.
func (*HTMLSelectElement) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}

// SelectedIndex prop Sets or retrieves the index of the selected option in a select object.
func (*HTMLSelectElement) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.selectedIndex = $1", selectedIndex)
}

// SelectedOptions prop
func (*HTMLSelectElement) SelectedOptions() (selectedOptions window.HTMLCollection) {
	js.Rewrite("$<.selectedOptions")
	return selectedOptions
}

// Size prop Sets or retrieves the number of rows in the list box.
func (*HTMLSelectElement) Size() (size uint) {
	js.Rewrite("$<.size")
	return size
}

// Size prop Sets or retrieves the number of rows in the list box.
func (*HTMLSelectElement) SetSize(size uint) {
	js.Rewrite("$<.size = $1", size)
}

// Type prop Retrieves the type of select control based on the value of the MULTIPLE attribute.
func (*HTMLSelectElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLSelectElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLSelectElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

// Value prop Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLSelectElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLSelectElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLSelectElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
