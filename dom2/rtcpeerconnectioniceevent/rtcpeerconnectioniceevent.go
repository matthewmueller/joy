package rtcpeerconnectioniceevent

import (
	"github.com/matthewmueller/golly/dom2/rtcicecandidate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCPeerConnectionIceEvent,omit"
type RTCPeerConnectionIceEvent struct {
	window.Event
}

// Candidate
func (*RTCPeerConnectionIceEvent) Candidate() (candidate *rtcicecandidate.RTCIceCandidate) {
	js.Rewrite("$<.Candidate")
	return candidate
}
