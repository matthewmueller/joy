package mslocalclienteventbase

import "github.com/matthewmueller/golly/dom/rtcstats"

type MSLocalClientEventBase struct {
	*rtcstats.RTCStats

	networkBandwidthLowEventRatio   *float32
	networkReceiveQualityEventRatio *float32
}
