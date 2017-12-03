package mslocalclienteventbase

import "github.com/matthewmueller/joy/dom/rtcstats"

type MSLocalClientEventBase struct {
	*rtcstats.RTCStats

	networkBandwidthLowEventRatio   *float32
	networkReceiveQualityEventRatio *float32
}
