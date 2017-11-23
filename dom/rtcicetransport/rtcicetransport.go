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
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *RTCIceTransport {
	js.Rewrite("RTCIceTransport")
	return &RTCIceTransport{}
}

// RTCIceTransport struct
// js:"RTCIceTransport,omit"
type RTCIceTransport struct {
	rtcstatsprovider.RTCStatsProvider
}

// AddRemoteCandidate fn
func (*RTCIceTransport) AddRemoteCandidate(remoteCandidate interface{}) {
	js.Rewrite("$<.addRemoteCandidate($1)", remoteCandidate)
}

// CreateAssociatedTransport fn
func (*RTCIceTransport) CreateAssociatedTransport() (r *RTCIceTransport) {
	js.Rewrite("$<.createAssociatedTransport()")
	return r
}

// GetNominatedCandidatePair fn
func (*RTCIceTransport) GetNominatedCandidatePair() (r *rtcicecandidatepair.RTCIceCandidatePair) {
	js.Rewrite("$<.getNominatedCandidatePair()")
	return r
}

// GetRemoteCandidates fn
func (*RTCIceTransport) GetRemoteCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.getRemoteCandidates()")
	return r
}

// GetRemoteParameters fn
func (*RTCIceTransport) GetRemoteParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$<.getRemoteParameters()")
	return r
}

// SetRemoteCandidates fn
func (*RTCIceTransport) SetRemoteCandidates(remoteCandidates []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.setRemoteCandidates($1)", remoteCandidates)
}

// Start fn
func (*RTCIceTransport) Start(gatherer *rtcicegatherer.RTCIceGatherer, remoteParameters *rtciceparameters.RTCIceParameters, role *rtcicerole.RTCIceRole) {
	js.Rewrite("$<.start($1, $2, $3)", gatherer, remoteParameters, role)
}

// Stop fn
func (*RTCIceTransport) Stop() {
	js.Rewrite("$<.stop()")
}

// Component prop
func (*RTCIceTransport) Component() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$<.component")
	return component
}

// IceGatherer prop
func (*RTCIceTransport) IceGatherer() (iceGatherer *rtcicegatherer.RTCIceGatherer) {
	js.Rewrite("$<.iceGatherer")
	return iceGatherer
}

// Oncandidatepairchange prop
func (*RTCIceTransport) Oncandidatepairchange() (oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	js.Rewrite("$<.oncandidatepairchange")
	return oncandidatepairchange
}

// Oncandidatepairchange prop
func (*RTCIceTransport) SetOncandidatepairchange(oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	js.Rewrite("$<.oncandidatepairchange = $1", oncandidatepairchange)
}

// Onicestatechange prop
func (*RTCIceTransport) Onicestatechange() (onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	js.Rewrite("$<.onicestatechange")
	return onicestatechange
}

// Onicestatechange prop
func (*RTCIceTransport) SetOnicestatechange(onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	js.Rewrite("$<.onicestatechange = $1", onicestatechange)
}

// Role prop
func (*RTCIceTransport) Role() (role *rtcicerole.RTCIceRole) {
	js.Rewrite("$<.role")
	return role
}

// State prop
func (*RTCIceTransport) State() (state *rtcicetransportstate.RTCIceTransportState) {
	js.Rewrite("$<.state")
	return state
}
