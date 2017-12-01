package rtcsessiondescriptioninit

import "github.com/matthewmueller/golly/dom/rtcsdptype"

type RTCSessionDescriptionInit struct {
	sdp  *string
	kind *rtcsdptype.RTCSdpType
}
