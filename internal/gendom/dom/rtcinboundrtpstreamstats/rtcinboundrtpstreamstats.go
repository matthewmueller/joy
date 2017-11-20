package rtcinboundrtpstreamstats

type RTCInboundRTPStreamStats struct {
	*RTCRTPStreamStats

	bytesReceived   *uint64
	fractionLost    *float32
	jitter          *float32
	packetsLost     *uint
	packetsReceived *uint
}
