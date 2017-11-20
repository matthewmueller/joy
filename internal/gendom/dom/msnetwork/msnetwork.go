package msnetwork

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msdelay"
	"github.com/matthewmueller/golly/internal/gendom/dom/msjitter"
	"github.com/matthewmueller/golly/internal/gendom/dom/mspacketloss"
	"github.com/matthewmueller/golly/internal/gendom/dom/msutilization"
)

type MSNetwork struct {
	*RTCStats

	delay       *msdelay.MSDelay
	jitter      *msjitter.MSJitter
	packetLoss  *mspacketloss.MSPacketLoss
	utilization *msutilization.MSUtilization
}
