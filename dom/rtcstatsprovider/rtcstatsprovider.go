package rtcstatsprovider

import (
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCStatsProvider,omit"
type RTCStatsProvider interface {
	window.EventTarget

	// GetStats
	// js:"getStats,rewrite=getstats"
	GetStats() (r *rtcstatsreport.RTCStatsReport)

	// MsGetStats
	// js:"msGetStats,rewrite=msgetstats"
	MsGetStats() (r *rtcstatsreport.RTCStatsReport)
}

// getstats fn
func getstats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $<.getStats()")
	return r
}

// msgetstats fn
func msgetstats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $<.msGetStats()")
	return r
}
