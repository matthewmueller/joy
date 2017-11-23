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

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLTextAreaElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// Select fn Highlights the input area of a form element.
func (*HTMLTextAreaElement) Select() {
	js.Rewrite("$<.select()")
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLTextAreaElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

// SetSelectionRange fn Sets the start and end positions of a selection in a text field.
//     * @param start The offset into the text field for the start of the selection.
//     * @param end The offset into the text field for the end of the selection.
func (*HTMLTextAreaElement) SetSelectionRange(start uint, end uint) {
	js.Rewrite("$<.setSelectionRange($1, $2)", start, end)
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLTextAreaElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLTextAreaElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

// Cols prop Sets or retrieves the width of the object.
func (*HTMLTextAreaElement) Cols() (cols uint) {
	js.Rewrite("$<.cols")
	return cols
}

// Cols prop Sets or retrieves the width of the object.
func (*HTMLTextAreaElement) SetCols(cols uint) {
	js.Rewrite("$<.cols = $1", cols)
}

// DefaultValue prop Sets or retrieves the initial contents of the object.
func (*HTMLTextAreaElement) DefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

// DefaultValue prop Sets or retrieves the initial contents of the object.
func (*HTMLTextAreaElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

// Disabled prop
func (*HTMLTextAreaElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLTextAreaElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLTextAreaElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// MaxLength prop Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLTextAreaElement) MaxLength() (maxLength int) {
	js.Rewrite("$<.maxLength")
	return maxLength
}

// MaxLength prop Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLTextAreaElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.maxLength = $1", maxLength)
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLTextAreaElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLTextAreaElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Placeholder prop Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLTextAreaElement) Placeholder() (placeholder string) {
	js.Rewrite("$<.placeholder")
	return placeholder
}

// Placeholder prop Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLTextAreaElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.placeholder = $1", placeholder)
}

// ReadOnly prop Sets or retrieves the value indicated whether the content of the object is read-only.
func (*HTMLTextAreaElement) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

// ReadOnly prop Sets or retrieves the value indicated whether the content of the object is read-only.
func (*HTMLTextAreaElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.readOnly = $1", readOnly)
}

// Required prop When present, marks an element that can't be submitted without a value.
func (*HTMLTextAreaElement) Required() (required bool) {
	js.Rewrite("$<.required")
	return required
}

// Required prop When present, marks an element that can't be submitted without a value.
func (*HTMLTextAreaElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

// Rows prop Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTextAreaElement) Rows() (rows uint) {
	js.Rewrite("$<.rows")
	return rows
}

// Rows prop Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTextAreaElement) SetRows(rows uint) {
	js.Rewrite("$<.rows = $1", rows)
}

// SelectionEnd prop Gets or sets the end position or offset of a text selection.
func (*HTMLTextAreaElement) SelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.selectionEnd")
	return selectionEnd
}

// SelectionEnd prop Gets or sets the end position or offset of a text selection.
func (*HTMLTextAreaElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.selectionEnd = $1", selectionEnd)
}

// SelectionStart prop Gets or sets the starting position or offset of a text selection.
func (*HTMLTextAreaElement) SelectionStart() (selectionStart uint) {
	js.Rewrite("$<.selectionStart")
	return selectionStart
}

// SelectionStart prop Gets or sets the starting position or offset of a text selection.
func (*HTMLTextAreaElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.selectionStart = $1", selectionStart)
}

// Status prop Sets or retrieves the value indicating whether the control is selected.
func (*HTMLTextAreaElement) Status() (status interface{}) {
	js.Rewrite("$<.status")
	return status
}

// Status prop Sets or retrieves the value indicating whether the control is selected.
func (*HTMLTextAreaElement) SetStatus(status interface{}) {
	js.Rewrite("$<.status = $1", status)
}

// Type prop Retrieves the type of control.
func (*HTMLTextAreaElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLTextAreaElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLTextAreaElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

// Value prop Retrieves or sets the text in the entry field of the textArea element.
func (*HTMLTextAreaElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Retrieves or sets the text in the entry field of the textArea element.
func (*HTMLTextAreaElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLTextAreaElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

// Wrap prop Sets or retrieves how to handle wordwrapping in the object.
func (*HTMLTextAreaElement) Wrap() (wrap string) {
	js.Rewrite("$<.wrap")
	return wrap
}

// Wrap prop Sets or retrieves how to handle wordwrapping in the object.
func (*HTMLTextAreaElement) SetWrap(wrap string) {
	js.Rewrite("$<.wrap = $1", wrap)
}
