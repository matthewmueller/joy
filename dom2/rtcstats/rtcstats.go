package rtcstats

type RTCStats struct {
	id        *string
	msType    *msstatstype.MSStatsType
	timestamp *int
	kind      *rtcstatstype.RTCStatsType
}
