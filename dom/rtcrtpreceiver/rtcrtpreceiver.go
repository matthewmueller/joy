package rtcrtpreceiver

import (
	"github.com/matthewmueller/joy/dom/rtcdtlstransport"
	"github.com/matthewmueller/joy/dom/rtcrtpcapabilities"
	"github.com/matthewmueller/joy/dom/rtcrtpcontributingsource"
	"github.com/matthewmueller/joy/dom/rtcrtpparameters"
	"github.com/matthewmueller/joy/dom/rtcstatsprovider"
	"github.com/matthewmueller/joy/dom/rtcstatsreport"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCRtpReceiver)(nil)
var _ window.EventTarget = (*RTCRtpReceiver)(nil)

// New fn
func New(transport interface{}, kind string, rtcptransport *rtcdtlstransport.RTCDtlsTransport) *RTCRtpReceiver {
	macro.Rewrite("new RTCRtpReceiver($1, $2, $3)", transport, kind, rtcptransport)
	return &RTCRtpReceiver{}
}

// RTCRtpReceiver struct
// js:"RTCRtpReceiver,omit"
type RTCRtpReceiver struct {
}

// GetCapabilities fn
// js:"getCapabilities"
func (*RTCRtpReceiver) GetCapabilities(kind *string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	macro.Rewrite("$_.getCapabilities($1)", kind)
	return r
}

// GetContributingSources fn
// js:"getContributingSources"
func (*RTCRtpReceiver) GetContributingSources() (r []*rtcrtpcontributingsource.RTCRtpContributingSource) {
	macro.Rewrite("$_.getContributingSources()")
	return r
}

// Receive fn
// js:"receive"
func (*RTCRtpReceiver) Receive(parameters *rtcrtpparameters.RTCRtpParameters) {
	macro.Rewrite("$_.receive($1)", parameters)
}

// RequestSendCSRC fn
// js:"requestSendCSRC"
func (*RTCRtpReceiver) RequestSendCSRC(csrc uint) {
	macro.Rewrite("$_.requestSendCSRC($1)", csrc)
}

// SetTransport fn
// js:"setTransport"
func (*RTCRtpReceiver) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	macro.Rewrite("$_.setTransport($1, $2)", transport, rtcpTransport)
}

// Stop fn
// js:"stop"
func (*RTCRtpReceiver) Stop() {
	macro.Rewrite("$_.stop()")
}

// GetStats fn
// js:"getStats"
func (*RTCRtpReceiver) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	macro.Rewrite("await $_.getStats()")
	return r
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCRtpReceiver) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	macro.Rewrite("await $_.msGetStats()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCRtpReceiver) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCRtpReceiver) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCRtpReceiver) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onerror prop
// js:"onerror"
func (*RTCRtpReceiver) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCRtpReceiver) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// RtcpTransport prop
// js:"rtcpTransport"
func (*RTCRtpReceiver) RtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	macro.Rewrite("$_.rtcpTransport")
	return rtcpTransport
}

// Track prop
// js:"track"
func (*RTCRtpReceiver) Track() (track *window.MediaStreamTrack) {
	macro.Rewrite("$_.track")
	return track
}

// Transport prop
// js:"transport"
func (*RTCRtpReceiver) Transport() (transport interface{}) {
	macro.Rewrite("$_.transport")
	return transport
}
