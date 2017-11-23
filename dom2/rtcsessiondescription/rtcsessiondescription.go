package rtcsessiondescription

import (
	"github.com/matthewmueller/golly/dom2/rtcsdptype"
	"github.com/matthewmueller/golly/dom2/rtcsessiondescriptioninit"
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

// ToJSON
func (*RTCSessionDescription) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// Sdp
func (*RTCSessionDescription) Sdp() (sdp string) {
	js.Rewrite("$<.Sdp")
	return sdp
}

// Sdp
func (*RTCSessionDescription) SetSdp(sdp string) {
	js.Rewrite("$<.Sdp = $1", sdp)
}

// Type
func (*RTCSessionDescription) Type() (kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*RTCSessionDescription) SetType(kind *rtcsdptype.RTCSdpType) {
	js.Rewrite("$<.Type = $1", kind)
}
