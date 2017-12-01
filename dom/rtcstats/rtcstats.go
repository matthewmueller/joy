package rtcstats

import (
	"github.com/matthewmueller/golly/dom/msstatstype"
	"github.com/matthewmueller/golly/dom/rtcstatstype"
)

type RTCStats struct {
	id        *string
	msType    *msstatstype.MSStatsType
	timestamp *int
	kind      *rtcstatstype.RTCStatsType
}
