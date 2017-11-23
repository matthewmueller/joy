package rtcconfiguration

import (
	"github.com/matthewmueller/golly/dom2/rtcbundlepolicy"
	"github.com/matthewmueller/golly/dom2/rtciceserver"
)

type RTCConfiguration struct {
	bundlePolicy       *rtcbundlepolicy.RTCBundlePolicy
	iceServers         *[]*rtciceserver.RTCIceServer
	iceTransportPolicy *rtcicetransportpolicy.RTCIceTransportPolicy
	peerIdentity       *string
}
