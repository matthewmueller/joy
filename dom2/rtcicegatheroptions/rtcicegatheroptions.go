package rtcicegatheroptions

import (
	"github.com/matthewmueller/golly/dom2/msportrange"
	"github.com/matthewmueller/golly/dom2/rtcicegatherpolicy"
	"github.com/matthewmueller/golly/dom2/rtciceserver"
)

type RTCIceGatherOptions struct {
	gatherPolicy *rtcicegatherpolicy.RTCIceGatherPolicy
	iceservers   *[]*rtciceserver.RTCIceServer
	portRange    *msportrange.MSPortRange
}
