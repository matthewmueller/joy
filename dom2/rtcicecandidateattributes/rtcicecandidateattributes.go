package rtcicecandidateattributes

import "github.com/matthewmueller/golly/dom2/rtcstats"

type RTCIceCandidateAttributes struct {
	*rtcstats.RTCStats

	addressSourceUrl *string
	candidateType    *rtcstatsicecandidatetype.RTCStatsIceCandidateType
	ipAddress        *string
	portNumber       *int
	priority         *int
	transport        *string
}
