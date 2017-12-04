package rtcicegatherer

import (
	"github.com/matthewmueller/joy/dom/rtcicecandidatedictionary"
	"github.com/matthewmueller/joy/dom/rtcicecomponent"
	"github.com/matthewmueller/joy/dom/rtcicegathererevent"
	"github.com/matthewmueller/joy/dom/rtcicegatheroptions"
	"github.com/matthewmueller/joy/dom/rtciceparameters"
	"github.com/matthewmueller/joy/dom/rtcstatsprovider"
	"github.com/matthewmueller/joy/dom/rtcstatsreport"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCIceGatherer)(nil)
var _ window.EventTarget = (*RTCIceGatherer)(nil)

// New fn
func New(options *rtcicegatheroptions.RTCIceGatherOptions) *RTCIceGatherer {
	macro.Rewrite("new RTCIceGatherer($1)", options)
	return &RTCIceGatherer{}
}

// RTCIceGatherer struct
// js:"RTCIceGatherer,omit"
type RTCIceGatherer struct {
}

// CreateAssociatedGatherer fn
// js:"createAssociatedGatherer"
func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	macro.Rewrite("$_.createAssociatedGatherer()")
	return r
}

// GetLocalCandidates fn
// js:"getLocalCandidates"
func (*RTCIceGatherer) GetLocalCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	macro.Rewrite("$_.getLocalCandidates()")
	return r
}

// GetLocalParameters fn
// js:"getLocalParameters"
func (*RTCIceGatherer) GetLocalParameters() (r *rtciceparameters.RTCIceParameters) {
	macro.Rewrite("$_.getLocalParameters()")
	return r
}

// GetStats fn
// js:"getStats"
func (*RTCIceGatherer) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCIceGatherer) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCIceGatherer) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCIceGatherer) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCIceGatherer) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Component prop
// js:"component"
func (*RTCIceGatherer) Component() (component *rtcicecomponent.RTCIceComponent) {
	macro.Rewrite("$_.component")
	return component
}

// Onerror prop
// js:"onerror"
func (*RTCIceGatherer) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCIceGatherer) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onlocalcandidate prop
// js:"onlocalcandidate"
func (*RTCIceGatherer) Onlocalcandidate() (onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	macro.Rewrite("$_.onlocalcandidate")
	return onlocalcandidate
}

// SetOnlocalcandidate prop
// js:"onlocalcandidate"
func (*RTCIceGatherer) SetOnlocalcandidate(onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	macro.Rewrite("$_.onlocalcandidate = $1", onlocalcandidate)
}
