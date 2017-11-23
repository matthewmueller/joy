package rtcdtlstransport

import (
	"github.com/matthewmueller/golly/dom/rtcdtlsparameters"
	"github.com/matthewmueller/golly/dom/rtcdtlstransportstate"
	"github.com/matthewmueller/golly/dom/rtcdtlstransportstatechangedevent"
	"github.com/matthewmueller/golly/dom/rtcicetransport"
	"github.com/matthewmueller/golly/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(transport *rtcicetransport.RTCIceTransport) *RTCDtlsTransport {
	js.Rewrite("RTCDtlsTransport")
	return &RTCDtlsTransport{}
}

// RTCDtlsTransport struct
// js:"RTCDtlsTransport,omit"
type RTCDtlsTransport struct {
	rtcstatsprovider.RTCStatsProvider
}

// GetLocalParameters fn
func (*RTCDtlsTransport) GetLocalParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

// GetRemoteCertificates fn
func (*RTCDtlsTransport) GetRemoteCertificates() (b [][]byte) {
	js.Rewrite("$<.getRemoteCertificates()")
	return b
}

// GetRemoteParameters fn
func (*RTCDtlsTransport) GetRemoteParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.getRemoteParameters()")
	return r
}

// Start fn
func (*RTCDtlsTransport) Start(remoteParameters *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.start($1)", remoteParameters)
}

// Stop fn
func (*RTCDtlsTransport) Stop() {
	js.Rewrite("$<.stop()")
}

// Ondtlsstatechange prop
func (*RTCDtlsTransport) Ondtlsstatechange() (ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	js.Rewrite("$<.ondtlsstatechange")
	return ondtlsstatechange
}

// Ondtlsstatechange prop
func (*RTCDtlsTransport) SetOndtlsstatechange(ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	js.Rewrite("$<.ondtlsstatechange = $1", ondtlsstatechange)
}

// Onerror prop
func (*RTCDtlsTransport) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*RTCDtlsTransport) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// State prop
func (*RTCDtlsTransport) State() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$<.state")
	return state
}

// Transport prop
func (*RTCDtlsTransport) Transport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$<.transport")
	return transport
}
