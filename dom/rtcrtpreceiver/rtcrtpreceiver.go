package rtcrtpreceiver

import (
	"github.com/matthewmueller/golly/dom/rtcdtlstransport"
	"github.com/matthewmueller/golly/dom/rtcrtpcapabilities"
	"github.com/matthewmueller/golly/dom/rtcrtpcontributingsource"
	"github.com/matthewmueller/golly/dom/rtcrtpparameters"
	"github.com/matthewmueller/golly/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(transport interface{}, kind string, rtcptransport *rtcdtlstransport.RTCDtlsTransport) *RTCRtpReceiver {
	js.Rewrite("RTCRtpReceiver")
	return &RTCRtpReceiver{}
}

// RTCRtpReceiver struct
// js:"RTCRtpReceiver,omit"
type RTCRtpReceiver struct {
	rtcstatsprovider.RTCStatsProvider
}

// GetCapabilities fn
func (*RTCRtpReceiver) GetCapabilities(kind *string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	js.Rewrite("$<.getCapabilities($1)", kind)
	return r
}

// GetContributingSources fn
func (*RTCRtpReceiver) GetContributingSources() (r []*rtcrtpcontributingsource.RTCRtpContributingSource) {
	js.Rewrite("$<.getContributingSources()")
	return r
}

// Receive fn
func (*RTCRtpReceiver) Receive(parameters *rtcrtpparameters.RTCRtpParameters) {
	js.Rewrite("$<.receive($1)", parameters)
}

// RequestSendCSRC fn
func (*RTCRtpReceiver) RequestSendCSRC(csrc uint) {
	js.Rewrite("$<.requestSendCSRC($1)", csrc)
}

// SetTransport fn
func (*RTCRtpReceiver) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.setTransport($1, $2)", transport, rtcpTransport)
}

// Stop fn
func (*RTCRtpReceiver) Stop() {
	js.Rewrite("$<.stop()")
}

// Onerror prop
func (*RTCRtpReceiver) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*RTCRtpReceiver) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// RtcpTransport prop
func (*RTCRtpReceiver) RtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.rtcpTransport")
	return rtcpTransport
}

// Track prop
func (*RTCRtpReceiver) Track() (track *window.MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}

// Transport prop
func (*RTCRtpReceiver) Transport() (transport interface{}) {
	js.Rewrite("$<.transport")
	return transport
}
