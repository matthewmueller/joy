package rtcsessiondescriptioninit

import "github.com/matthewmueller/golly/internal/gendom/dom/rtcsdptype"

type RTCSessionDescriptionInit struct {
	sdp  *string
	kind *rtcsdptype.RTCSdpType
}
