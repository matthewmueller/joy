package rtcrtpsender

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlstransport"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpcapabilities"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcssrcconflictevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/js"
)

type RTCRtpSender struct {
	*rtcstatsprovider.RTCStatsProvider
}

func (*RTCRtpSender) GetCapabilities(kind string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	js.Rewrite("$<.getCapabilities($1)", kind)
	return r
}

func (*RTCRtpSender) Send(parameters *rtcrtpparameters.RTCRtpParameters) {
	js.Rewrite("$<.send($1)", parameters)
}

func (*RTCRtpSender) SetTrack(track *mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.setTrack($1)", track)
}

func (*RTCRtpSender) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.setTransport($1, $2)", transport, rtcpTransport)
}

func (*RTCRtpSender) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCRtpSender) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCRtpSender) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCRtpSender) GetOnssrcconflict() (ssrcconflict *rtcssrcconflictevent.RTCSsrcConflictEvent) {
	js.Rewrite("$<.onssrcconflict")
	return ssrcconflict
}

func (*RTCRtpSender) SetOnssrcconflict(ssrcconflict *rtcssrcconflictevent.RTCSsrcConflictEvent) {
	js.Rewrite("$<.onssrcconflict = $1", ssrcconflict)
}

func (*RTCRtpSender) GetRtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.rtcpTransport")
	return rtcpTransport
}

func (*RTCRtpSender) GetTrack() (track *mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}

func (*RTCRtpSender) GetTransport() (transport interface{}) {
	js.Rewrite("$<.transport")
	return transport
}
