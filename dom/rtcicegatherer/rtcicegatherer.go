package rtcicegatherer

import (
	"github.com/matthewmueller/golly/dom/rtcicecandidatedictionary"
	"github.com/matthewmueller/golly/dom/rtcicecomponent"
	"github.com/matthewmueller/golly/dom/rtcicegathererevent"
	"github.com/matthewmueller/golly/dom/rtciceparameters"
	"github.com/matthewmueller/golly/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCIceGatherer)(nil)
var _ window.EventTarget = (*RTCIceGatherer)(nil)

// RTCIceGatherer struct
// js:"RTCIceGatherer,omit"
type RTCIceGatherer struct {
}

// CreateAssociatedGatherer fn
// js:"createAssociatedGatherer"
func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	js.Rewrite("$_.createAssociatedGatherer()")
	return r
}

// GetLocalCandidates fn
// js:"getLocalCandidates"
func (*RTCIceGatherer) GetLocalCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$_.getLocalCandidates()")
	return r
}

// GetLocalParameters fn
// js:"getLocalParameters"
func (*RTCIceGatherer) GetLocalParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$_.getLocalParameters()")
	return r
}

// GetStats fn
// js:"getStats"
func (*RTCIceGatherer) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.getStats()")
	return r
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCIceGatherer) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.msGetStats()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCIceGatherer) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCIceGatherer) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCIceGatherer) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Component prop
// js:"component"
func (*RTCIceGatherer) Component() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$_.component")
	return component
}

// Onerror prop
// js:"onerror"
func (*RTCIceGatherer) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCIceGatherer) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onlocalcandidate prop
// js:"onlocalcandidate"
func (*RTCIceGatherer) Onlocalcandidate() (onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	js.Rewrite("$_.onlocalcandidate")
	return onlocalcandidate
}

// SetOnlocalcandidate prop
// js:"onlocalcandidate"
func (*RTCIceGatherer) SetOnlocalcandidate(onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	js.Rewrite("$_.onlocalcandidate = $1", onlocalcandidate)
}
