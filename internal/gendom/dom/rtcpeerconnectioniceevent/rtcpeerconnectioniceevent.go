package rtcpeerconnectioniceevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidate"
	"github.com/matthewmueller/golly/js"
)

type RTCPeerConnectionIceEvent struct {
	*event.Event
}

func (*RTCPeerConnectionIceEvent) GetCandidate() (candidate *rtcicecandidate.RTCIceCandidate) {
	js.Rewrite("$<.candidate")
	return candidate
}
