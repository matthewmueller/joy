package audioparam

import "github.com/matthewmueller/golly/js"

// AudioParam struct
// js:"AudioParam,omit"
type AudioParam struct {
}

// CancelScheduledValues
func (*AudioParam) CancelScheduledValues(startTime float32) (a *AudioParam) {
	js.Rewrite("$<.CancelScheduledValues($1)", startTime)
	return a
}

// ExponentialRampToValueAtTime
func (*AudioParam) ExponentialRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.ExponentialRampToValueAtTime($1, $2)", value, endTime)
	return a
}

// LinearRampToValueAtTime
func (*AudioParam) LinearRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.LinearRampToValueAtTime($1, $2)", value, endTime)
	return a
}

// SetTargetAtTime
func (*AudioParam) SetTargetAtTime(target float32, startTime float32, timeConstant float32) (a *AudioParam) {
	js.Rewrite("$<.SetTargetAtTime($1, $2, $3)", target, startTime, timeConstant)
	return a
}

// SetValueAtTime
func (*AudioParam) SetValueAtTime(value float32, startTime float32) (a *AudioParam) {
	js.Rewrite("$<.SetValueAtTime($1, $2)", value, startTime)
	return a
}

// SetValueCurveAtTime
func (*AudioParam) SetValueCurveAtTime(values []float32, startTime float32, duration float32) (a *AudioParam) {
	js.Rewrite("$<.SetValueCurveAtTime($1, $2, $3)", values, startTime, duration)
	return a
}

// DefaultValue
func (*AudioParam) DefaultValue() (defaultValue float32) {
	js.Rewrite("$<.DefaultValue")
	return defaultValue
}

// Value
func (*AudioParam) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*AudioParam) SetValue(value float32) {
	js.Rewrite("$<.Value = $1", value)
}
