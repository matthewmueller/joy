package htmlmeterelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLMeterElement struct
// js:"HTMLMeterElement,omit"
type HTMLMeterElement struct {
	window.HTMLElement
}

// High prop
func (*HTMLMeterElement) High() (high float32) {
	js.Rewrite("$<.high")
	return high
}

// High prop
func (*HTMLMeterElement) SetHigh(high float32) {
	js.Rewrite("$<.high = $1", high)
}

// Low prop
func (*HTMLMeterElement) Low() (low float32) {
	js.Rewrite("$<.low")
	return low
}

// Low prop
func (*HTMLMeterElement) SetLow(low float32) {
	js.Rewrite("$<.low = $1", low)
}

// Max prop
func (*HTMLMeterElement) Max() (max float32) {
	js.Rewrite("$<.max")
	return max
}

// Max prop
func (*HTMLMeterElement) SetMax(max float32) {
	js.Rewrite("$<.max = $1", max)
}

// Min prop
func (*HTMLMeterElement) Min() (min float32) {
	js.Rewrite("$<.min")
	return min
}

// Min prop
func (*HTMLMeterElement) SetMin(min float32) {
	js.Rewrite("$<.min = $1", min)
}

// Optimum prop
func (*HTMLMeterElement) Optimum() (optimum float32) {
	js.Rewrite("$<.optimum")
	return optimum
}

// Optimum prop
func (*HTMLMeterElement) SetOptimum(optimum float32) {
	js.Rewrite("$<.optimum = $1", optimum)
}

// Value prop
func (*HTMLMeterElement) Value() (value float32) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*HTMLMeterElement) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
