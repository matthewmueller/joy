package rtcstatsprovider

import (
	"github.com/matthewmueller/joy/dom/rtcstatsreport"
	"github.com/matthewmueller/joy/dom/window"
)

// RTCStatsProvider interface
// js:"RTCStatsProvider"
type RTCStatsProvider interface {
	window.EventTarget

	// GetStats
	// js:"getStats"
	// jsrewrite:"await $_.getStats()"
	GetStats() (r *rtcstatsreport.RTCStatsReport)

	// MsGetStats
	// js:"msGetStats"
	// jsrewrite:"await $_.msGetStats()"
	MsGetStats() (r *rtcstatsreport.RTCStatsReport)
}
