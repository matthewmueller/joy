package rtcdtlstransportstatechangedevent

import (
	"github.com/matthewmueller/golly/dom2/rtcdtlstransportstate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// RTCDtlsTransportStateChangedEvent struct
// js:"RTCDtlsTransportStateChangedEvent,omit"
type RTCDtlsTransportStateChangedEvent struct {
	window.Event
}

// State
func (*RTCDtlsTransportStateChangedEvent) State() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$<.State")
	return state
}
