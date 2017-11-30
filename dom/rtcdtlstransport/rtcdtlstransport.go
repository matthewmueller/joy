package rtcdtlstransport

import (
	"github.com/matthewmueller/golly/dom/rtcdtlsparameters"
	"github.com/matthewmueller/golly/dom/rtcdtlstransportstate"
	"github.com/matthewmueller/golly/dom/rtcdtlstransportstatechangedevent"
	"github.com/matthewmueller/golly/dom/rtcicetransport"
	"github.com/matthewmueller/golly/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCDtlsTransport)(nil)
var _ window.EventTarget = (*RTCDtlsTransport)(nil)

// RTCDtlsTransport struct
// js:"RTCDtlsTransport,omit"
type RTCDtlsTransport struct {
}

// GetLocalParameters fn
// js:"getLocalParameters"
func (*RTCDtlsTransport) GetLocalParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$_.getLocalParameters()")
	return r
}

// GetRemoteCertificates fn
// js:"getRemoteCertificates"
func (*RTCDtlsTransport) GetRemoteCertificates() (b [][]byte) {
	js.Rewrite("$_.getRemoteCertificates()")
	return b
}

// GetRemoteParameters fn
// js:"getRemoteParameters"
func (*RTCDtlsTransport) GetRemoteParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$_.getRemoteParameters()")
	return r
}

// Start fn
// js:"start"
func (*RTCDtlsTransport) Start(remoteParameters *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$_.start($1)", remoteParameters)
}

// Stop fn
// js:"stop"
func (*RTCDtlsTransport) Stop() {
	js.Rewrite("$_.stop()")
}

// GetStats fn
// js:"getStats"
func (*RTCDtlsTransport) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.getStats()")
	return r
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCDtlsTransport) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.msGetStats()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCDtlsTransport) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCDtlsTransport) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCDtlsTransport) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Ondtlsstatechange prop
// js:"ondtlsstatechange"
func (*RTCDtlsTransport) Ondtlsstatechange() (ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	js.Rewrite("$_.ondtlsstatechange")
	return ondtlsstatechange
}

// SetOndtlsstatechange prop
// js:"ondtlsstatechange"
func (*RTCDtlsTransport) SetOndtlsstatechange(ondtlsstatechange func(*rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent)) {
	js.Rewrite("$_.ondtlsstatechange = $1", ondtlsstatechange)
}

// Onerror prop
// js:"onerror"
func (*RTCDtlsTransport) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCDtlsTransport) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// State prop
// js:"state"
func (*RTCDtlsTransport) State() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$_.state")
	return state
}

// Transport prop
// js:"transport"
func (*RTCDtlsTransport) Transport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$_.transport")
	return transport
}
