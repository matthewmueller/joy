package htmlinputelement

import (
	"time"

	"github.com/matthewmueller/golly/dom2/filelist"
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLInputElement struct
// js:"HTMLInputElement,omit"
type HTMLInputElement struct {
	window.HTMLElement
}

// CheckValidity Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLInputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// Select Makes the selection equal to the current object.
func (*HTMLInputElement) Select() {
	js.Rewrite("$<.Select()")
}

// SetCustomValidity Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLInputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.SetCustomValidity($1)", err)
}

// SetSelectionRange Sets the start and end positions of a selection in a text field.
//     * @param start The offset into the text field for the start of the selection.
//     * @param end The offset into the text field for the end of the selection.
func (*HTMLInputElement) SetSelectionRange(start *uint, end *uint, direction *string) {
	js.Rewrite("$<.SetSelectionRange($1, $2, $3)", start, end, direction)
}

// StepDown Decrements a range input control's value by the value given by the Step attribute. If the optional parameter is used, it will decrement the input control's step value multiplied by the parameter's value.
//     * @param n Value to decrement the value by.
func (*HTMLInputElement) StepDown(n *int) {
	js.Rewrite("$<.StepDown($1)", n)
}

// StepUp Increments a range input control's value by the value given by the Step attribute. If the optional parameter is used, will increment the input control's value by that value.
//     * @param n Value to increment the value by.
func (*HTMLInputElement) StepUp(n *int) {
	js.Rewrite("$<.StepUp($1)", n)
}

// Accept Sets or retrieves a comma-separated list of content types.
func (*HTMLInputElement) Accept() (accept string) {
	js.Rewrite("$<.Accept")
	return accept
}

// Accept Sets or retrieves a comma-separated list of content types.
func (*HTMLInputElement) SetAccept(accept string) {
	js.Rewrite("$<.Accept = $1", accept)
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLInputElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLInputElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLInputElement) Alt() (alt string) {
	js.Rewrite("$<.Alt")
	return alt
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLInputElement) SetAlt(alt string) {
	js.Rewrite("$<.Alt = $1", alt)
}

// Autocomplete Specifies whether autocomplete is applied to an editable text field.
func (*HTMLInputElement) Autocomplete() (autocomplete string) {
	js.Rewrite("$<.Autocomplete")
	return autocomplete
}

// Autocomplete Specifies whether autocomplete is applied to an editable text field.
func (*HTMLInputElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.Autocomplete = $1", autocomplete)
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLInputElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.Autofocus")
	return autofocus
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLInputElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.Autofocus = $1", autofocus)
}

// Border Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// Checked Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) Checked() (checked bool) {
	js.Rewrite("$<.Checked")
	return checked
}

// Checked Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) SetChecked(checked bool) {
	js.Rewrite("$<.Checked = $1", checked)
}

// Complete Retrieves whether the object is fully loaded.
func (*HTMLInputElement) Complete() (complete bool) {
	js.Rewrite("$<.Complete")
	return complete
}

// DefaultChecked Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) DefaultChecked() (defaultChecked bool) {
	js.Rewrite("$<.DefaultChecked")
	return defaultChecked
}

// DefaultChecked Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) SetDefaultChecked(defaultChecked bool) {
	js.Rewrite("$<.DefaultChecked = $1", defaultChecked)
}

// DefaultValue Sets or retrieves the initial contents of the object.
func (*HTMLInputElement) DefaultValue() (defaultValue string) {
	js.Rewrite("$<.DefaultValue")
	return defaultValue
}

// DefaultValue Sets or retrieves the initial contents of the object.
func (*HTMLInputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.DefaultValue = $1", defaultValue)
}

// Disabled
func (*HTMLInputElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLInputElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Files Returns a FileList object on a file type input object.
func (*HTMLInputElement) Files() (files *filelist.FileList) {
	js.Rewrite("$<.Files")
	return files
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLInputElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// FormAction Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLInputElement) FormAction() (formAction string) {
	js.Rewrite("$<.FormAction")
	return formAction
}

// FormAction Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLInputElement) SetFormAction(formAction string) {
	js.Rewrite("$<.FormAction = $1", formAction)
}

// FormEnctype Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLInputElement) FormEnctype() (formEnctype string) {
	js.Rewrite("$<.FormEnctype")
	return formEnctype
}

// FormEnctype Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLInputElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.FormEnctype = $1", formEnctype)
}

// FormMethod Overrides the submit method attribute previously specified on a form element.
func (*HTMLInputElement) FormMethod() (formMethod string) {
	js.Rewrite("$<.FormMethod")
	return formMethod
}

// FormMethod Overrides the submit method attribute previously specified on a form element.
func (*HTMLInputElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.FormMethod = $1", formMethod)
}

// FormNoValidate Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLInputElement) FormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.FormNoValidate")
	return formNoValidate
}

// FormNoValidate Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLInputElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.FormNoValidate = $1", formNoValidate)
}

// FormTarget Overrides the target attribute on a form element.
func (*HTMLInputElement) FormTarget() (formTarget string) {
	js.Rewrite("$<.FormTarget")
	return formTarget
}

// FormTarget Overrides the target attribute on a form element.
func (*HTMLInputElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.FormTarget = $1", formTarget)
}

// Height Sets or retrieves the height of the object.
func (*HTMLInputElement) Height() (height string) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLInputElement) SetHeight(height string) {
	js.Rewrite("$<.Height = $1", height)
}

// Hspace Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) Hspace() (hspace int) {
	js.Rewrite("$<.Hspace")
	return hspace
}

// Hspace Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) SetHspace(hspace int) {
	js.Rewrite("$<.Hspace = $1", hspace)
}

// Indeterminate
func (*HTMLInputElement) Indeterminate() (indeterminate bool) {
	js.Rewrite("$<.Indeterminate")
	return indeterminate
}

// Indeterminate
func (*HTMLInputElement) SetIndeterminate(indeterminate bool) {
	js.Rewrite("$<.Indeterminate = $1", indeterminate)
}

// List Specifies the ID of a pre-defined datalist of options for an input element.
func (*HTMLInputElement) List() (list window.HTMLElement) {
	js.Rewrite("$<.List")
	return list
}

// Max Defines the maximum acceptable value for an input element with type="number".When used with the min and step attributes, lets you control the range and increment (such as only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) Max() (max string) {
	js.Rewrite("$<.Max")
	return max
}

// Max Defines the maximum acceptable value for an input element with type="number".When used with the min and step attributes, lets you control the range and increment (such as only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) SetMax(max string) {
	js.Rewrite("$<.Max = $1", max)
}

// MaxLength Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLInputElement) MaxLength() (maxLength int) {
	js.Rewrite("$<.MaxLength")
	return maxLength
}

// MaxLength Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLInputElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.MaxLength = $1", maxLength)
}

// Min Defines the minimum acceptable value for an input element with type="number". When used with the max and step attributes, lets you control the range and increment (such as even numbers only) that the user can enter into an input field.
func (*HTMLInputElement) Min() (min string) {
	js.Rewrite("$<.Min")
	return min
}

// Min Defines the minimum acceptable value for an input element with type="number". When used with the max and step attributes, lets you control the range and increment (such as even numbers only) that the user can enter into an input field.
func (*HTMLInputElement) SetMin(min string) {
	js.Rewrite("$<.Min = $1", min)
}

// Multiple Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLInputElement) Multiple() (multiple bool) {
	js.Rewrite("$<.Multiple")
	return multiple
}

// Multiple Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLInputElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.Multiple = $1", multiple)
}

// Name Sets or retrieves the name of the object.
func (*HTMLInputElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLInputElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Pattern Gets or sets a string containing a regular expression that the user's input must match.
func (*HTMLInputElement) Pattern() (pattern string) {
	js.Rewrite("$<.Pattern")
	return pattern
}

// Pattern Gets or sets a string containing a regular expression that the user's input must match.
func (*HTMLInputElement) SetPattern(pattern string) {
	js.Rewrite("$<.Pattern = $1", pattern)
}

// Placeholder Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLInputElement) Placeholder() (placeholder string) {
	js.Rewrite("$<.Placeholder")
	return placeholder
}

// Placeholder Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLInputElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.Placeholder = $1", placeholder)
}

// ReadOnly
func (*HTMLInputElement) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.ReadOnly")
	return readOnly
}

// ReadOnly
func (*HTMLInputElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.ReadOnly = $1", readOnly)
}

// Required When present, marks an element that can't be submitted without a value.
func (*HTMLInputElement) Required() (required bool) {
	js.Rewrite("$<.Required")
	return required
}

// Required When present, marks an element that can't be submitted without a value.
func (*HTMLInputElement) SetRequired(required bool) {
	js.Rewrite("$<.Required = $1", required)
}

// SelectionDirection
func (*HTMLInputElement) SelectionDirection() (selectionDirection string) {
	js.Rewrite("$<.SelectionDirection")
	return selectionDirection
}

// SelectionDirection
func (*HTMLInputElement) SetSelectionDirection(selectionDirection string) {
	js.Rewrite("$<.SelectionDirection = $1", selectionDirection)
}

// SelectionEnd Gets or sets the end position or offset of a text selection.
func (*HTMLInputElement) SelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.SelectionEnd")
	return selectionEnd
}

// SelectionEnd Gets or sets the end position or offset of a text selection.
func (*HTMLInputElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.SelectionEnd = $1", selectionEnd)
}

// SelectionStart Gets or sets the starting position or offset of a text selection.
func (*HTMLInputElement) SelectionStart() (selectionStart uint) {
	js.Rewrite("$<.SelectionStart")
	return selectionStart
}

// SelectionStart Gets or sets the starting position or offset of a text selection.
func (*HTMLInputElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.SelectionStart = $1", selectionStart)
}

// Size
func (*HTMLInputElement) Size() (size uint) {
	js.Rewrite("$<.Size")
	return size
}

// Size
func (*HTMLInputElement) SetSize(size uint) {
	js.Rewrite("$<.Size = $1", size)
}

// Src The address or URL of the a media resource that is to be considered.
func (*HTMLInputElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src The address or URL of the a media resource that is to be considered.
func (*HTMLInputElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Status
func (*HTMLInputElement) Status() (status bool) {
	js.Rewrite("$<.Status")
	return status
}

// Status
func (*HTMLInputElement) SetStatus(status bool) {
	js.Rewrite("$<.Status = $1", status)
}

// Step Defines an increment or jump between values that you want to allow the user to enter. When used with the max and min attributes, lets you control the range and increment (for example, allow only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) Step() (step string) {
	js.Rewrite("$<.Step")
	return step
}

// Step Defines an increment or jump between values that you want to allow the user to enter. When used with the max and min attributes, lets you control the range and increment (for example, allow only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) SetStep(step string) {
	js.Rewrite("$<.Step = $1", step)
}

// Type Returns the content type of the object.
func (*HTMLInputElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Returns the content type of the object.
func (*HTMLInputElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLInputElement) UseMap() (useMap string) {
	js.Rewrite("$<.UseMap")
	return useMap
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLInputElement) SetUseMap(useMap string) {
	js.Rewrite("$<.UseMap = $1", useMap)
}

// ValidationMessage Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLInputElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.ValidationMessage")
	return validationMessage
}

// Validity Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLInputElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.Validity")
	return validity
}

// Value Returns the value of the data at the cursor's current position.
func (*HTMLInputElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value Returns the value of the data at the cursor's current position.
func (*HTMLInputElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}

// ValueAsDate
func (*HTMLInputElement) ValueAsDate() (valueAsDate time.Time) {
	js.Rewrite("$<.ValueAsDate")
	return valueAsDate
}

// ValueAsDate
func (*HTMLInputElement) SetValueAsDate(valueAsDate time.Time) {
	js.Rewrite("$<.ValueAsDate = $1", valueAsDate)
}

// ValueAsNumber Returns the input field value as a number.
func (*HTMLInputElement) ValueAsNumber() (valueAsNumber float32) {
	js.Rewrite("$<.ValueAsNumber")
	return valueAsNumber
}

// ValueAsNumber Returns the input field value as a number.
func (*HTMLInputElement) SetValueAsNumber(valueAsNumber float32) {
	js.Rewrite("$<.ValueAsNumber = $1", valueAsNumber)
}

// Vspace Sets or retrieves the vertical margin for the object.
func (*HTMLInputElement) Vspace() (vspace int) {
	js.Rewrite("$<.Vspace")
	return vspace
}

// Vspace Sets or retrieves the vertical margin for the object.
func (*HTMLInputElement) SetVspace(vspace int) {
	js.Rewrite("$<.Vspace = $1", vspace)
}

// Webkitdirectory
func (*HTMLInputElement) Webkitdirectory() (webkitdirectory bool) {
	js.Rewrite("$<.Webkitdirectory")
	return webkitdirectory
}

// Webkitdirectory
func (*HTMLInputElement) SetWebkitdirectory(webkitdirectory bool) {
	js.Rewrite("$<.Webkitdirectory = $1", webkitdirectory)
}

// Width Sets or retrieves the width of the object.
func (*HTMLInputElement) Width() (width string) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLInputElement) SetWidth(width string) {
	js.Rewrite("$<.Width = $1", width)
}

// WillValidate Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLInputElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.WillValidate")
	return willValidate
}
