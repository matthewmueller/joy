package rtcicetransportstatechangedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicetransportstate"
	"github.com/matthewmueller/golly/js"
)

type RTCIceTransportStateChangedEvent struct {
	*event.Event
}

func (*RTCIceTransportStateChangedEvent) GetState() (state *rtcicetransportstate.RTCIceTransportState) {
	js.Rewrite("$<.state")
	return state
}
