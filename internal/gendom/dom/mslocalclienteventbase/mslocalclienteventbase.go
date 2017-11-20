package mslocalclienteventbase

type MSLocalClientEventBase struct {
	*RTCStats

	networkBandwidthLowEventRatio   *float32
	networkReceiveQualityEventRatio *float32
}
