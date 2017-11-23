package rtcdtlstransportstatechangedevent

import (
	"github.com/matthewmueller/golly/dom/rtcdtlstransportstate"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// RTCDtlsTransportStateChangedEvent struct
// js:"RTCDtlsTransportStateChangedEvent,omit"
type RTCDtlsTransportStateChangedEvent struct {
	window.Event
}

// State prop
func (*RTCDtlsTransportStateChangedEvent) State() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$<.state")
	return state
}
