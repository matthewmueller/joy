package rtcicecandidateattributes

import (
	"github.com/matthewmueller/golly/dom/rtcstats"
	"github.com/matthewmueller/golly/dom/rtcstatsicecandidatetype"
)

type RTCIceCandidateAttributes struct {
	*rtcstats.RTCStats

	addressSourceUrl *string
	candidateType    *rtcstatsicecandidatetype.RTCStatsIceCandidateType
	ipAddress        *string
	portNumber       *int
	priority         *int
	transport        *string
}
