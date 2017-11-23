package rtcinboundrtpstreamstats

import "github.com/matthewmueller/golly/dom2/rtcrtpstreamstats"

type RTCInboundRTPStreamStats struct {
	*rtcrtpstreamstats.RTCRTPStreamStats

	bytesReceived   *uint64
	fractionLost    *float32
	jitter          *float32
	packetsLost     *uint
	packetsReceived *uint
}
