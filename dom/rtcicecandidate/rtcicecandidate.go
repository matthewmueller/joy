package rtcicecandidate

import (
	"github.com/matthewmueller/golly/dom2/rtcicecandidateinit"
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
func (*RTCIceCandidate) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// Candidate prop
func (*RTCIceCandidate) Candidate() (candidate string) {
	js.Rewrite("$<.candidate")
	return candidate
}

// Candidate prop
func (*RTCIceCandidate) SetCandidate(candidate string) {
	js.Rewrite("$<.candidate = $1", candidate)
}

// SdpMid prop
func (*RTCIceCandidate) SdpMid() (sdpMid string) {
	js.Rewrite("$<.sdpMid")
	return sdpMid
}

// SdpMid prop
func (*RTCIceCandidate) SetSdpMid(sdpMid string) {
	js.Rewrite("$<.sdpMid = $1", sdpMid)
}

// SdpMLineIndex prop
func (*RTCIceCandidate) SdpMLineIndex() (sdpMLineIndex uint8) {
	js.Rewrite("$<.sdpMLineIndex")
	return sdpMLineIndex
}

// SdpMLineIndex prop
func (*RTCIceCandidate) SetSdpMLineIndex(sdpMLineIndex uint8) {
	js.Rewrite("$<.sdpMLineIndex = $1", sdpMLineIndex)
}
