package htmlmeterelement

import "github.com/matthewmueller/golly/js"

// js:"HTMLMeterElement,omit"
type HTMLMeterElement struct {
	window.HTMLElement
}

// High
func (*HTMLMeterElement) High() (high float32) {
	js.Rewrite("$<.High")
	return high
}

// High
func (*HTMLMeterElement) SetHigh(high float32) {
	js.Rewrite("$<.High = $1", high)
}

// Low
func (*HTMLMeterElement) Low() (low float32) {
	js.Rewrite("$<.Low")
	return low
}

// Low
func (*HTMLMeterElement) SetLow(low float32) {
	js.Rewrite("$<.Low = $1", low)
}

// Max
func (*HTMLMeterElement) Max() (max float32) {
	js.Rewrite("$<.Max")
	return max
}

// Max
func (*HTMLMeterElement) SetMax(max float32) {
	js.Rewrite("$<.Max = $1", max)
}

// Min
func (*HTMLMeterElement) Min() (min float32) {
	js.Rewrite("$<.Min")
	return min
}

// Min
func (*HTMLMeterElement) SetMin(min float32) {
	js.Rewrite("$<.Min = $1", min)
}

// Optimum
func (*HTMLMeterElement) Optimum() (optimum float32) {
	js.Rewrite("$<.Optimum")
	return optimum
}

// Optimum
func (*HTMLMeterElement) SetOptimum(optimum float32) {
	js.Rewrite("$<.Optimum = $1", optimum)
}

// Value
func (*HTMLMeterElement) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*HTMLMeterElement) SetValue(value float32) {
	js.Rewrite("$<.Value = $1", value)
}
