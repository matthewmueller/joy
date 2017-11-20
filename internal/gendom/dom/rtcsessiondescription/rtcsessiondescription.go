package rtcsessiondescription

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcsdptype"
	"github.com/matthewmueller/golly/js"
)

type RTCSessionDescription struct {
}

func (*RTCSessionDescription) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*RTCSessionDescription) GetSdp() (sdp string) {
	js.Rewrite("$<.sdp")
	return sdp
}

func (*RTCSessionDescription) SetSdp(sdp string) {
	js.Rewrite("$<.sdp = $1", sdp)
}

func (*RTCSessionDescription) GetType() (kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$<.type")
	return kind
}

func (*RTCSessionDescription) SetType(kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$<.type = $1", kind)
}
