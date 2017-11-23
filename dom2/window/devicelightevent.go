package window

import "github.com/matthewmueller/golly/js"

// js:"DeviceLightEvent,omit"
type DeviceLightEvent struct {
	Event
}

// Value
func (*DeviceLightEvent) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}
