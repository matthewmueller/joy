package rtcpeerconnectioniceeventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidate"

type RTCPeerConnectionIceEventInit struct {
	*EventInit

	candidate *rtcicecandidate.RTCIceCandidate
}
