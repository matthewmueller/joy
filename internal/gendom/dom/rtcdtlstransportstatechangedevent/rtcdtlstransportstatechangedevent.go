package rtcdtlstransportstatechangedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlstransportstate"
	"github.com/matthewmueller/golly/js"
)

type RTCDtlsTransportStateChangedEvent struct {
	*event.Event
}

func (*RTCDtlsTransportStateChangedEvent) GetState() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$<.state")
	return state
}
