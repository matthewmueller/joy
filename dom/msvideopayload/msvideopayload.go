package msvideopayload

import "github.com/matthewmueller/golly/dom/mspayloadbase"

type MSVideoPayload struct {
	*mspayloadbase.MSPayloadBase

	durationSeconds     *float32
	resolution          *string
	videoBitRateAvg     *uint
	videoBitRateMax     *uint
	videoFrameRateAvg   *float32
	videoPacketLossRate *float32
}
