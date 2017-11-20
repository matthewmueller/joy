package rtcicecandidate

import "github.com/matthewmueller/golly/js"

type RTCIceCandidate struct {
}

func (*RTCIceCandidate) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*RTCIceCandidate) GetCandidate() (candidate string) {
	js.Rewrite("$<.candidate")
	return candidate
}

func (*RTCIceCandidate) SetCandidate(candidate string) {
	js.Rewrite("$<.candidate = $1", candidate)
}

func (*RTCIceCandidate) GetSdpMid() (sdpMid string) {
	js.Rewrite("$<.sdpMid")
	return sdpMid
}

func (*RTCIceCandidate) SetSdpMid(sdpMid string) {
	js.Rewrite("$<.sdpMid = $1", sdpMid)
}

func (*RTCIceCandidate) GetSdpMLineIndex() (sdpMLineIndex uint8) {
	js.Rewrite("$<.sdpMLineIndex")
	return sdpMLineIndex
}

func (*RTCIceCandidate) SetSdpMLineIndex(sdpMLineIndex uint8) {
	js.Rewrite("$<.sdpMLineIndex = $1", sdpMLineIndex)
}
