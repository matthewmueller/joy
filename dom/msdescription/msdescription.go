package msdescription

import (
	"github.com/matthewmueller/golly/dom/msconnectivity"
	"github.com/matthewmueller/golly/dom/msipaddressinfo"
	"github.com/matthewmueller/golly/dom/msnetworkconnectivityinfo"
	"github.com/matthewmueller/golly/dom/rtciceprotocol"
	"github.com/matthewmueller/golly/dom/rtcstats"
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
