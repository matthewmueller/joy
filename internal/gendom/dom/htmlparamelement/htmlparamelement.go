package htmlparamelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLParamElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLParamElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLParamElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLParamElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLParamElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLParamElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLParamElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLParamElement) GetValueType() (valueType string) {
	js.Rewrite("$<.valueType")
	return valueType
}

func (*HTMLParamElement) SetValueType(valueType string) {
	js.Rewrite("$<.valueType = $1", valueType)
}
