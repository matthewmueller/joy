package rtcrtpsender

import (
	"github.com/matthewmueller/golly/dom2/rtcdtlstransport"
	"github.com/matthewmueller/golly/dom2/rtcrtpcapabilities"
	"github.com/matthewmueller/golly/dom2/rtcrtpparameters"
	"github.com/matthewmueller/golly/dom2/rtcssrcconflictevent"
	"github.com/matthewmueller/golly/dom2/rtcstatsprovider"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(track *window.MediaStreamTrack, transport interface{}, rtcptransport *rtcdtlstransport.RTCDtlsTransport) *RTCRtpSender {
	js.Rewrite("RTCRtpSender")
	return &RTCRtpSender{}
}

// RTCRtpSender struct
// js:"RTCRtpSender,omit"
type RTCRtpSender struct {
	rtcstatsprovider.RTCStatsProvider
}

// GetCapabilities fn
func (*RTCRtpSender) GetCapabilities(kind *string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	js.Rewrite("$<.getCapabilities($1)", kind)
	return r
}

// Send fn
func (*RTCRtpSender) Send(parameters *rtcrtpparameters.RTCRtpParameters) {
	js.Rewrite("$<.send($1)", parameters)
}

// SetTrack fn
func (*RTCRtpSender) SetTrack(track *window.MediaStreamTrack) {
	js.Rewrite("$<.setTrack($1)", track)
}

// SetTransport fn
func (*RTCRtpSender) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.setTransport($1, $2)", transport, rtcpTransport)
}

// Stop fn
func (*RTCRtpSender) Stop() {
	js.Rewrite("$<.stop()")
}

// Onerror prop
func (*RTCRtpSender) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*RTCRtpSender) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onssrcconflict prop
func (*RTCRtpSender) Onssrcconflict() (onssrcconflict func(*rtcssrcconflictevent.RTCSsrcConflictEvent)) {
	js.Rewrite("$<.onssrcconflict")
	return onssrcconflict
}

// Onssrcconflict prop
func (*RTCRtpSender) SetOnssrcconflict(onssrcconflict func(*rtcssrcconflictevent.RTCSsrcConflictEvent)) {
	js.Rewrite("$<.onssrcconflict = $1", onssrcconflict)
}

// RtcpTransport prop
func (*RTCRtpSender) RtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.rtcpTransport")
	return rtcpTransport
}

// Track prop
func (*RTCRtpSender) Track() (track *window.MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}

// Transport prop
func (*RTCRtpSender) Transport() (transport interface{}) {
	js.Rewrite("$<.transport")
	return transport
}
