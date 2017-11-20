package htmlinputelement

import (
	"time"

	"github.com/matthewmueller/golly/internal/gendom/dom/filelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/validitystate"
	"github.com/matthewmueller/golly/js"
)

type HTMLInputElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLInputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLInputElement) Select() {
	js.Rewrite("$<.select()")
}

func (*HTMLInputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLInputElement) SetSelectionRange(start uint, end uint, direction string) {
	js.Rewrite("$<.setSelectionRange($1, $2, $3)", start, end, direction)
}

func (*HTMLInputElement) StepDown(n int) {
	js.Rewrite("$<.stepDown($1)", n)
}

func (*HTMLInputElement) StepUp(n int) {
	js.Rewrite("$<.stepUp($1)", n)
}

func (*HTMLInputElement) GetAccept() (accept string) {
	js.Rewrite("$<.accept")
	return accept
}

func (*HTMLInputElement) SetAccept(accept string) {
	js.Rewrite("$<.accept = $1", accept)
}

func (*HTMLInputElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLInputElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLInputElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLInputElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLInputElement) GetAutocomplete() (autocomplete string) {
	js.Rewrite("$<.autocomplete")
	return autocomplete
}

func (*HTMLInputElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.autocomplete = $1", autocomplete)
}

func (*HTMLInputElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLInputElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLInputElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLInputElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLInputElement) GetChecked() (checked bool) {
	js.Rewrite("$<.checked")
	return checked
}

func (*HTMLInputElement) SetChecked(checked bool) {
	js.Rewrite("$<.checked = $1", checked)
}

func (*HTMLInputElement) GetComplete() (complete bool) {
	js.Rewrite("$<.complete")
	return complete
}

func (*HTMLInputElement) GetDefaultChecked() (defaultChecked bool) {
	js.Rewrite("$<.defaultChecked")
	return defaultChecked
}

func (*HTMLInputElement) SetDefaultChecked(defaultChecked bool) {
	js.Rewrite("$<.defaultChecked = $1", defaultChecked)
}

func (*HTMLInputElement) GetDefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*HTMLInputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

func (*HTMLInputElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLInputElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLInputElement) GetFiles() (files *filelist.FileList) {
	js.Rewrite("$<.files")
	return files
}

func (*HTMLInputElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLInputElement) GetFormAction() (formAction string) {
	js.Rewrite("$<.formAction")
	return formAction
}

func (*HTMLInputElement) SetFormAction(formAction string) {
	js.Rewrite("$<.formAction = $1", formAction)
}

func (*HTMLInputElement) GetFormEnctype() (formEnctype string) {
	js.Rewrite("$<.formEnctype")
	return formEnctype
}

func (*HTMLInputElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.formEnctype = $1", formEnctype)
}

func (*HTMLInputElement) GetFormMethod() (formMethod string) {
	js.Rewrite("$<.formMethod")
	return formMethod
}

func (*HTMLInputElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.formMethod = $1", formMethod)
}

func (*HTMLInputElement) GetFormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.formNoValidate")
	return formNoValidate
}

func (*HTMLInputElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.formNoValidate = $1", formNoValidate)
}

func (*HTMLInputElement) GetFormTarget() (formTarget string) {
	js.Rewrite("$<.formTarget")
	return formTarget
}

func (*HTMLInputElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.formTarget = $1", formTarget)
}

func (*HTMLInputElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLInputElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLInputElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLInputElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLInputElement) GetIndeterminate() (indeterminate bool) {
	js.Rewrite("$<.indeterminate")
	return indeterminate
}

func (*HTMLInputElement) SetIndeterminate(indeterminate bool) {
	js.Rewrite("$<.indeterminate = $1", indeterminate)
}

func (*HTMLInputElement) GetList() (list *htmlelement.HTMLElement) {
	js.Rewrite("$<.list")
	return list
}

func (*HTMLInputElement) GetMax() (max string) {
	js.Rewrite("$<.max")
	return max
}

func (*HTMLInputElement) SetMax(max string) {
	js.Rewrite("$<.max = $1", max)
}

func (*HTMLInputElement) GetMaxLength() (maxLength int) {
	js.Rewrite("$<.maxLength")
	return maxLength
}

func (*HTMLInputElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.maxLength = $1", maxLength)
}

func (*HTMLInputElement) GetMin() (min string) {
	js.Rewrite("$<.min")
	return min
}

func (*HTMLInputElement) SetMin(min string) {
	js.Rewrite("$<.min = $1", min)
}

func (*HTMLInputElement) GetMultiple() (multiple bool) {
	js.Rewrite("$<.multiple")
	return multiple
}

func (*HTMLInputElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.multiple = $1", multiple)
}

func (*HTMLInputElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLInputElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLInputElement) GetPattern() (pattern string) {
	js.Rewrite("$<.pattern")
	return pattern
}

func (*HTMLInputElement) SetPattern(pattern string) {
	js.Rewrite("$<.pattern = $1", pattern)
}

func (*HTMLInputElement) GetPlaceholder() (placeholder string) {
	js.Rewrite("$<.placeholder")
	return placeholder
}

func (*HTMLInputElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.placeholder = $1", placeholder)
}

func (*HTMLInputElement) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*HTMLInputElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.readOnly = $1", readOnly)
}

func (*HTMLInputElement) GetRequired() (required bool) {
	js.Rewrite("$<.required")
	return required
}

func (*HTMLInputElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

func (*HTMLInputElement) GetSelectionDirection() (selectionDirection string) {
	js.Rewrite("$<.selectionDirection")
	return selectionDirection
}

func (*HTMLInputElement) SetSelectionDirection(selectionDirection string) {
	js.Rewrite("$<.selectionDirection = $1", selectionDirection)
}

func (*HTMLInputElement) GetSelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.selectionEnd")
	return selectionEnd
}

func (*HTMLInputElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.selectionEnd = $1", selectionEnd)
}

func (*HTMLInputElement) GetSelectionStart() (selectionStart uint) {
	js.Rewrite("$<.selectionStart")
	return selectionStart
}

func (*HTMLInputElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.selectionStart = $1", selectionStart)
}

func (*HTMLInputElement) GetSize() (size uint) {
	js.Rewrite("$<.size")
	return size
}

func (*HTMLInputElement) SetSize(size uint) {
	js.Rewrite("$<.size = $1", size)
}

func (*HTMLInputElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLInputElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLInputElement) GetStatus() (status bool) {
	js.Rewrite("$<.status")
	return status
}

func (*HTMLInputElement) SetStatus(status bool) {
	js.Rewrite("$<.status = $1", status)
}

func (*HTMLInputElement) GetStep() (step string) {
	js.Rewrite("$<.step")
	return step
}

func (*HTMLInputElement) SetStep(step string) {
	js.Rewrite("$<.step = $1", step)
}

func (*HTMLInputElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLInputElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLInputElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLInputElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLInputElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLInputElement) GetValidity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLInputElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLInputElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLInputElement) GetValueAsDate() (valueAsDate time.Time) {
	js.Rewrite("$<.valueAsDate")
	return valueAsDate
}

func (*HTMLInputElement) SetValueAsDate(valueAsDate time.Time) {
	js.Rewrite("$<.valueAsDate = $1", valueAsDate)
}

func (*HTMLInputElement) GetValueAsNumber() (valueAsNumber float32) {
	js.Rewrite("$<.valueAsNumber")
	return valueAsNumber
}

func (*HTMLInputElement) SetValueAsNumber(valueAsNumber float32) {
	js.Rewrite("$<.valueAsNumber = $1", valueAsNumber)
}

func (*HTMLInputElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLInputElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLInputElement) GetWebkitdirectory() (webkitdirectory bool) {
	js.Rewrite("$<.webkitdirectory")
	return webkitdirectory
}

func (*HTMLInputElement) SetWebkitdirectory(webkitdirectory bool) {
	js.Rewrite("$<.webkitdirectory = $1", webkitdirectory)
}

func (*HTMLInputElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLInputElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

func (*HTMLInputElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
