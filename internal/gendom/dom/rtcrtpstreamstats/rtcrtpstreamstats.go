package rtcrtpstreamstats

type RTCRTPStreamStats struct {
	*RTCStats

	associateStatsId *string
	codecId          *string
	firCount         *uint
	isRemote         *bool
	mediaTrackId     *string
	nackCount        *uint
	pliCount         *uint
	sliCount         *uint
	ssrc             *string
	transportId      *string
}
