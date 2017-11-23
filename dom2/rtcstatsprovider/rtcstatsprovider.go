package rtcstatsprovider

import (
	"github.com/matthewmueller/golly/dom2/rtcstatsreport"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"RTCStatsProvider,omit"
type RTCStatsProvider interface {
	window.EventTarget

	// GetStats
	GetStats() (r *rtcstatsreport.RTCStatsReport)

	// MsGetStats
	MsGetStats() (r *rtcstatsreport.RTCStatsReport)
}
