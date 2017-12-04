package rtcsessiondescription

import (
	"github.com/matthewmueller/joy/dom/rtcsdptype"
	"github.com/matthewmueller/joy/dom/rtcsessiondescriptioninit"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New(descriptioninitdict *rtcsessiondescriptioninit.RTCSessionDescriptionInit) *RTCSessionDescription {
	macro.Rewrite("new RTCSessionDescription($1)", descriptioninitdict)
	return &RTCSessionDescription{}
}

// RTCSessionDescription struct
// js:"RTCSessionDescription,omit"
type RTCSessionDescription struct {
}

// ToJSON fn
// js:"toJSON"
func (*RTCSessionDescription) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

// Sdp prop
// js:"sdp"
func (*RTCSessionDescription) Sdp() (sdp string) {
	macro.Rewrite("$_.sdp")
	return sdp
}

// SetSdp prop
// js:"sdp"
func (*RTCSessionDescription) SetSdp(sdp string) {
	macro.Rewrite("$_.sdp = $1", sdp)
}

// Type prop
// js:"type"
func (*RTCSessionDescription) Type() (kind *rtcsdptype.RTCSdpType) {
	macro.Rewrite("$_.type")
	return kind
}

// SetType prop
// js:"type"
func (*RTCSessionDescription) SetType(kind *rtcsdptype.RTCSdpType) {
	macro.Rewrite("$_.type = $1", kind)
}
