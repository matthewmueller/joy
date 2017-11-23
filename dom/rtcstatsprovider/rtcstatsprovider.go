package rtcstatsprovider

import (
	"github.com/matthewmueller/golly/dom2/rtcstatsreport"
	"github.com/matthewmueller/golly/dom2/window"
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
