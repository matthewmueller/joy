package msaudiorecvpayload

import (
	"github.com/matthewmueller/joy/dom/msaudiorecvsignal"
	"github.com/matthewmueller/joy/dom/mspayloadbase"
)

type MSAudioRecvPayload struct {
	*mspayloadbase.MSPayloadBase

	burstLossLength1          *float32
	burstLossLength2          *float32
	burstLossLength3          *float32
	burstLossLength4          *float32
	burstLossLength5          *float32
	burstLossLength6          *float32
	burstLossLength7          *float32
	burstLossLength8OrHigher  *float32
	fecRecvDistance1          *float32
	fecRecvDistance2          *float32
	fecRecvDistance3          *float32
	packetReorderDepthAvg     *int
	packetReorderDepthMax     *int
	packetReorderRatio        *float32
	ratioCompressedSamplesAvg *float32
	ratioConcealedSamplesAvg  *float32
	ratioStretchedSamplesAvg  *float32
	samplingRate              *uint
	signal                    *msaudiorecvsignal.MSAudioRecvSignal
}
