package rtcpeerconnectioniceevent

import (
	"github.com/matthewmueller/golly/dom/rtcicecandidate"
	"github.com/matthewmueller/golly/dom/rtcpeerconnectioniceeventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(kind string, eventinitdict *rtcpeerconnectioniceeventinit.RTCPeerConnectionIceEventInit) *RTCPeerConnectionIceEvent {
	js.Rewrite("RTCPeerConnectionIceEvent")
	return &RTCPeerConnectionIceEvent{}
}

// RTCPeerConnectionIceEvent struct
// js:"RTCPeerConnectionIceEvent,omit"
type RTCPeerConnectionIceEvent struct {
	window.Event
}

// Candidate prop
func (*RTCPeerConnectionIceEvent) Candidate() (candidate *rtcicecandidate.RTCIceCandidate) {
	js.Rewrite("$<.candidate")
	return candidate
}
