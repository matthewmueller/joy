package rtcsessiondescription

import (
	"github.com/matthewmueller/golly/dom/rtcsdptype"
	"github.com/matthewmueller/golly/dom/rtcsessiondescriptioninit"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(descriptioninitdict *rtcsessiondescriptioninit.RTCSessionDescriptionInit) *RTCSessionDescription {
	js.Rewrite("RTCSessionDescription")
	return &RTCSessionDescription{}
}

// RTCSessionDescription struct
// js:"RTCSessionDescription,omit"
type RTCSessionDescription struct {
}

// ToJSON fn
func (*RTCSessionDescription) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// Sdp prop
func (*RTCSessionDescription) Sdp() (sdp string) {
	js.Rewrite("$<.sdp")
	return sdp
}

// Sdp prop
func (*RTCSessionDescription) SetSdp(sdp string) {
	js.Rewrite("$<.sdp = $1", sdp)
}

// Type prop
func (*RTCSessionDescription) Type() (kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*RTCSessionDescription) SetType(kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$<.type = $1", kind)
}
