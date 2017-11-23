package rtcrtpsender

import (
	"github.com/matthewmueller/golly/dom2/rtcrtpparameters"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCRtpSender,omit"
type RTCRtpSender struct {
	rtcstatsprovider.RTCStatsProvider
}

// GetCapabilities
func (*RTCRtpSender) GetCapabilities(kind *string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	js.Rewrite("$<.GetCapabilities($1)", kind)
	return r
}

// Send
func (*RTCRtpSender) Send(parameters *rtcrtpparameters.RTCRtpParameters) {
	js.Rewrite("$<.Send($1)", parameters)
}

// SetTrack
func (*RTCRtpSender) SetTrack(track *window.MediaStreamTrack) {
	js.Rewrite("$<.SetTrack($1)", track)
}

// SetTransport
func (*RTCRtpSender) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.SetTransport($1, $2)", transport, rtcpTransport)
}

// Stop
func (*RTCRtpSender) Stop() {
	js.Rewrite("$<.Stop()")
}

// Onerror
func (*RTCRtpSender) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*RTCRtpSender) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onssrcconflict
func (*RTCRtpSender) Onssrcconflict() (onssrcconflict func(*rtcssrcconflictevent.RTCSsrcConflictEvent)) {
	js.Rewrite("$<.Onssrcconflict")
	return onssrcconflict
}

// Onssrcconflict
func (*RTCRtpSender) SetOnssrcconflict(onssrcconflict func(*rtcssrcconflictevent.RTCSsrcConflictEvent)) {
	js.Rewrite("$<.Onssrcconflict = $1", onssrcconflict)
}

// RtcpTransport
func (*RTCRtpSender) RtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.RtcpTransport")
	return rtcpTransport
}

// Track
func (*RTCRtpSender) Track() (track *window.MediaStreamTrack) {
	js.Rewrite("$<.Track")
	return track
}

// Transport
func (*RTCRtpSender) Transport() (transport interface{}) {
	js.Rewrite("$<.Transport")
	return transport
}
