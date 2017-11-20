package msvideorecvpayload

import "github.com/matthewmueller/golly/internal/gendom/dom/msvideoresolutiondistribution"

type MSVideoRecvPayload struct {
	*MSVideoPayload

	lowBitRateCallPercent                                *float32
	lowFrameRateCallPercent                              *float32
	recvBitRateAverage                                   *uint
	recvBitRateMaximum                                   *uint
	recvCodecType                                        *string
	recvFpsHarmonicAverage                               *float32
	recvFrameRateAverage                                 *float32
	recvNumResSwitches                                   *float32
	recvReorderBufferMaxSuccessfullyOrderedExtent        *uint
	recvReorderBufferMaxSuccessfullyOrderedLateTime      *uint
	recvReorderBufferPacketsDroppedDueToBufferExhaustion *uint
	recvReorderBufferPacketsDroppedDueToTimeout          *uint
	recvReorderBufferReorderedPackets                    *uint
	recvResolutionHeight                                 *uint
	recvResolutionWidth                                  *uint
	recvVideoStreamsMax                                  *uint
	recvVideoStreamsMin                                  *uint
	recvVideoStreamsMode                                 *int
	reorderBufferTotalPackets                            *uint
	videoFrameLossRate                                   *float32
	videoPostFECPLR                                      *float32
	videoResolutions                                     *msvideoresolutiondistribution.MSVideoResolutionDistribution
}
