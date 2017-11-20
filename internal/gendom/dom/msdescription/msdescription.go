package msdescription

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msconnectivity"
	"github.com/matthewmueller/golly/internal/gendom/dom/msipaddressinfo"
	"github.com/matthewmueller/golly/internal/gendom/dom/msnetworkconnectivityinfo"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtciceprotocol"
)

type MSDescription struct {
	*RTCStats

	connectivity         *msconnectivity.MSConnectivity
	deviceDevName        *string
	localAddr            *msipaddressinfo.MSIPAddressInfo
	networkconnectivity  *msnetworkconnectivityinfo.MSNetworkConnectivityInfo
	reflexiveLocalIPAddr *msipaddressinfo.MSIPAddressInfo
	remoteAddr           *msipaddressinfo.MSIPAddressInfo
	transport            *rtciceprotocol.RTCIceProtocol
}
