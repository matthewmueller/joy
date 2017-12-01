package rtcicecandidatepairstats

import (
	"github.com/matthewmueller/golly/dom/rtcstats"
	"github.com/matthewmueller/golly/dom/rtcstatsicecandidatepairstate"
)

type RTCIceCandidatePairStats struct {
	*rtcstats.RTCStats

	availableIncomingBitrate *float32
	availableOutgoingBitrate *float32
	bytesReceived            *uint64
	bytesSent                *uint64
	localCandidateId         *string
	nominated                *bool
	priority                 *uint64
	readable                 *bool
	remoteCandidateId        *string
	roundTripTime            *float32
	state                    *rtcstatsicecandidatepairstate.RTCStatsIceCandidatePairState
	transportId              *string
	writable                 *bool
}
