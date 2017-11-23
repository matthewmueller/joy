package rtcpeerconnectioniceevent

import "github.com/matthewmueller/golly/js"

// js:"RTCPeerConnectionIceEvent,omit"
type RTCPeerConnectionIceEvent struct {
	window.Event
}

// Candidate
func (*RTCPeerConnectionIceEvent) Candidate() (candidate *rtcicecandidate.RTCIceCandidate) {
	js.Rewrite("$<.Candidate")
	return candidate
}
