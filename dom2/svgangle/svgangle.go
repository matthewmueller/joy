package svgangle

import "github.com/matthewmueller/golly/js"

// SVGAngle struct
// js:"SVGAngle,omit"
type SVGAngle struct {
}

// ConvertToSpecifiedUnits
func (*SVGAngle) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.ConvertToSpecifiedUnits($1)", unitType)
}

// NewValueSpecifiedUnits
func (*SVGAngle) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.NewValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

// UnitType
func (*SVGAngle) UnitType() (unitType uint8) {
	js.Rewrite("$<.UnitType")
	return unitType
}

// Value
func (*SVGAngle) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*SVGAngle) SetValue(value float32) {
	js.Rewrite("$<.Value = $1", value)
}

// ValueAsString
func (*SVGAngle) ValueAsString() (valueAsString string) {
	js.Rewrite("$<.ValueAsString")
	return valueAsString
}

// ValueAsString
func (*SVGAngle) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.ValueAsString = $1", valueAsString)
}

// ValueInSpecifiedUnits
func (*SVGAngle) ValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.ValueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

// ValueInSpecifiedUnits
func (*SVGAngle) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.ValueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
