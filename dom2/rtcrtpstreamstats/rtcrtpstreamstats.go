package rtcrtpstreamstats

import "github.com/matthewmueller/golly/dom2/rtcstats"

type RTCRTPStreamStats struct {
	*rtcstats.RTCStats

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
