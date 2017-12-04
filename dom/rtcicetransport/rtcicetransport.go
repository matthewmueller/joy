package rtcicetransport

import (
	"github.com/matthewmueller/joy/dom/rtcicecandidatedictionary"
	"github.com/matthewmueller/joy/dom/rtcicecandidatepair"
	"github.com/matthewmueller/joy/dom/rtcicecandidatepairchangedevent"
	"github.com/matthewmueller/joy/dom/rtcicecomponent"
	"github.com/matthewmueller/joy/dom/rtcicegatherer"
	"github.com/matthewmueller/joy/dom/rtciceparameters"
	"github.com/matthewmueller/joy/dom/rtcicerole"
	"github.com/matthewmueller/joy/dom/rtcicetransportstate"
	"github.com/matthewmueller/joy/dom/rtcicetransportstatechangedevent"
	"github.com/matthewmueller/joy/dom/rtcstatsprovider"
	"github.com/matthewmueller/joy/dom/rtcstatsreport"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCIceTransport)(nil)
var _ window.EventTarget = (*RTCIceTransport)(nil)

// New fn
func New() *RTCIceTransport {
	macro.Rewrite("new RTCIceTransport()")
	return &RTCIceTransport{}
}

// RTCIceTransport struct
// js:"RTCIceTransport,omit"
type RTCIceTransport struct {
}

// AddRemoteCandidate fn
// js:"addRemoteCandidate"
func (*RTCIceTransport) AddRemoteCandidate(remoteCandidate interface{}) {
	macro.Rewrite("$_.addRemoteCandidate($1)", remoteCandidate)
}

// CreateAssociatedTransport fn
// js:"createAssociatedTransport"
func (*RTCIceTransport) CreateAssociatedTransport() (r *RTCIceTransport) {
	macro.Rewrite("$_.createAssociatedTransport()")
	return r
}

// GetNominatedCandidatePair fn
// js:"getNominatedCandidatePair"
func (*RTCIceTransport) GetNominatedCandidatePair() (r *rtcicecandidatepair.RTCIceCandidatePair) {
	macro.Rewrite("$_.getNominatedCandidatePair()")
	return r
}

// GetRemoteCandidates fn
// js:"getRemoteCandidates"
func (*RTCIceTransport) GetRemoteCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	macro.Rewrite("$_.getRemoteCandidates()")
	return r
}

// GetRemoteParameters fn
// js:"getRemoteParameters"
func (*RTCIceTransport) GetRemoteParameters() (r *rtciceparameters.RTCIceParameters) {
	macro.Rewrite("$_.getRemoteParameters()")
	return r
}

// SetRemoteCandidates fn
// js:"setRemoteCandidates"
func (*RTCIceTransport) SetRemoteCandidates(remoteCandidates []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	macro.Rewrite("$_.setRemoteCandidates($1)", remoteCandidates)
}

// Start fn
// js:"start"
func (*RTCIceTransport) Start(gatherer *rtcicegatherer.RTCIceGatherer, remoteParameters *rtciceparameters.RTCIceParameters, role *rtcicerole.RTCIceRole) {
	macro.Rewrite("$_.start($1, $2, $3)", gatherer, remoteParameters, role)
}

// Stop fn
// js:"stop"
func (*RTCIceTransport) Stop() {
	macro.Rewrite("$_.stop()")
}

// GetStats fn
// js:"getStats"
func (*RTCIceTransport) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCIceTransport) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCIceTransport) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCIceTransport) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCIceTransport) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Component prop
// js:"component"
func (*RTCIceTransport) Component() (component *rtcicecomponent.RTCIceComponent) {
	macro.Rewrite("$_.component")
	return component
}

// IceGatherer prop
// js:"iceGatherer"
func (*RTCIceTransport) IceGatherer() (iceGatherer *rtcicegatherer.RTCIceGatherer) {
	macro.Rewrite("$_.iceGatherer")
	return iceGatherer
}

// Oncandidatepairchange prop
// js:"oncandidatepairchange"
func (*RTCIceTransport) Oncandidatepairchange() (oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	macro.Rewrite("$_.oncandidatepairchange")
	return oncandidatepairchange
}

// SetOncandidatepairchange prop
// js:"oncandidatepairchange"
func (*RTCIceTransport) SetOncandidatepairchange(oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	macro.Rewrite("$_.oncandidatepairchange = $1", oncandidatepairchange)
}

// Onicestatechange prop
// js:"onicestatechange"
func (*RTCIceTransport) Onicestatechange() (onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	macro.Rewrite("$_.onicestatechange")
	return onicestatechange
}

// SetOnicestatechange prop
// js:"onicestatechange"
func (*RTCIceTransport) SetOnicestatechange(onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	macro.Rewrite("$_.onicestatechange = $1", onicestatechange)
}

// Role prop
// js:"role"
func (*RTCIceTransport) Role() (role *rtcicerole.RTCIceRole) {
	macro.Rewrite("$_.role")
	return role
}

// State prop
// js:"state"
func (*RTCIceTransport) State() (state *rtcicetransportstate.RTCIceTransportState) {
	macro.Rewrite("$_.state")
	return state
}
