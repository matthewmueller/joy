package msnetwork

import (
	"github.com/matthewmueller/golly/dom/msdelay"
	"github.com/matthewmueller/golly/dom/msjitter"
	"github.com/matthewmueller/golly/dom/mspacketloss"
	"github.com/matthewmueller/golly/dom/msutilization"
	"github.com/matthewmueller/golly/dom/rtcstats"
)

type MSNetwork struct {
	*rtcstats.RTCStats

	delay       *msdelay.MSDelay
	jitter      *msjitter.MSJitter
	packetLoss  *mspacketloss.MSPacketLoss
	utilization *msutilization.MSUtilization
}
