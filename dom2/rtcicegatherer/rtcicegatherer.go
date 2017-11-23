package rtcicegatherer

import (
	"github.com/matthewmueller/golly/dom2/rtcicecandidatedictionary"
	"github.com/matthewmueller/golly/dom2/rtcicecomponent"
	"github.com/matthewmueller/golly/dom2/rtcicegathererevent"
	"github.com/matthewmueller/golly/dom2/rtciceparameters"
	"github.com/matthewmueller/golly/dom2/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCIceGatherer,omit"
type RTCIceGatherer struct {
	rtcstatsprovider.RTCStatsProvider
}

// CreateAssociatedGatherer
func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	js.Rewrite("$<.CreateAssociatedGatherer()")
	return r
}

// GetLocalCandidates
func (*RTCIceGatherer) GetLocalCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.GetLocalCandidates()")
	return r
}

// GetLocalParameters
func (*RTCIceGatherer) GetLocalParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$<.GetLocalParameters()")
	return r
}

// Component
func (*RTCIceGatherer) Component() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$<.Component")
	return component
}

// Onerror
func (*RTCIceGatherer) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*RTCIceGatherer) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onlocalcandidate
func (*RTCIceGatherer) Onlocalcandidate() (onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	js.Rewrite("$<.Onlocalcandidate")
	return onlocalcandidate
}

// Onlocalcandidate
func (*RTCIceGatherer) SetOnlocalcandidate(onlocalcandidate func(*rtcicegathererevent.RTCIceGathererEvent)) {
	js.Rewrite("$<.Onlocalcandidate = $1", onlocalcandidate)
}
