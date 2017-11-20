package rtcicecandidateattributes

import "github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsicecandidatetype"

type RTCIceCandidateAttributes struct {
	*RTCStats

	addressSourceUrl *string
	candidateType    *rtcstatsicecandidatetype.RTCStatsIceCandidateType
	ipAddress        *string
	portNumber       *int
	priority         *int
	transport        *string
}
