package rtcconfiguration

import (
	"github.com/matthewmueller/golly/dom/rtcbundlepolicy"
	"github.com/matthewmueller/golly/dom/rtciceserver"
	"github.com/matthewmueller/golly/dom/rtcicetransportpolicy"
)

type RTCConfiguration struct {
	bundlePolicy       *rtcbundlepolicy.RTCBundlePolicy
	iceServers         *[]*rtciceserver.RTCIceServer
	iceTransportPolicy *rtcicetransportpolicy.RTCIceTransportPolicy
	peerIdentity       *string
}
