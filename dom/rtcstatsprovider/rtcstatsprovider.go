package rtcstatsprovider

import (
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
)

// RTCStatsProvider interface
// js:"RTCStatsProvider"
type RTCStatsProvider interface {
	window.EventTarget

	// GetStats
	// js:"getStats"
	GetStats() (r *rtcstatsreport.RTCStatsReport)

	// MsGetStats
	// js:"msGetStats"
	MsGetStats() (r *rtcstatsreport.RTCStatsReport)
}
