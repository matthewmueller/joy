package htmlselectelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLSelectElement,omit"
type HTMLSelectElement struct {
	window.HTMLElement
}

// Add Adds an element to the areas, controlRange, or options collection.
//     * @param element Variant of type Number that specifies the index position in the collection where the element is placed. If no value is given, the method places the element at the end of the collection.
//     * @param before Variant of type Object that specifies an element to insert before, or null to append the object to the collection.
func (*HTMLSelectElement) Add(element window.HTMLElement, before *interface{}) {
	js.Rewrite("$<.Add($1, $2)", element, before)
}

// CheckValidity Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLSelectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// Item Retrieves a select object or an object from an options collection.
//     * @param name Variant of type Number or String that specifies the object or collection to retrieve. If this parameter is an integer, it is the zero-based index of the object. If this parameter is a string, all objects with matching name or id properties are retrieved, and a collection is returned if more than one match is made.
//     * @param index Variant of type Number that specifies the zero-based index of the object to retrieve when a collection is returned.
func (*HTMLSelectElement) Item(name *interface{}, index *interface{}) (i interface{}) {
	js.Rewrite("$<.Item($1, $2)", name, index)
	return i
}

// NamedItem Retrieves a select object or an object from an options collection.
//     * @param namedItem A String that specifies the name or id property of the object to retrieve. A collection is returned if more than one match is made.
func (*HTMLSelectElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.NamedItem($1)", name)
	return i
}

// Remove Removes an element from the collection.
//     * @param index Number that specifies the zero-based index of the element to remove from the collection.
func (*HTMLSelectElement) Remove(index *int) {
	js.Rewrite("$<.Remove($1)", index)
}

// SetCustomValidity Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLSelectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.SetCustomValidity($1)", err)
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLSelectElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.Autofocus")
	return autofocus
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLSelectElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.Autofocus = $1", autofocus)
}

// Disabled
func (*HTMLSelectElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLSelectElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLSelectElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// Length Sets or retrieves the number of objects in a collection.
func (*HTMLSelectElement) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Length Sets or retrieves the number of objects in a collection.
func (*HTMLSelectElement) SetLength(length uint) {
	js.Rewrite("$<.Length = $1", length)
}

// Multiple Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLSelectElement) Multiple() (multiple bool) {
	js.Rewrite("$<.Multiple")
	return multiple
}

// Multiple Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLSelectElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.Multiple = $1", multiple)
}

// Name Sets or retrieves the name of the object.
func (*HTMLSelectElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLSelectElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Options
func (*HTMLSelectElement) Options() (options *htmloptionscollection.HTMLOptionsCollection) {
	js.Rewrite("$<.Options")
	return options
}

// Required When present, marks an element that can't be submitted without a value.
func (*HTMLSelectElement) Required() (required bool) {
	js.Rewrite("$<.Required")
	return required
}

// Required When present, marks an element that can't be submitted without a value.
func (*HTMLSelectElement) SetRequired(required bool) {
	js.Rewrite("$<.Required = $1", required)
}

// SelectedIndex Sets or retrieves the index of the selected option in a select object.
func (*HTMLSelectElement) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.SelectedIndex")
	return selectedIndex
}

// SelectedIndex Sets or retrieves the index of the selected option in a select object.
func (*HTMLSelectElement) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.SelectedIndex = $1", selectedIndex)
}

// SelectedOptions
func (*HTMLSelectElement) SelectedOptions() (selectedOptions window.HTMLCollection) {
	js.Rewrite("$<.SelectedOptions")
	return selectedOptions
}

// Size Sets or retrieves the number of rows in the list box.
func (*HTMLSelectElement) Size() (size uint) {
	js.Rewrite("$<.Size")
	return size
}

// Size Sets or retrieves the number of rows in the list box.
func (*HTMLSelectElement) SetSize(size uint) {
	js.Rewrite("$<.Size = $1", size)
}

// Type Retrieves the type of select control based on the value of the MULTIPLE attribute.
func (*HTMLSelectElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// ValidationMessage Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLSelectElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.ValidationMessage")
	return validationMessage
}

// Validity Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLSelectElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.Validity")
	return validity
}

// Value Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLSelectElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value Sets or retrieves the value which is returned to the server when the form control is submitted.
func (*HTMLSelectElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}

// WillValidate Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLSelectElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.WillValidate")
	return willValidate
}
