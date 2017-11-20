package rtcrtpreceiver

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlstransport"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpcapabilities"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpcontributingsource"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/js"
)

type RTCRtpReceiver struct {
	*rtcstatsprovider.RTCStatsProvider
}

func (*RTCRtpReceiver) GetCapabilities(kind string) (r *rtcrtpcapabilities.RTCRtpCapabilities) {
	js.Rewrite("$<.getCapabilities($1)", kind)
	return r
}

func (*RTCRtpReceiver) GetContributingSources() (r []*rtcrtpcontributingsource.RTCRtpContributingSource) {
	js.Rewrite("$<.getContributingSources()")
	return r
}

func (*RTCRtpReceiver) Receive(parameters *rtcrtpparameters.RTCRtpParameters) {
	js.Rewrite("$<.receive($1)", parameters)
}

func (*RTCRtpReceiver) RequestSendCSRC(csrc uint) {
	js.Rewrite("$<.requestSendCSRC($1)", csrc)
}

func (*RTCRtpReceiver) SetTransport(transport interface{}, rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.setTransport($1, $2)", transport, rtcpTransport)
}

func (*RTCRtpReceiver) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCRtpReceiver) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCRtpReceiver) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCRtpReceiver) GetRtcpTransport() (rtcpTransport *rtcdtlstransport.RTCDtlsTransport) {
	js.Rewrite("$<.rtcpTransport")
	return rtcpTransport
}

func (*RTCRtpReceiver) GetTrack() (track *mediastreamtrack.MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}

func (*RTCRtpReceiver) GetTransport() (transport interface{}) {
	js.Rewrite("$<.transport")
	return transport
}
