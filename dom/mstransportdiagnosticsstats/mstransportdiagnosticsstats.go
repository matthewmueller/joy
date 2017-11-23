package mstransportdiagnosticsstats

import (
	"github.com/matthewmueller/golly/dom2/msiceaddrtype"
	"github.com/matthewmueller/golly/dom2/msicewarningflags"
	"github.com/matthewmueller/golly/dom2/msnetworkinterfacetype"
	"github.com/matthewmueller/golly/dom2/rtciceprotocol"
	"github.com/matthewmueller/golly/dom2/rtcicerole"
	"github.com/matthewmueller/golly/dom2/rtcstats"
)

type MSTransportDiagnosticsStats struct {
	*rtcstats.RTCStats

	allocationTimeInMs     *uint
	baseAddress            *string
	baseInterface          *msnetworkinterfacetype.MSNetworkInterfaceType
	iceRole                *rtcicerole.RTCIceRole
	iceWarningFlags        *msicewarningflags.MSIceWarningFlags
	interfaces             *msnetworkinterfacetype.MSNetworkInterfaceType
	localAddress           *string
	localAddrType          *msiceaddrtype.MSIceAddrType
	localInterface         *msnetworkinterfacetype.MSNetworkInterfaceType
	localMR                *string
	localMRTCPPort         *uint8
	localSite              *string
	msRtcEngineVersion     *string
	networkName            *string
	numConsentReqReceived  *uint
	numConsentReqSent      *uint
	numConsentRespReceived *uint
	numConsentRespSent     *uint
	portRangeMax           *uint8
	portRangeMin           *uint8
	protocol               *rtciceprotocol.RTCIceProtocol
	remoteAddress          *string
	remoteAddrType         *msiceaddrtype.MSIceAddrType
	remoteMR               *string
	remoteMRTCPPort        *uint8
	remoteSite             *string
	rtpRtcpMux             *bool
	stunVer                *uint
}
