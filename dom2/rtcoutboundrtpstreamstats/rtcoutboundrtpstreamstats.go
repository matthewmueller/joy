package rtcoutboundrtpstreamstats

type RTCOutboundRTPStreamStats struct {
	*rtcrtpstreamstats.RTCRTPStreamStats

	bytesSent     *uint64
	packetsSent   *uint
	roundTripTime *float32
	targetBitrate *float32
}
