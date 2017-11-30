package rtcrtpsender

import (
	"github.com/matthewmueller/golly/dom/rtcdtlstransport"
	"github.com/matthewmueller/golly/dom/rtcrtpcapabilities"
	"github.com/matthewmueller/golly/dom/rtcrtpparameters"
	"github.com/matthewmueller/golly/dom/rtcssrcconflictevent"
	"github.com/matthewmueller/golly/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ rtcstatsprovider.RTCStatsProvider = (*RTCRtpSender)(nil)
var _ window.EventTarget = (*RTCRtpSender)(nil)

// RTCRtpSender struct
// js:"RTCRtpSender,omit"
type RTCRtpSender struct {
}

// GetCapabilities fn
// js:"getCapabilities"
func (*RTCRtpSender) GetCapabilities(kind *string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	js.Rewrite("$_.getCapabilities($1)", kind)
	return r
}

// Send fn
// js:"send"
func (*RTCRtpSender) Send(parameters *rtcrtpparameters.RTCRtpParameters) {
	js.Rewrite("$_.send($1)", parameters)
}

// SetTrack fn
// js:"setTrack"
func (*RTCRtpSender) SetTrack(track *window.MediaStreamTrack) {
	js.Rewrite("$_.setTrack($1)", track)
}

// SetTransport fn
// js:"setTransport"
func (*RTCRtpSender) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$_.setTransport($1, $2)", transport, rtcpTransport)
}

// Stop fn
// js:"stop"
func (*RTCRtpSender) Stop() {
	js.Rewrite("$_.stop()")
}

// GetStats fn
// js:"getStats"
func (*RTCRtpSender) GetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.getStats()")
	return r
}

// MsGetStats fn
// js:"msGetStats"
func (*RTCRtpSender) MsGetStats() (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $_.msGetStats()")
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*RTCRtpSender) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*RTCRtpSender) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*RTCRtpSender) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onerror prop
// js:"onerror"
func (*RTCRtpSender) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*RTCRtpSender) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onssrcconflict prop
// js:"onssrcconflict"
func (*RTCRtpSender) Onssrcconflict() (onssrcconflict func(*rtcssrcconflictevent.RTCSsrcConflictEvent)) {
	js.Rewrite("$_.onssrcconflict")
	return onssrcconflict
}

// SetOnssrcconflict prop
// js:"onssrcconflict"
func (*RTCRtpSender) SetOnssrcconflict(onssrcconflict func(*rtcssrcconflictevent.RTCSsrcConflictEvent)) {
	js.Rewrite("$_.onssrcconflict = $1", onssrcconflict)
}

// RtcpTransport prop
// js:"rtcpTransport"
func (*RTCRtpSender) RtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$_.rtcpTransport")
	return rtcpTransport
}

// Track prop
// js:"track"
func (*RTCRtpSender) Track() (track *window.MediaStreamTrack) {
	js.Rewrite("$_.track")
	return track
}

// Transport prop
// js:"transport"
func (*RTCRtpSender) Transport() (transport interface{}) {
	js.Rewrite("$_.transport")
	return transport
}
