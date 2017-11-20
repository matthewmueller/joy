package rtcicetransport

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidatedictionary"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidatepair"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidatepairchangedevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecomponent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicegatherer"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtciceparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicerole"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicetransportstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicetransportstatechangedevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/js"
)

type RTCIceTransport struct {
	*rtcstatsprovider.RTCStatsProvider
}

func (*RTCIceTransport) AddRemoteCandidate(remoteCandidate interface{}) {
	js.Rewrite("$<.addRemoteCandidate($1)", remoteCandidate)
}

func (*RTCIceTransport) CreateAssociatedTransport() (r *RTCIceTransport) {
	js.Rewrite("$<.createAssociatedTransport()")
	return r
}

func (*RTCIceTransport) GetNominatedCandidatePair() (r *rtcicecandidatepair.RTCIceCandidatePair) {
	js.Rewrite("$<.getNominatedCandidatePair()")
	return r
}

func (*RTCIceTransport) GetRemoteCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.getRemoteCandidates()")
	return r
}

func (*RTCIceTransport) GetRemoteParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$<.getRemoteParameters()")
	return r
}

func (*RTCIceTransport) SetRemoteCandidates(remoteCandidates []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.setRemoteCandidates($1)", remoteCandidates)
}

func (*RTCIceTransport) Start(gatherer *rtcicegatherer.RTCIceGatherer, remoteParameters *rtciceparameters.RTCIceParameters, role *rtcicerole.RTCIceRole) {
	js.Rewrite("$<.start($1, $2, $3)", gatherer, remoteParameters, role)
}

func (*RTCIceTransport) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCIceTransport) GetComponent() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$<.component")
	return component
}

func (*RTCIceTransport) GetIceGatherer() (iceGatherer *rtcicegatherer.RTCIceGatherer) {
	js.Rewrite("$<.iceGatherer")
	return iceGatherer
}

func (*RTCIceTransport) GetOncandidatepairchange() (candidatepairchange *rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent) {
	js.Rewrite("$<.oncandidatepairchange")
	return candidatepairchange
}

func (*RTCIceTransport) SetOncandidatepairchange(candidatepairchange *rtcicecandidatepairchangedevent.RTCIceCandidatePairChangedEvent) {
	js.Rewrite("$<.oncandidatepairchange = $1", candidatepairchange)
}

func (*RTCIceTransport) GetOnicestatechange() (icestatechange *rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent) {
	js.Rewrite("$<.onicestatechange")
	return icestatechange
}

func (*RTCIceTransport) SetOnicestatechange(icestatechange *rtcicetransportstatechangedevent.RTCIceTransportStateChangedEvent) {
	js.Rewrite("$<.onicestatechange = $1", icestatechange)
}

func (*RTCIceTransport) GetRole() (role *rtcicerole.RTCIceRole) {
	js.Rewrite("$<.role")
	return role
}

func (*RTCIceTransport) GetState() (state *rtcicetransportstate.RTCIceTransportState) {
	js.Rewrite("$<.state")
	return state
}
