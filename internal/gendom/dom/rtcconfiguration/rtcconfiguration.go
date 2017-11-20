package rtcconfiguration

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcbundlepolicy"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtciceserver"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicetransportpolicy"
)

type RTCConfiguration struct {
	bundlePolicy       *rtcbundlepolicy.RTCBundlePolicy
	iceServers         *[]*rtciceserver.RTCIceServer
	iceTransportPolicy *rtcicetransportpolicy.RTCIceTransportPolicy
	peerIdentity       *string
}
