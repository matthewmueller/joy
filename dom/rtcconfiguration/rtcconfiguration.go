package rtcconfiguration

import (
	"github.com/matthewmueller/joy/dom/rtcbundlepolicy"
	"github.com/matthewmueller/joy/dom/rtciceserver"
	"github.com/matthewmueller/joy/dom/rtcicetransportpolicy"
)

type RTCConfiguration struct {
	bundlePolicy       *rtcbundlepolicy.RTCBundlePolicy
	iceServers         *[]*rtciceserver.RTCIceServer
	iceTransportPolicy *rtcicetransportpolicy.RTCIceTransportPolicy
	peerIdentity       *string
}
