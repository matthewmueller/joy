package svgangle

import "github.com/matthewmueller/golly/js"

// SVGAngle struct
// js:"SVGAngle,omit"
type SVGAngle struct {
}

// ConvertToSpecifiedUnits fn
func (*SVGAngle) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.convertToSpecifiedUnits($1)", unitType)
}

// NewValueSpecifiedUnits fn
func (*SVGAngle) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

// UnitType prop
func (*SVGAngle) UnitType() (unitType uint8) {
	js.Rewrite("$<.unitType")
	return unitType
}

// Value prop
func (*SVGAngle) Value() (value float32) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*SVGAngle) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

// ValueAsString prop
func (*SVGAngle) ValueAsString() (valueAsString string) {
	js.Rewrite("$<.valueAsString")
	return valueAsString
}

// ValueAsString prop
func (*SVGAngle) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.valueAsString = $1", valueAsString)
}

// ValueInSpecifiedUnits prop
func (*SVGAngle) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

// ValueInSpecifiedUnits prop
func (*SVGAngle) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
