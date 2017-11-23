package htmlinputelement

import (
	"time"

	"github.com/matthewmueller/golly/dom/filelist"
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/validitystate"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLInputElement struct
// js:"HTMLInputElement,omit"
type HTMLInputElement struct {
	window.HTMLElement
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLInputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// Select fn Makes the selection equal to the current object.
func (*HTMLInputElement) Select() {
	js.Rewrite("$<.select()")
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLInputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

// SetSelectionRange fn Sets the start and end positions of a selection in a text field.
//     * @param start The offset into the text field for the start of the selection.
//     * @param end The offset into the text field for the end of the selection.
func (*HTMLInputElement) SetSelectionRange(start *uint, end *uint, direction *string) {
	js.Rewrite("$<.setSelectionRange($1, $2, $3)", start, end, direction)
}

// StepDown fn Decrements a range input control's value by the value given by the Step attribute. If the optional parameter is used, it will decrement the input control's step value multiplied by the parameter's value.
//     * @param n Value to decrement the value by.
func (*HTMLInputElement) StepDown(n *int) {
	js.Rewrite("$<.stepDown($1)", n)
}

// StepUp fn Increments a range input control's value by the value given by the Step attribute. If the optional parameter is used, will increment the input control's value by that value.
//     * @param n Value to increment the value by.
func (*HTMLInputElement) StepUp(n *int) {
	js.Rewrite("$<.stepUp($1)", n)
}

// Accept prop Sets or retrieves a comma-separated list of content types.
func (*HTMLInputElement) Accept() (accept string) {
	js.Rewrite("$<.accept")
	return accept
}

// Accept prop Sets or retrieves a comma-separated list of content types.
func (*HTMLInputElement) SetAccept(accept string) {
	js.Rewrite("$<.accept = $1", accept)
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLInputElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLInputElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLInputElement) Alt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLInputElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

// Autocomplete prop Specifies whether autocomplete is applied to an editable text field.
func (*HTMLInputElement) Autocomplete() (autocomplete string) {
	js.Rewrite("$<.autocomplete")
	return autocomplete
}

// Autocomplete prop Specifies whether autocomplete is applied to an editable text field.
func (*HTMLInputElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.autocomplete = $1", autocomplete)
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLInputElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLInputElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

// Border prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// Checked prop Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) Checked() (checked bool) {
	js.Rewrite("$<.checked")
	return checked
}

// Checked prop Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) SetChecked(checked bool) {
	js.Rewrite("$<.checked = $1", checked)
}

// Complete prop Retrieves whether the object is fully loaded.
func (*HTMLInputElement) Complete() (complete bool) {
	js.Rewrite("$<.complete")
	return complete
}

// DefaultChecked prop Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) DefaultChecked() (defaultChecked bool) {
	js.Rewrite("$<.defaultChecked")
	return defaultChecked
}

// DefaultChecked prop Sets or retrieves the state of the check box or radio button.
func (*HTMLInputElement) SetDefaultChecked(defaultChecked bool) {
	js.Rewrite("$<.defaultChecked = $1", defaultChecked)
}

// DefaultValue prop Sets or retrieves the initial contents of the object.
func (*HTMLInputElement) DefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

// DefaultValue prop Sets or retrieves the initial contents of the object.
func (*HTMLInputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

// Disabled prop
func (*HTMLInputElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLInputElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Files prop Returns a FileList object on a file type input object.
func (*HTMLInputElement) Files() (files *filelist.FileList) {
	js.Rewrite("$<.files")
	return files
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLInputElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// FormAction prop Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLInputElement) FormAction() (formAction string) {
	js.Rewrite("$<.formAction")
	return formAction
}

// FormAction prop Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLInputElement) SetFormAction(formAction string) {
	js.Rewrite("$<.formAction = $1", formAction)
}

// FormEnctype prop Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLInputElement) FormEnctype() (formEnctype string) {
	js.Rewrite("$<.formEnctype")
	return formEnctype
}

// FormEnctype prop Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLInputElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.formEnctype = $1", formEnctype)
}

// FormMethod prop Overrides the submit method attribute previously specified on a form element.
func (*HTMLInputElement) FormMethod() (formMethod string) {
	js.Rewrite("$<.formMethod")
	return formMethod
}

// FormMethod prop Overrides the submit method attribute previously specified on a form element.
func (*HTMLInputElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.formMethod = $1", formMethod)
}

// FormNoValidate prop Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLInputElement) FormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.formNoValidate")
	return formNoValidate
}

// FormNoValidate prop Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLInputElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.formNoValidate = $1", formNoValidate)
}

// FormTarget prop Overrides the target attribute on a form element.
func (*HTMLInputElement) FormTarget() (formTarget string) {
	js.Rewrite("$<.formTarget")
	return formTarget
}

// FormTarget prop Overrides the target attribute on a form element.
func (*HTMLInputElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.formTarget = $1", formTarget)
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLInputElement) Height() (height string) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLInputElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

// Hspace prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) Hspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

// Hspace prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLInputElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

// Indeterminate prop
func (*HTMLInputElement) Indeterminate() (indeterminate bool) {
	js.Rewrite("$<.indeterminate")
	return indeterminate
}

// Indeterminate prop
func (*HTMLInputElement) SetIndeterminate(indeterminate bool) {
	js.Rewrite("$<.indeterminate = $1", indeterminate)
}

// List prop Specifies the ID of a pre-defined datalist of options for an input element.
func (*HTMLInputElement) List() (list window.HTMLElement) {
	js.Rewrite("$<.list")
	return list
}

// Max prop Defines the maximum acceptable value for an input element with type="number".When used with the min and step attributes, lets you control the range and increment (such as only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) Max() (max string) {
	js.Rewrite("$<.max")
	return max
}

// Max prop Defines the maximum acceptable value for an input element with type="number".When used with the min and step attributes, lets you control the range and increment (such as only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) SetMax(max string) {
	js.Rewrite("$<.max = $1", max)
}

// MaxLength prop Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLInputElement) MaxLength() (maxLength int) {
	js.Rewrite("$<.maxLength")
	return maxLength
}

// MaxLength prop Sets or retrieves the maximum number of characters that the user can enter in a text control.
func (*HTMLInputElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.maxLength = $1", maxLength)
}

// Min prop Defines the minimum acceptable value for an input element with type="number". When used with the max and step attributes, lets you control the range and increment (such as even numbers only) that the user can enter into an input field.
func (*HTMLInputElement) Min() (min string) {
	js.Rewrite("$<.min")
	return min
}

// Min prop Defines the minimum acceptable value for an input element with type="number". When used with the max and step attributes, lets you control the range and increment (such as even numbers only) that the user can enter into an input field.
func (*HTMLInputElement) SetMin(min string) {
	js.Rewrite("$<.min = $1", min)
}

// Multiple prop Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLInputElement) Multiple() (multiple bool) {
	js.Rewrite("$<.multiple")
	return multiple
}

// Multiple prop Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (*HTMLInputElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.multiple = $1", multiple)
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLInputElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLInputElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Pattern prop Gets or sets a string containing a regular expression that the user's input must match.
func (*HTMLInputElement) Pattern() (pattern string) {
	js.Rewrite("$<.pattern")
	return pattern
}

// Pattern prop Gets or sets a string containing a regular expression that the user's input must match.
func (*HTMLInputElement) SetPattern(pattern string) {
	js.Rewrite("$<.pattern = $1", pattern)
}

// Placeholder prop Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLInputElement) Placeholder() (placeholder string) {
	js.Rewrite("$<.placeholder")
	return placeholder
}

// Placeholder prop Gets or sets a text string that is displayed in an input field as a hint or prompt to users as the format or type of information they need to enter.The text appears in an input field until the user puts focus on the field.
func (*HTMLInputElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.placeholder = $1", placeholder)
}

// ReadOnly prop
func (*HTMLInputElement) ReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

// ReadOnly prop
func (*HTMLInputElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.readOnly = $1", readOnly)
}

// Required prop When present, marks an element that can't be submitted without a value.
func (*HTMLInputElement) Required() (required bool) {
	js.Rewrite("$<.required")
	return required
}

// Required prop When present, marks an element that can't be submitted without a value.
func (*HTMLInputElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

// SelectionDirection prop
func (*HTMLInputElement) SelectionDirection() (selectionDirection string) {
	js.Rewrite("$<.selectionDirection")
	return selectionDirection
}

// SelectionDirection prop
func (*HTMLInputElement) SetSelectionDirection(selectionDirection string) {
	js.Rewrite("$<.selectionDirection = $1", selectionDirection)
}

// SelectionEnd prop Gets or sets the end position or offset of a text selection.
func (*HTMLInputElement) SelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.selectionEnd")
	return selectionEnd
}

// SelectionEnd prop Gets or sets the end position or offset of a text selection.
func (*HTMLInputElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.selectionEnd = $1", selectionEnd)
}

// SelectionStart prop Gets or sets the starting position or offset of a text selection.
func (*HTMLInputElement) SelectionStart() (selectionStart uint) {
	js.Rewrite("$<.selectionStart")
	return selectionStart
}

// SelectionStart prop Gets or sets the starting position or offset of a text selection.
func (*HTMLInputElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.selectionStart = $1", selectionStart)
}

// Size prop
func (*HTMLInputElement) Size() (size uint) {
	js.Rewrite("$<.size")
	return size
}

// Size prop
func (*HTMLInputElement) SetSize(size uint) {
	js.Rewrite("$<.size = $1", size)
}

// Src prop The address or URL of the a media resource that is to be considered.
func (*HTMLInputElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop The address or URL of the a media resource that is to be considered.
func (*HTMLInputElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Status prop
func (*HTMLInputElement) Status() (status bool) {
	js.Rewrite("$<.status")
	return status
}

// Status prop
func (*HTMLInputElement) SetStatus(status bool) {
	js.Rewrite("$<.status = $1", status)
}

// Step prop Defines an increment or jump between values that you want to allow the user to enter. When used with the max and min attributes, lets you control the range and increment (for example, allow only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) Step() (step string) {
	js.Rewrite("$<.step")
	return step
}

// Step prop Defines an increment or jump between values that you want to allow the user to enter. When used with the max and min attributes, lets you control the range and increment (for example, allow only even numbers) that the user can enter into an input field.
func (*HTMLInputElement) SetStep(step string) {
	js.Rewrite("$<.step = $1", step)
}

// Type prop Returns the content type of the object.
func (*HTMLInputElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Returns the content type of the object.
func (*HTMLInputElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLInputElement) UseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLInputElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLInputElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLInputElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

// Value prop Returns the value of the data at the cursor's current position.
func (*HTMLInputElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Returns the value of the data at the cursor's current position.
func (*HTMLInputElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

// ValueAsDate prop
func (*HTMLInputElement) ValueAsDate() (valueAsDate time.Time) {
	js.Rewrite("$<.valueAsDate")
	return valueAsDate
}

// ValueAsDate prop
func (*HTMLInputElement) SetValueAsDate(valueAsDate time.Time) {
	js.Rewrite("$<.valueAsDate = $1", valueAsDate)
}

// ValueAsNumber prop Returns the input field value as a number.
func (*HTMLInputElement) ValueAsNumber() (valueAsNumber float32) {
	js.Rewrite("$<.valueAsNumber")
	return valueAsNumber
}

// ValueAsNumber prop Returns the input field value as a number.
func (*HTMLInputElement) SetValueAsNumber(valueAsNumber float32) {
	js.Rewrite("$<.valueAsNumber = $1", valueAsNumber)
}

// Vspace prop Sets or retrieves the vertical margin for the object.
func (*HTMLInputElement) Vspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

// Vspace prop Sets or retrieves the vertical margin for the object.
func (*HTMLInputElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

// Webkitdirectory prop
func (*HTMLInputElement) Webkitdirectory() (webkitdirectory bool) {
	js.Rewrite("$<.webkitdirectory")
	return webkitdirectory
}

// Webkitdirectory prop
func (*HTMLInputElement) SetWebkitdirectory(webkitdirectory bool) {
	js.Rewrite("$<.webkitdirectory = $1", webkitdirectory)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLInputElement) Width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLInputElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLInputElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
