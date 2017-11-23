package rtcicegatherer

import (
	"github.com/matthewmueller/golly/dom2/rtcicecandidatedictionary"
	"github.com/matthewmueller/golly/dom2/rtcicecomponent"
	"github.com/matthewmueller/golly/dom2/rtcicegathererevent"
	"github.com/matthewmueller/golly/dom2/rtcicegatheroptions"
	"github.com/matthewmueller/golly/dom2/rtciceparameters"
	"github.com/matthewmueller/golly/dom2/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(options *rtcicegatheroptions.RTCIceGatherOptions) *RTCIceGatherer {
	js.Rewrite("RTCIceGatherer")
	return &RTCIceGatherer{}
}

// RTCIceGatherer struct
// js:"RTCIceGatherer,omit"
type RTCIceGatherer struct {
	rtcstatsprovider.RTCStatsProvider
}

// CreateAssociatedGatherer fn
func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	js.Rewrite("$<.createAssociatedGatherer()")
	return r
}

// GetLocalCandidates fn
func (*RTCIceGatherer) GetLocalCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.getLocalCandidates()")
	return r
}

// GetLocalParameters fn
func (*RTCIceGatherer) GetLocalParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

// Component prop
func (*RTCIceGatherer) Component() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$<.component")
	return component
}

// Onerror prop
func (*RTCIceGatherer) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*RTCIceGatherer) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onlocalcandidate prop
func (*RTCIceGatherer) Onlocalcandidate() (onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	js.Rewrite("$<.onlocalcandidate")
	return onlocalcandidate
}

// Onlocalcandidate prop
func (*RTCIceGatherer) SetOnlocalcandidate(onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	js.Rewrite("$<.onlocalcandidate = $1", onlocalcandidate)
}
