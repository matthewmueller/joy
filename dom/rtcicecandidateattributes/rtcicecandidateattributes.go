package rtcicecandidateattributes

import (
	"github.com/matthewmueller/joy/dom/rtcstats"
	"github.com/matthewmueller/joy/dom/rtcstatsicecandidatetype"
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
