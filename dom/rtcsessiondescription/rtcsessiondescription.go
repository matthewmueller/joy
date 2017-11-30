package rtcsessiondescription

import (
	"github.com/matthewmueller/golly/dom/rtcsdptype"
	"github.com/matthewmueller/golly/js"
)

// RTCSessionDescription struct
// js:"RTCSessionDescription,omit"
type RTCSessionDescription struct {
}

// ToJSON fn
// js:"toJSON"
func (*RTCSessionDescription) ToJSON() (i interface{}) {
	js.Rewrite("$_.toJSON()")
	return i
}

// Sdp prop
// js:"sdp"
func (*RTCSessionDescription) Sdp() (sdp string) {
	js.Rewrite("$_.sdp")
	return sdp
}

// SetSdp prop
// js:"sdp"
func (*RTCSessionDescription) SetSdp(sdp string) {
	js.Rewrite("$_.sdp = $1", sdp)
}

// Type prop
// js:"type"
func (*RTCSessionDescription) Type() (kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$_.type")
	return kind
}

// SetType prop
// js:"type"
func (*RTCSessionDescription) SetType(kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$_.type = $1", kind)
}
