package rtcicetransport

import (
	"github.com/matthewmueller/golly/dom2/rtcicegatherer"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCIceTransport,omit"
type RTCIceTransport struct {
	rtcstatsprovider.RTCStatsProvider
}

// AddRemoteCandidate
func (*RTCIceTransport) AddRemoteCandidate(remoteCandidate interface{}) {
	js.Rewrite("$<.AddRemoteCandidate($1)", remoteCandidate)
}

// CreateAssociatedTransport
func (*RTCIceTransport) CreateAssociatedTransport() (r *RTCIceTransport) {
	js.Rewrite("$<.CreateAssociatedTransport()")
	return r
}

// GetNominatedCandidatePair
func (*RTCIceTransport) GetNominatedCandidatePair() (r *rtcicecandidatepair.RTCIceCandidatePair) {
	js.Rewrite("$<.GetNominatedCandidatePair()")
	return r
}

// GetRemoteCandidates
func (*RTCIceTransport) GetRemoteCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.GetRemoteCandidates()")
	return r
}

// GetRemoteParameters
func (*RTCIceTransport) GetRemoteParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$<.GetRemoteParameters()")
	return r
}

// SetRemoteCandidates
func (*RTCIceTransport) SetRemoteCandidates(remoteCandidates []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.SetRemoteCandidates($1)", remoteCandidates)
}

// Start
func (*RTCIceTransport) Start(gatherer *rtcicegatherer.RTCIceGatherer, remoteParameters *rtciceparameters.RTCIceParameters, role *rtcicerole.RTCIceRole) {
	js.Rewrite("$<.Start($1, $2, $3)", gatherer, remoteParameters, role)
}

// Stop
func (*RTCIceTransport) Stop() {
	js.Rewrite("$<.Stop()")
}

// Component
func (*RTCIceTransport) Component() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$<.Component")
	return component
}

// IceGatherer
func (*RTCIceTransport) IceGatherer() (iceGatherer *rtcicegatherer.RTCIceGatherer) {
	js.Rewrite("$<.IceGatherer")
	return iceGatherer
}

// Oncandidatepairchange
func (*RTCIceTransport) Oncandidatepairchange() (oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	js.Rewrite("$<.Oncandidatepairchange")
	return oncandidatepairchange
}

// Oncandidatepairchange
func (*RTCIceTransport) SetOncandidatepairchange(oncandidatepairchange func(*rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent)) {
	js.Rewrite("$<.Oncandidatepairchange = $1", oncandidatepairchange)
}

// Onicestatechange
func (*RTCIceTransport) Onicestatechange() (onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	js.Rewrite("$<.Onicestatechange")
	return onicestatechange
}

// Onicestatechange
func (*RTCIceTransport) SetOnicestatechange(onicestatechange func(*rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent)) {
	js.Rewrite("$<.Onicestatechange = $1", onicestatechange)
}

// Role
func (*RTCIceTransport) Role() (role *rtcicerole.RTCIceRole) {
	js.Rewrite("$<.Role")
	return role
}

// State
func (*RTCIceTransport) State() (state *rtcicetransportstate.RTCIceTransportState) {
	js.Rewrite("$<.State")
	return state
}
