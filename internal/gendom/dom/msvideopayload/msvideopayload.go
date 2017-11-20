package msvideopayload

type MSVideoPayload struct {
	*MSPayloadBase

	durationSeconds     *float32
	resolution          *string
	videoBitRateAvg     *uint
	videoBitRateMax     *uint
	videoFrameRateAvg   *float32
	videoPacketLossRate *float32
}
