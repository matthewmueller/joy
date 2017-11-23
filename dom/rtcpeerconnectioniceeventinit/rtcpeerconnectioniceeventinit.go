package rtcpeerconnectioniceeventinit

import (
	"github.com/matthewmueller/golly/dom/eventinit"
	"github.com/matthewmueller/golly/dom/rtcicecandidate"
)

type RTCPeerConnectionIceEventInit struct {
	*eventinit.EventInit

	candidate *rtcicecandidate.RTCIceCandidate
}
