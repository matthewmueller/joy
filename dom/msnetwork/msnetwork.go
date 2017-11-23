package msnetwork

import (
	"github.com/matthewmueller/golly/dom2/msdelay"
	"github.com/matthewmueller/golly/dom2/msjitter"
	"github.com/matthewmueller/golly/dom2/mspacketloss"
	"github.com/matthewmueller/golly/dom2/msutilization"
	"github.com/matthewmueller/golly/dom2/rtcstats"
)

type MSNetwork struct {
	*rtcstats.RTCStats

	delay       *msdelay.MSDelay
	jitter      *msjitter.MSJitter
	packetLoss  *mspacketloss.MSPacketLoss
	utilization *msutilization.MSUtilization
}
