package htmltextareaelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/validitystate"
	"github.com/matthewmueller/golly/js"
)

type HTMLTextAreaElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLTextAreaElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLTextAreaElement) Select() {
	js.Rewrite("$<.select()")
}

func (*HTMLTextAreaElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLTextAreaElement) SetSelectionRange(start uint, end uint) {
	js.Rewrite("$<.setSelectionRange($1, $2)", start, end)
}

func (*HTMLTextAreaElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLTextAreaElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLTextAreaElement) GetCols() (cols uint) {
	js.Rewrite("$<.cols")
	return cols
}

func (*HTMLTextAreaElement) SetCols(cols uint) {
	js.Rewrite("$<.cols = $1", cols)
}

func (*HTMLTextAreaElement) GetDefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*HTMLTextAreaElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

func (*HTMLTextAreaElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLTextAreaElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLTextAreaElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLTextAreaElement) GetMaxLength() (maxLength int) {
	js.Rewrite("$<.maxLength")
	return maxLength
}

func (*HTMLTextAreaElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.maxLength = $1", maxLength)
}

func (*HTMLTextAreaElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLTextAreaElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLTextAreaElement) GetPlaceholder() (placeholder string) {
	js.Rewrite("$<.placeholder")
	return placeholder
}

func (*HTMLTextAreaElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.placeholder = $1", placeholder)
}

func (*HTMLTextAreaElement) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*HTMLTextAreaElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.readOnly = $1", readOnly)
}

func (*HTMLTextAreaElement) GetRequired() (required bool) {
	js.Rewrite("$<.required")
	return required
}

func (*HTMLTextAreaElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

func (*HTMLTextAreaElement) GetRows() (rows uint) {
	js.Rewrite("$<.rows")
	return rows
}

func (*HTMLTextAreaElement) SetRows(rows uint) {
	js.Rewrite("$<.rows = $1", rows)
}

func (*HTMLTextAreaElement) GetSelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.selectionEnd")
	return selectionEnd
}

func (*HTMLTextAreaElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.selectionEnd = $1", selectionEnd)
}

func (*HTMLTextAreaElement) GetSelectionStart() (selectionStart uint) {
	js.Rewrite("$<.selectionStart")
	return selectionStart
}

func (*HTMLTextAreaElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.selectionStart = $1", selectionStart)
}

func (*HTMLTextAreaElement) GetStatus() (status interface{}) {
	js.Rewrite("$<.status")
	return status
}

func (*HTMLTextAreaElement) SetStatus(status interface{}) {
	js.Rewrite("$<.status = $1", status)
}

func (*HTMLTextAreaElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLTextAreaElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLTextAreaElement) GetValidity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLTextAreaElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLTextAreaElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLTextAreaElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

func (*HTMLTextAreaElement) GetWrap() (wrap string) {
	js.Rewrite("$<.wrap")
	return wrap
}

func (*HTMLTextAreaElement) SetWrap(wrap string) {
	js.Rewrite("$<.wrap = $1", wrap)
}
