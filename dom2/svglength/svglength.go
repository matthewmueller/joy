package svglength

import "github.com/matthewmueller/golly/js"

// js:"SVGLength,omit"
type SVGLength struct {
}

// ConvertToSpecifiedUnits
func (*SVGLength) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.ConvertToSpecifiedUnits($1)", unitType)
}

// NewValueSpecifiedUnits
func (*SVGLength) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.NewValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

// UnitType
func (*SVGLength) UnitType() (unitType uint8) {
	js.Rewrite("$<.UnitType")
	return unitType
}

// Value
func (*SVGLength) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*SVGLength) SetValue(value float32) {
	js.Rewrite("$<.Value = $1", value)
}

// ValueAsString
func (*SVGLength) ValueAsString() (valueAsString string) {
	js.Rewrite("$<.ValueAsString")
	return valueAsString
}

// ValueAsString
func (*SVGLength) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.ValueAsString = $1", valueAsString)
}

// ValueInSpecifiedUnits
func (*SVGLength) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.ValueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

// ValueInSpecifiedUnits
func (*SVGLength) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.ValueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
