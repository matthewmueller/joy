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

// ToJSON
func (*RTCIceCandidate) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// Candidate
func (*RTCIceCandidate) Candidate() (candidate string) {
	js.Rewrite("$<.Candidate")
	return candidate
}

// Candidate
func (*RTCIceCandidate) SetCandidate(candidate string) {
	js.Rewrite("$<.Candidate = $1", candidate)
}

// SdpMid
func (*RTCIceCandidate) SdpMid() (sdpMid string) {
	js.Rewrite("$<.SdpMid")
	return sdpMid
}

// SdpMid
func (*RTCIceCandidate) SetSdpMid(sdpMid string) {
	js.Rewrite("$<.SdpMid = $1", sdpMid)
}

// SdpMLineIndex
func (*RTCIceCandidate) SdpMLineIndex() (sdpMLineIndex uint8) {
	js.Rewrite("$<.SdpMLineIndex")
	return sdpMLineIndex
}

// SdpMLineIndex
func (*RTCIceCandidate) SetSdpMLineIndex(sdpMLineIndex uint8) {
	js.Rewrite("$<.SdpMLineIndex = $1", sdpMLineIndex)
}
