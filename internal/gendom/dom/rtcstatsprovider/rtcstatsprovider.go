package rtcstatsprovider

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/js"
)

type RTCStatsProvider struct {
	*eventtarget.EventTarget
}

func (*RTCStatsProvider) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $<.getStats()")
	return r
}

func (*RTCStatsProvider) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $<.msGetStats()")
	return r
}
