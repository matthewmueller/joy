package htmlparamelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLParamElement struct
// js:"HTMLParamElement,omit"
type HTMLParamElement struct {
	window.HTMLElement
}

// Name prop Sets or retrieves the name of an input parameter for an element.
func (*HTMLParamElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of an input parameter for an element.
func (*HTMLParamElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Type prop Sets or retrieves the content type of the resource designated by the value attribute.
func (*HTMLParamElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Sets or retrieves the content type of the resource designated by the value attribute.
func (*HTMLParamElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

// Value prop Sets or retrieves the value of an input parameter for an element.
func (*HTMLParamElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Sets or retrieves the value of an input parameter for an element.
func (*HTMLParamElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

// ValueType prop Sets or retrieves the data type of the value attribute.
func (*HTMLParamElement) ValueType() (valueType string) {
	js.Rewrite("$<.valueType")
	return valueType
}

// ValueType prop Sets or retrieves the data type of the value attribute.
func (*HTMLParamElement) SetValueType(valueType string) {
	js.Rewrite("$<.valueType = $1", valueType)
}
