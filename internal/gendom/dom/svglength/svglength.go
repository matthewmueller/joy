package svglength

import "github.com/matthewmueller/golly/js"

type SVGLength struct {
}

func (*SVGLength) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.convertToSpecifiedUnits($1)", unitType)
}

func (*SVGLength) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

func (*SVGLength) GetUnitType() (unitType uint8) {
	js.Rewrite("$<.unitType")
	return unitType
}

func (*SVGLength) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*SVGLength) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

func (*SVGLength) GetValueAsString() (valueAsString string) {
	js.Rewrite("$<.valueAsString")
	return valueAsString
}

func (*SVGLength) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.valueAsString = $1", valueAsString)
}

func (*SVGLength) GetValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

func (*SVGLength) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}
