package rtcdtlstransport

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlsparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlstransportstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlstransportstatechangedevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicetransport"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/js"
)

type RTCDtlsTransport struct {
	*rtcstatsprovider.RTCStatsProvider
}

func (*RTCDtlsTransport) GetLocalParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

func (*RTCDtlsTransport) GetRemoteCertificates() (b [][]byte) {
	js.Rewrite("$<.getRemoteCertificates()")
	return b
}

func (*RTCDtlsTransport) GetRemoteParameters() (r *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.getRemoteParameters()")
	return r
}

func (*RTCDtlsTransport) Start(remoteParameters *rtcdtlsparameters.RTCDtlsParameters) {
	js.Rewrite("$<.start($1)", remoteParameters)
}

func (*RTCDtlsTransport) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCDtlsTransport) GetOndtlsstatechange() (dtlsstatechange *rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent) {
	js.Rewrite("$<.ondtlsstatechange")
	return dtlsstatechange
}

func (*RTCDtlsTransport) SetOndtlsstatechange(dtlsstatechange *rtcdtlstransportstatechangedevent.RTCDtlsTransportStateChangedEvent) {
	js.Rewrite("$<.ondtlsstatechange = $1", dtlsstatechange)
}

func (*RTCDtlsTransport) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCDtlsTransport) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCDtlsTransport) GetState() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$<.state")
	return state
}

func (*RTCDtlsTransport) GetTransport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$<.transport")
	return transport
}
