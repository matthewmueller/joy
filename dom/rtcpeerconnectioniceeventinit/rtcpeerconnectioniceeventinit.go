package rtcpeerconnectioniceeventinit

import (
	"github.com/matthewmueller/joy/dom/eventinit"
	"github.com/matthewmueller/joy/dom/rtcicecandidate"
)

type RTCPeerConnectionIceEventInit struct {
	*eventinit.EventInit

	candidate *rtcicecandidate.RTCIceCandidate
}
