package svglength

import "github.com/matthewmueller/golly/js"

// SVGLength struct
// js:"SVGLength,omit"
type SVGLength struct {
}

// ConvertToSpecifiedUnits fn
func (*SVGLength) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.convertToSpecifiedUnits($1)", unitType)
}

// NewValueSpecifiedUnits fn
func (*SVGLength) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

// UnitType prop
func (*SVGLength) UnitType() (unitType uint8) {
	js.Rewrite("$<.unitType")
	return unitType
}

// Value prop
func (*SVGLength) Value() (value float32) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*SVGLength) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

// ValueAsString prop
func (*SVGLength) ValueAsString() (valueAsString string) {
	js.Rewrite("$<.valueAsString")
	return valueAsString
}

// ValueAsString prop
func (*SVGLength) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.valueAsString = $1", valueAsString)
}

// ValueInSpecifiedUnits prop
func (*SVGLength) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

// ValueInSpecifiedUnits prop
func (*SVGLength) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
