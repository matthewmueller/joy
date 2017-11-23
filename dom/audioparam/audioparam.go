package audioparam

import "github.com/matthewmueller/golly/js"

// AudioParam struct
// js:"AudioParam,omit"
type AudioParam struct {
}

// CancelScheduledValues fn
func (*AudioParam) CancelScheduledValues(startTime float32) (a *AudioParam) {
	js.Rewrite("$<.cancelScheduledValues($1)", startTime)
	return a
}

// ExponentialRampToValueAtTime fn
func (*AudioParam) ExponentialRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.exponentialRampToValueAtTime($1, $2)", value, endTime)
	return a
}

// LinearRampToValueAtTime fn
func (*AudioParam) LinearRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.linearRampToValueAtTime($1, $2)", value, endTime)
	return a
}

// SetTargetAtTime fn
func (*AudioParam) SetTargetAtTime(target float32, startTime float32, timeConstant float32) (a *AudioParam) {
	js.Rewrite("$<.setTargetAtTime($1, $2, $3)", target, startTime, timeConstant)
	return a
}

// SetValueAtTime fn
func (*AudioParam) SetValueAtTime(value float32, startTime float32) (a *AudioParam) {
	js.Rewrite("$<.setValueAtTime($1, $2)", value, startTime)
	return a
}

// SetValueCurveAtTime fn
func (*AudioParam) SetValueCurveAtTime(values []float32, startTime float32, duration float32) (a *AudioParam) {
	js.Rewrite("$<.setValueCurveAtTime($1, $2, $3)", values, startTime, duration)
	return a
}

// DefaultValue prop
func (*AudioParam) DefaultValue() (defaultValue float32) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

// Value prop
func (*AudioParam) Value() (value float32) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*AudioParam) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
