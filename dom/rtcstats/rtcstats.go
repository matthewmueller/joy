package rtcstats

import (
	"github.com/matthewmueller/joy/dom/msstatstype"
	"github.com/matthewmueller/joy/dom/rtcstatstype"
)

type RTCStats struct {
	id        *string
	msType    *msstatstype.MSStatsType
	timestamp *int
	kind      *rtcstatstype.RTCStatsType
}
