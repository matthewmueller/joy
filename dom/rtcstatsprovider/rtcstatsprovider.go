package rtcstatsprovider

import (
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
)

// js:"RTCStatsProvider,omit"
type RTCStatsProvider interface {
	window.EventTarget

	// GetStats
	// js:"getStats"
	GetStats() (r *rtcstatsreport.RTCStatsReport)

	// MsGetStats
	// js:"msGetStats"
	MsGetStats() (r *rtcstatsreport.RTCStatsReport)
}
