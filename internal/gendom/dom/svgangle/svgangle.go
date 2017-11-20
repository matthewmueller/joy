package svgangle

import "github.com/matthewmueller/golly/js"

type SVGAngle struct {
}

func (*SVGAngle) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.convertToSpecifiedUnits($1)", unitType)
}

func (*SVGAngle) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

func (*SVGAngle) GetUnitType() (unitType uint8) {
	js.Rewrite("$<.unitType")
	return unitType
}

func (*SVGAngle) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*SVGAngle) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

func (*SVGAngle) GetValueAsString() (valueAsString string) {
	js.Rewrite("$<.valueAsString")
	return valueAsString
}

func (*SVGAngle) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.valueAsString = $1", valueAsString)
}

func (*SVGAngle) GetValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

func (*SVGAngle) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
