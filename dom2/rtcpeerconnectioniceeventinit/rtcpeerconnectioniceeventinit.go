package rtcpeerconnectioniceeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type RTCPeerConnectionIceEventInit struct {
	*eventinit.EventInit

	candidate *rtcicecandidate.RTCIceCandidate
}
