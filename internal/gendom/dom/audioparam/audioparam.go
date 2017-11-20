package audioparam

import "github.com/matthewmueller/golly/js"

type AudioParam struct {
}

func (*AudioParam) CancelScheduledValues(startTime float32) (a *AudioParam) {
	js.Rewrite("$<.cancelScheduledValues($1)", startTime)
	return a
}

func (*AudioParam) ExponentialRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.exponentialRampToValueAtTime($1, $2)", value, endTime)
	return a
}

func (*AudioParam) LinearRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.linearRampToValueAtTime($1, $2)", value, endTime)
	return a
}

func (*AudioParam) SetTargetAtTime(target float32, startTime float32, timeConstant float32) (a *AudioParam) {
	js.Rewrite("$<.setTargetAtTime($1, $2, $3)", target, startTime, timeConstant)
	return a
}

func (*AudioParam) SetValueAtTime(value float32, startTime float32) (a *AudioParam) {
	js.Rewrite("$<.setValueAtTime($1, $2)", value, startTime)
	return a
}

func (*AudioParam) SetValueCurveAtTime(values []float32, startTime float32, duration float32) (a *AudioParam) {
	js.Rewrite("$<.setValueCurveAtTime($1, $2, $3)", values, startTime, duration)
	return a
}

func (*AudioParam) GetDefaultValue() (defaultValue float32) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*AudioParam) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*AudioParam) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
