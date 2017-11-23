package rtcconfiguration

type RTCConfiguration struct {
	bundlePolicy       *rtcbundlepolicy.RTCBundlePolicy
	iceServers         *[]*rtciceserver.RTCIceServer
	iceTransportPolicy *rtcicetransportpolicy.RTCIceTransportPolicy
	peerIdentity       *string
}
