package rtcstats

import (
	"github.com/matthewmueller/golly/dom2/msstatstype"
	"github.com/matthewmueller/golly/dom2/rtcstatstype"
)

type RTCStats struct {
	id        *string
	msType    *msstatstype.MSStatsType
	timestamp *int
	kind      *rtcstatstype.RTCStatsType
}
