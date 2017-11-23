package rtcsessiondescriptioninit

import "github.com/matthewmueller/golly/dom2/rtcsdptype"

type RTCSessionDescriptionInit struct {
	sdp  *string
	kind *rtcsdptype.RTCSdpType
}
