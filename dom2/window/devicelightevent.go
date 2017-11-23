package window

import (
	"github.com/matthewmueller/golly/dom2/devicelighteventinit"
	"github.com/matthewmueller/golly/js"
)

// NewDeviceLightEvent fn
func NewDeviceLightEvent(typearg string, eventinitdict *devicelighteventinit.DeviceLightEventInit) *DeviceLightEvent {
	js.Rewrite("DeviceLightEvent")
	return &DeviceLightEvent{}
}

// DeviceLightEvent struct
// js:"DeviceLightEvent,omit"
type DeviceLightEvent struct {
	Event
}

// Value
func (*DeviceLightEvent) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}
