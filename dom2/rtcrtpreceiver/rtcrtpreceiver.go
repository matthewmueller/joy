package rtcrtpreceiver

import (
	"github.com/matthewmueller/golly/dom2/rtcdtlstransport"
	"github.com/matthewmueller/golly/dom2/rtcrtpcapabilities"
	"github.com/matthewmueller/golly/dom2/rtcrtpcontributingsource"
	"github.com/matthewmueller/golly/dom2/rtcrtpparameters"
	"github.com/matthewmueller/golly/dom2/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCRtpReceiver,omit"
type RTCRtpReceiver struct {
	rtcstatsprovider.RTCStatsProvider
}

// GetCapabilities
func (*RTCRtpReceiver) GetCapabilities(kind *string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	js.Rewrite("$<.GetCapabilities($1)", kind)
	return r
}

// GetContributingSources
func (*RTCRtpReceiver) GetContributingSources() (r []*rtcrtpcontributingsource.RTCRtpContributingSource) {
	js.Rewrite("$<.GetContributingSources()")
	return r
}

// Receive
func (*RTCRtpReceiver) Receive(parameters *rtcrtpparameters.RTCRtpParameters) {
	js.Rewrite("$<.Receive($1)", parameters)
}

// RequestSendCSRC
func (*RTCRtpReceiver) RequestSendCSRC(csrc uint) {
	js.Rewrite("$<.RequestSendCSRC($1)", csrc)
}

// SetTransport
func (*RTCRtpReceiver) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.SetTransport($1, $2)", transport, rtcpTransport)
}

// Stop
func (*RTCRtpReceiver) Stop() {
	js.Rewrite("$<.Stop()")
}

// Onerror
func (*RTCRtpReceiver) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*RTCRtpReceiver) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// RtcpTransport
func (*RTCRtpReceiver) RtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.RtcpTransport")
	return rtcpTransport
}

// Track
func (*RTCRtpReceiver) Track() (track *window.MediaStreamTrack) {
	js.Rewrite("$<.Track")
	return track
}

// Transport
func (*RTCRtpReceiver) Transport() (transport interface{}) {
	js.Rewrite("$<.Transport")
	return transport
}
