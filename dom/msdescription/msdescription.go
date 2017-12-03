package msdescription

import (
	"github.com/matthewmueller/joy/dom/msconnectivity"
	"github.com/matthewmueller/joy/dom/msipaddressinfo"
	"github.com/matthewmueller/joy/dom/msnetworkconnectivityinfo"
	"github.com/matthewmueller/joy/dom/rtciceprotocol"
	"github.com/matthewmueller/joy/dom/rtcstats"
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
