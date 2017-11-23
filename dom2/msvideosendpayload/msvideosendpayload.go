package msvideosendpayload

import "github.com/matthewmueller/golly/dom2/msvideopayload"

type MSVideoSendPayload struct {
	*msvideopayload.MSVideoPayload

	sendBitRateAverage   *uint64
	sendBitRateMaximum   *uint64
	sendFrameRateAverage *float32
	sendResolutionHeight *uint
	sendResolutionWidth  *uint
	sendVideoStreamsMax  *uint
}
