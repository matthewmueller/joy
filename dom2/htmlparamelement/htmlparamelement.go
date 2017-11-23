package htmlparamelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLParamElement struct
// js:"HTMLParamElement,omit"
type HTMLParamElement struct {
	window.HTMLElement
}

// Name Sets or retrieves the name of an input parameter for an element.
func (*HTMLParamElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of an input parameter for an element.
func (*HTMLParamElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Type Sets or retrieves the content type of the resource designated by the value attribute.
func (*HTMLParamElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Sets or retrieves the content type of the resource designated by the value attribute.
func (*HTMLParamElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}

// Value Sets or retrieves the value of an input parameter for an element.
func (*HTMLParamElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value Sets or retrieves the value of an input parameter for an element.
func (*HTMLParamElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}

// ValueType Sets or retrieves the data type of the value attribute.
func (*HTMLParamElement) ValueType() (valueType string) {
	js.Rewrite("$<.ValueType")
	return valueType
}

// ValueType Sets or retrieves the data type of the value attribute.
func (*HTMLParamElement) SetValueType(valueType string) {
	js.Rewrite("$<.ValueType = $1", valueType)
}
