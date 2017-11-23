package msdescription

import (
	"github.com/matthewmueller/golly/dom2/msconnectivity"
	"github.com/matthewmueller/golly/dom2/rtcstats"
)

type MSDescription struct {
	*rtcstats.RTCStats

	connectivity         *msconnectivity.MSConnectivity
	deviceDevName        *string
	localAddr            *msipaddressinfo.MSIPAddressInfo
	networkconnectivity  *msnetworkconnectivityinfo.MSNetworkConnectivityInfo
	reflexiveLocalIPAddr *msipaddressinfo.MSIPAddressInfo
	remoteAddr           *msipaddressinfo.MSIPAddressInfo
	transport            *rtciceprotocol.RTCIceProtocol
}
