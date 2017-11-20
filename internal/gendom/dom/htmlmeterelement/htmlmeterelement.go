package htmlmeterelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLMeterElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLMeterElement) GetHigh() (high float32) {
	js.Rewrite("$<.high")
	return high
}

func (*HTMLMeterElement) SetHigh(high float32) {
	js.Rewrite("$<.high = $1", high)
}

func (*HTMLMeterElement) GetLow() (low float32) {
	js.Rewrite("$<.low")
	return low
}

func (*HTMLMeterElement) SetLow(low float32) {
	js.Rewrite("$<.low = $1", low)
}

func (*HTMLMeterElement) GetMax() (max float32) {
	js.Rewrite("$<.max")
	return max
}

func (*HTMLMeterElement) SetMax(max float32) {
	js.Rewrite("$<.max = $1", max)
}

func (*HTMLMeterElement) GetMin() (min float32) {
	js.Rewrite("$<.min")
	return min
}

func (*HTMLMeterElement) SetMin(min float32) {
	js.Rewrite("$<.min = $1", min)
}

func (*HTMLMeterElement) GetOptimum() (optimum float32) {
	js.Rewrite("$<.optimum")
	return optimum
}

func (*HTMLMeterElement) SetOptimum(optimum float32) {
	js.Rewrite("$<.optimum = $1", optimum)
}

func (*HTMLMeterElement) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLMeterElement) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
