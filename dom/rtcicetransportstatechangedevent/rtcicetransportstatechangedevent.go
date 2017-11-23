package rtcicetransportstatechangedevent

import (
	"github.com/matthewmueller/golly/dom/rtcicetransportstate"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// RTCIceTransportStateChangedEvent struct
// js:"RTCIceTransportStateChangedEvent,omit"
type RTCIceTransportStateChangedEvent struct {
	window.Event
}

// State prop
func (*RTCIceTransportStateChangedEvent) State() (state *rtcicetransportstate.RTCIceTransportState) {
	js.Rewrite("$<.state")
	return state
}
