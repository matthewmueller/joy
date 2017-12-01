package rtcicetransport

import (
	"github.com/matthewmueller/golly/dom/rtcicecandidatedictionary"
	"github.com/matthewmueller/golly/dom/rtcicecandidatepair"
	"github.com/matthewmueller/golly/dom/rtcicecandidatepairchangedevent"
	"github.com/matthewmueller/golly/dom/rtcicecomponent"
	"github.com/matthewmueller/golly/dom/rtcicegatherer"
	"github.com/matthewmueller/golly/dom/rtciceparameters"
	"github.com/matthewmueller/golly/dom/rtcicerole"
	"github.com/matthewmueller/golly/dom/rtcicetransportstate"
	"github.com/matthewmueller/golly/dom/rtcicetransportstatechangedevent"
	"github.com/matthewmueller/golly/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCIceTransport)(nil)
var _ window.EventTarget = (*RTCIceTransport)(nil)

// New fn
func New() *RTCIceTransport {
	js.Rewrite("RTCIceTransport")
	return &RTCIceTransport{}
}

// RTCIceTransport struct
// js:"RTCIceTransport,omit"
type RTCIceTransport struct {
}

// AddRemoteCandidate fn
// js:"addRemoteCandidate"
func (*RTCIceTransport) AddRemoteCandidate(remoteCandidate interface{}) {
	js.Rewrite("$_.addRemoteCandidate($1)", remoteCandidate)
}

// CreateAssociatedTransport fn
// js:"createAssociatedTransport"
func (*RTCIceTransport) CreateAssociatedTransport() (r *RTCIceTransport) {
	js.Rewrite("$_.createAssociatedTransport()")
	return r
}

// GetNominatedCandidatePair fn
// js:"getNominatedCandidatePair"
func (*RTCIceTransport) GetNominatedCandidatePair() (r *rtcicecandidatepair.RTCIceCandidatePair) {
	js.Rewrite("$_.getNominatedCandidatePair()")
	return r
}

// GetRemoteCandidates fn
// js:"getRemoteCandidates"
func (*RTCIceTransport) GetRemoteCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$_.getRemoteCandidates()")
	return r
}

// GetRemoteParameters fn
// js:"getRemoteParameters"
func (*RTCIceTransport) GetRemoteParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$_.getRemoteParameters()")
	return r
}

// SetRemoteCandidates fn
// js:"setRemoteCandidates"
func (*RTCIceTransport) SetRemoteCandidates(remoteCandidates []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$_.setRemoteCandidates($1)", remoteCandidates)
}

// Start fn
// js:"start"
func (*RTCIceTransport) Start(gatherer *rtcicegatherer.RTCIceGatherer, remoteParameters *rtciceparameters.RTCIceParameters, role *rtcicerole.RTCIceRole) {
	js.Rewrite("$_.start($1, $2, $3)", gatherer, remoteParameters, role)
}

// Stop fn
// js:"stop"
func (*RTCIceTransport) Stop() {
	js.Rewrite("$_.stop()")
}

// GetStats fn
// js:"getStats"
func (*RTCIceTransport) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.getStats()")
	return r
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCIceTransport) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.msGetStats()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCIceTransport) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCIceTransport) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCIceTransport) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Component prop
// js:"component"
func (*RTCIceTransport) Component() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$_.component")
	return component
}

// IceGatherer prop
// js:"iceGatherer"
func (*RTCIceTransport) IceGatherer() (iceGatherer *rtcicegatherer.RTCIceGatherer) {
	js.Rewrite("$_.iceGatherer")
	return iceGatherer
}

// Oncandidatepairchange prop
// js:"oncandidatepairchange"
func (*RTCIceTransport) Oncandidatepairchange() (oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	js.Rewrite("$_.oncandidatepairchange")
	return oncandidatepairchange
}

// SetOncandidatepairchange prop
// js:"oncandidatepairchange"
func (*RTCIceTransport) SetOncandidatepairchange(oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	js.Rewrite("$_.oncandidatepairchange = $1", oncandidatepairchange)
}

// Onicestatechange prop
// js:"onicestatechange"
func (*RTCIceTransport) Onicestatechange() (onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	js.Rewrite("$_.onicestatechange")
	return onicestatechange
}

// SetOnicestatechange prop
// js:"onicestatechange"
func (*RTCIceTransport) SetOnicestatechange(onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	js.Rewrite("$_.onicestatechange = $1", onicestatechange)
}

// Role prop
// js:"role"
func (*RTCIceTransport) Role() (role *rtcicerole.RTCIceRole) {
	js.Rewrite("$_.role")
	return role
}

// State prop
// js:"state"
func (*RTCIceTransport) State() (state *rtcicetransportstate.RTCIceTransportState) {
	js.Rewrite("$_.state")
	return state
}
