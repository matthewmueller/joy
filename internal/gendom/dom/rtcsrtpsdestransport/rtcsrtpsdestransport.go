package rtcsrtpsdestransport

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicetransport"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcsrtpsdesparameters"
	"github.com/matthewmueller/golly/js"
)

type RTCSrtpSdesTransport struct {
	*eventtarget.EventTarget
}

func (*RTCSrtpSdesTransport) GetLocalParameters() (r []*rtcsrtpsdesparameters.RTCSrtpSdesParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

func (*RTCSrtpSdesTransport) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCSrtpSdesTransport) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCSrtpSdesTransport) GetTransport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$<.transport")
	return transport
}
