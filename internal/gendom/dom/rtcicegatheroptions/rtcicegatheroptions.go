package rtcicegatheroptions

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msportrange"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicegatherpolicy"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtciceserver"
)

type RTCIceGatherOptions struct {
	gatherPolicy *rtcicegatherpolicy.RTCIceGatherPolicy
	iceservers   *[]*rtciceserver.RTCIceServer
	portRange    *msportrange.MSPortRange
}
