package htmltextareaelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTextAreaElement struct
// js:"HTMLTextAreaElement,omit"
type HTMLTextAreaElement struct {
	window.HTMLElement
}

// CheckValidity Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLTextAreaElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// Select Highlights the input area of a form element.
func (*HTMLTextAreaElement) Select() {
	js.Rewrite("$<.Select()")
}

// SetCustomValidity Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLTextAreaElement) SetCustomValidity(err string) {
	js.Rewrite("$<.SetCustomValidity($1)", err)
}

// SetSelectionRange Sets the start and end positions of a selection in a text field.
//     * @param start The offset into the text field for the start of the selection.
//     * @param end The offset into the text field for the end of the selection.
func (*HTMLTextAreaElement) SetSelectionRange(start uint, end uint) {
	js.Rewrite("$<.SetSelectionRange($1, $2)", start, end)
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLTextAreaElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.Autofocus")
	return autofocus
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLTextAreaElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.Autofocus = $1", autofocus)
}

// Cols Sets or retrieves the width of the object.
func (*HTMLTextAreaElement) Cols() (cols uint) {
	js.Rewrite("$<.Cols")
	return cols
}

// Cols Sets or retrieves the width of the object.
func (*HTMLTextAreaElement) SetCols(cols uint) {
	js.Rewrite("$<.Cols = $1", cols)
}

// DefaultValue Sets or retrieves the initial contents of the object.
func (*HTMLTextAreaElement) DefaultValue() (defaultValue string) {
	js.Rewrite("$<.DefaultValue")
	return defaultValue
}

// DefaultValue Sets or retrieves the initial contents of the object.
func (*HTMLTextAreaElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.DefaultValue = $1", defaultValue)
}

// Disabled
func (*HTMLTextAreaElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLTextAreaElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLTextAreaElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// MaxLength Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLTextAreaElement) MaxLength() (maxLength int) {
	js.Rewrite("$<.MaxLength")
	return maxLength
}

// MaxLength Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLTextAreaElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.MaxLength = $1", maxLength)
}

// Name Sets or retrieves the name of the object.
func (*HTMLTextAreaElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLTextAreaElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Placeholder Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLTextAreaElement) Placeholder() (placeholder string) {
	js.Rewrite("$<.Placeholder")
	return placeholder
}

// Placeholder Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLTextAreaElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.Placeholder = $1", placeholder)
}

// ReadOnly Sets or retrieves the value indicated whether the content of the object is read-only.
func (*HTMLTextAreaElement) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.ReadOnly")
	return readOnly
}

// ReadOnly Sets or retrieves the value indicated whether the content of the object is read-only.
func (*HTMLTextAreaElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.ReadOnly = $1", readOnly)
}

// Required When present, marks an element that can't be submitted without a value.
func (*HTMLTextAreaElement) Required() (required bool) {
	js.Rewrite("$<.Required")
	return required
}

// Required When present, marks an element that can't be submitted without a value.
func (*HTMLTextAreaElement) SetRequired(required bool) {
	js.Rewrite("$<.Required = $1", required)
}

// Rows Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTextAreaElement) Rows() (rows uint) {
	js.Rewrite("$<.Rows")
	return rows
}

// Rows Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTextAreaElement) SetRows(rows uint) {
	js.Rewrite("$<.Rows = $1", rows)
}

// SelectionEnd Gets or sets the end position or offset of a text selection.
func (*HTMLTextAreaElement) SelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.SelectionEnd")
	return selectionEnd
}

// SelectionEnd Gets or sets the end position or offset of a text selection.
func (*HTMLTextAreaElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.SelectionEnd = $1", selectionEnd)
}

// SelectionStart Gets or sets the starting position or offset of a text selection.
func (*HTMLTextAreaElement) SelectionStart() (selectionStart uint) {
	js.Rewrite("$<.SelectionStart")
	return selectionStart
}

// SelectionStart Gets or sets the starting position or offset of a text selection.
func (*HTMLTextAreaElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.SelectionStart = $1", selectionStart)
}

// Status Sets or retrieves the value indicating whether the control is selected.
func (*HTMLTextAreaElement) Status() (status interface{}) {
	js.Rewrite("$<.Status")
	return status
}

// Status Sets or retrieves the value indicating whether the control is selected.
func (*HTMLTextAreaElement) SetStatus(status interface{}) {
	js.Rewrite("$<.Status = $1", status)
}

// Type Retrieves the type of control.
func (*HTMLTextAreaElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// ValidationMessage Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLTextAreaElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.ValidationMessage")
	return validationMessage
}

// Validity Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLTextAreaElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.Validity")
	return validity
}

// Value Retrieves or sets the text in the entry field of the textArea element.
func (*HTMLTextAreaElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value Retrieves or sets the text in the entry field of the textArea element.
func (*HTMLTextAreaElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}

// WillValidate Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLTextAreaElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.WillValidate")
	return willValidate
}

// Wrap Sets or retrieves how to handle wordwrapping in the object.
func (*HTMLTextAreaElement) Wrap() (wrap string) {
	js.Rewrite("$<.Wrap")
	return wrap
}

// Wrap Sets or retrieves how to handle wordwrapping in the object.
func (*HTMLTextAreaElement) SetWrap(wrap string) {
	js.Rewrite("$<.Wrap = $1", wrap)
}
