package msnetwork

import (
	"github.com/matthewmueller/joy/dom/msdelay"
	"github.com/matthewmueller/joy/dom/msjitter"
	"github.com/matthewmueller/joy/dom/mspacketloss"
	"github.com/matthewmueller/joy/dom/msutilization"
	"github.com/matthewmueller/joy/dom/rtcstats"
)

type MSNetwork struct {
	*rtcstats.RTCStats

	delay       *msdelay.MSDelay
	jitter      *msjitter.MSJitter
	packetLoss  *mspacketloss.MSPacketLoss
	utilization *msutilization.MSUtilization
}
