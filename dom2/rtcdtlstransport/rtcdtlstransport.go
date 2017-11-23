package rtcdtlstransport

import (
	"github.com/matthewmueller/golly/dom2/rtcdtlsparameters"
	"github.com/matthewmueller/golly/dom2/rtcdtlstransportstate"
	"github.com/matthewmueller/golly/dom2/rtcdtlstransportstatechangedevent"
	"github.com/matthewmueller/golly/dom2/rtcicetransport"
	"github.com/matthewmueller/golly/dom2/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCDtlsTransport,omit"
type RTCDtlsTransport struct {
	rtcstatsprovider.RTCStatsProvider
}

// GetLocalParameters
func (*RTCDtlsTransport) GetLocalParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.GetLocalParameters()")
	return r
}

// GetRemoteCertificates
func (*RTCDtlsTransport) GetRemoteCertificates() (b [][]byte) {
	js.Rewrite("$<.GetRemoteCertificates()")
	return b
}

// GetRemoteParameters
func (*RTCDtlsTransport) GetRemoteParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.GetRemoteParameters()")
	return r
}

// Start
func (*RTCDtlsTransport) Start(remoteParameters *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.Start($1)", remoteParameters)
}

// Stop
func (*RTCDtlsTransport) Stop() {
	js.Rewrite("$<.Stop()")
}

// Ondtlsstatechange
func (*RTCDtlsTransport) Ondtlsstatechange() (ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	js.Rewrite("$<.Ondtlsstatechange")
	return ondtlsstatechange
}

// Ondtlsstatechange
func (*RTCDtlsTransport) SetOndtlsstatechange(ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	js.Rewrite("$<.Ondtlsstatechange = $1", ondtlsstatechange)
}

// Onerror
func (*RTCDtlsTransport) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*RTCDtlsTransport) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// State
func (*RTCDtlsTransport) State() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$<.State")
	return state
}

// Transport
func (*RTCDtlsTransport) Transport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$<.Transport")
	return transport
}
