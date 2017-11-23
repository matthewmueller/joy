package webkitrtcpeerconnection

import (
	"github.com/matthewmueller/golly/dom/rtcconfiguration"
	"github.com/matthewmueller/golly/dom/rtcpeerconnection"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(configuration *rtcconfiguration.RTCConfiguration) *WebkitRTCPeerConnection {
	js.Rewrite("webkitRTCPeerConnection")
	return &WebkitRTCPeerConnection{}
}

// WebkitRTCPeerConnection struct
// js:"WebkitRTCPeerConnection,omit"
type WebkitRTCPeerConnection struct {
	rtcpeerconnection.RTCPeerConnection
}
