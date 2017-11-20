package rtcicegatherer

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidatedictionary"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecomponent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicegathererevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtciceparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsprovider"
	"github.com/matthewmueller/golly/js"
)

type RTCIceGatherer struct {
	*rtcstatsprovider.RTCStatsProvider
}

func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	js.Rewrite("$<.createAssociatedGatherer()")
	return r
}

func (*RTCIceGatherer) GetLocalCandidates() (r []*rtcicecandidatedictionary.RTCIceCandidateDictionary) {
	js.Rewrite("$<.getLocalCandidates()")
	return r
}

func (*RTCIceGatherer) GetLocalParameters() (r *rtciceparameters.RTCIceParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

func (*RTCIceGatherer) GetComponent() (component *rtcicecomponent.RTCIceComponent) {
	js.Rewrite("$<.component")
	return component
}

func (*RTCIceGatherer) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCIceGatherer) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCIceGatherer) GetOnlocalcandidate() (localcandidate *rtcicegathererevent.RTCIceGathererEvent) {
	js.Rewrite("$<.onlocalcandidate")
	return localcandidate
}

func (*RTCIceGatherer) SetOnlocalcandidate(localcandidate *rtcicegathererevent.RTCIceGathererEvent) {
	js.Rewrite("$<.onlocalcandidate = $1", localcandidate)
}
