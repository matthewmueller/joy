package rtcstats

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msstatstype"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatstype"
)

type RTCStats struct {
	id        *string
	msType    *msstatstype.MSStatsType
	timestamp *int
	kind      *rtcstatstype.RTCStatsType
}
