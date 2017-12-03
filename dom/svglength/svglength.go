package svglength

import "github.com/matthewmueller/joy/js"

// SVGLength struct
// js:"SVGLength,omit"
type SVGLength struct {
}

// ConvertToSpecifiedUnits fn
// js:"convertToSpecifiedUnits"
func (*SVGLength) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$_.convertToSpecifiedUnits($1)", unitType)
}

// NewValueSpecifiedUnits fn
// js:"newValueSpecifiedUnits"
func (*SVGLength) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$_.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

// UnitType prop
// js:"unitType"
func (*SVGLength) UnitType() (unitType uint8) {
	js.Rewrite("$_.unitType")
	return unitType
}

// Value prop
// js:"value"
func (*SVGLength) Value() (value float32) {
	js.Rewrite("$_.value")
	return value
}

// SetValue prop
// js:"value"
func (*SVGLength) SetValue(value float32) {
	js.Rewrite("$_.value = $1", value)
}

// ValueAsString prop
// js:"valueAsString"
func (*SVGLength) ValueAsString() (valueAsString string) {
	js.Rewrite("$_.valueAsString")
	return valueAsString
}

// SetValueAsString prop
// js:"valueAsString"
func (*SVGLength) SetValueAsString(valueAsString string) {
	js.Rewrite("$_.valueAsString = $1", valueAsString)
}

// ValueInSpecifiedUnits prop
// js:"valueInSpecifiedUnits"
func (*SVGLength) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$_.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

// SetValueInSpecifiedUnits prop
// js:"valueInSpecifiedUnits"
func (*SVGLength) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$_.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
