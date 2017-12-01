package rtcicecandidate

import (
	"github.com/matthewmueller/golly/dom/rtcicecandidateinit"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(candidateinitdict *rtcicecandidateinit.RTCIceCandidateInit) *RTCIceCandidate {
	js.Rewrite("RTCIceCandidate")
	return &RTCIceCandidate{}
}

// RTCIceCandidate struct
// js:"RTCIceCandidate,omit"
type RTCIceCandidate struct {
}

// ToJSON fn
// js:"toJSON"
func (*RTCIceCandidate) ToJSON() (i interface{}) {
	js.Rewrite("$_.toJSON()")
	return i
}

// Candidate prop
// js:"candidate"
func (*RTCIceCandidate) Candidate() (candidate string) {
	js.Rewrite("$_.candidate")
	return candidate
}

// SetCandidate prop
// js:"candidate"
func (*RTCIceCandidate) SetCandidate(candidate string) {
	js.Rewrite("$_.candidate = $1", candidate)
}

// SdpMid prop
// js:"sdpMid"
func (*RTCIceCandidate) SdpMid() (sdpMid string) {
	js.Rewrite("$_.sdpMid")
	return sdpMid
}

// SetSdpMid prop
// js:"sdpMid"
func (*RTCIceCandidate) SetSdpMid(sdpMid string) {
	js.Rewrite("$_.sdpMid = $1", sdpMid)
}

// SdpMLineIndex prop
// js:"sdpMLineIndex"
func (*RTCIceCandidate) SdpMLineIndex() (sdpMLineIndex uint8) {
	js.Rewrite("$_.sdpMLineIndex")
	return sdpMLineIndex
}

// SetSdpMLineIndex prop
// js:"sdpMLineIndex"
func (*RTCIceCandidate) SetSdpMLineIndex(sdpMLineIndex uint8) {
	js.Rewrite("$_.sdpMLineIndex = $1", sdpMLineIndex)
}
