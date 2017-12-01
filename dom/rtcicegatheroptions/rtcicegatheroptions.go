package rtcicegatheroptions

import (
	"github.com/matthewmueller/golly/dom/msportrange"
	"github.com/matthewmueller/golly/dom/rtcicegatherpolicy"
	"github.com/matthewmueller/golly/dom/rtciceserver"
)

type RTCIceGatherOptions struct {
	gatherPolicy *rtcicegatherpolicy.RTCIceGatherPolicy
	iceservers   *[]*rtciceserver.RTCIceServer
	portRange    *msportrange.MSPortRange
}
