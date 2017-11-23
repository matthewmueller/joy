package rtcpeerconnectioniceevent

import (
	"github.com/matthewmueller/golly/dom2/rtcicecandidate"
	"github.com/matthewmueller/golly/dom2/rtcpeerconnectioniceeventinit"
	"github.com/matthewmueller/golly/dom2/window"
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
