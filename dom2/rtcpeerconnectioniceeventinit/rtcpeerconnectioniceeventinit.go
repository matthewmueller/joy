package rtcpeerconnectioniceeventinit

import (
	"github.com/matthewmueller/golly/dom2/eventinit"
	"github.com/matthewmueller/golly/dom2/rtcicecandidate"
)

type RTCPeerConnectionIceEventInit struct {
	*eventinit.EventInit

	candidate *rtcicecandidate.RTCIceCandidate
}
