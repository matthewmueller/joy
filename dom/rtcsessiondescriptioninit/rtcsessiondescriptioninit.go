package rtcsessiondescriptioninit

import "github.com/matthewmueller/joy/dom/rtcsdptype"

type RTCSessionDescriptionInit struct {
	sdp  *string
	kind *rtcsdptype.RTCSdpType
}
