package rtcicegatheroptions

import (
	"github.com/matthewmueller/joy/dom/msportrange"
	"github.com/matthewmueller/joy/dom/rtcicegatherpolicy"
	"github.com/matthewmueller/joy/dom/rtciceserver"
)

type RTCIceGatherOptions struct {
	gatherPolicy *rtcicegatherpolicy.RTCIceGatherPolicy
	iceservers   *[]*rtciceserver.RTCIceServer
	portRange    *msportrange.MSPortRange
}
