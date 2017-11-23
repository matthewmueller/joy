package rtcoutboundrtpstreamstats

import "github.com/matthewmueller/golly/dom2/rtcrtpstreamstats"

type RTCOutboundRTPStreamStats struct {
	*rtcrtpstreamstats.RTCRTPStreamStats

	bytesSent     *uint64
	packetsSent   *uint
	roundTripTime *float32
	targetBitrate *float32
}
