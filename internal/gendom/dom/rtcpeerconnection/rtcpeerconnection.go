package rtcpeerconnection

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastream"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcconfiguration"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidate"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtciceconnectionstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicegatheringstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcofferoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcpeerconnectioniceevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcsessiondescription"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcsignalingstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/js"
)

type RTCPeerConnection struct {
	*eventtarget.EventTarget
}

func (*RTCPeerConnection) AddIceCandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback func(), failureCallback func(err *domerror.DOMError)) {
	js.Rewrite("await $<.addIceCandidate($1, $2, $3)", candidate, successCallback, failureCallback)
}

func (*RTCPeerConnection) AddStream(stream *mediastream.MediaStream) {
	js.Rewrite("$<.addStream($1)", stream)
}

func (*RTCPeerConnection) Close() {
	js.Rewrite("$<.close()")
}

func (*RTCPeerConnection) CreateAnswer(successCallback func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("await $<.createAnswer($1, $2)", successCallback, failureCallback)
	return r
}

func (*RTCPeerConnection) CreateOffer(successCallback func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("await $<.createOffer($1, $2, $3)", successCallback, failureCallback, options)
	return r
}

func (*RTCPeerConnection) GetConfiguration() (r *rtcconfiguration.RTCConfiguration) {
	js.Rewrite("$<.getConfiguration()")
	return r
}

func (*RTCPeerConnection) GetLocalStreams() (m []*mediastream.MediaStream) {
	js.Rewrite("$<.getLocalStreams()")
	return m
}

func (*RTCPeerConnection) GetRemoteStreams() (m []*mediastream.MediaStream) {
	js.Rewrite("$<.getRemoteStreams()")
	return m
}

func (*RTCPeerConnection) GetStats(selector *mediastreamtrack.MediaStreamTrack, successCallback func(report *rtcstatsreport.RTCStatsReport), failureCallback func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $<.getStats($1, $2, $3)", selector, successCallback, failureCallback)
	return r
}

func (*RTCPeerConnection) GetStreamByID(streamId string) (m *mediastream.MediaStream) {
	js.Rewrite("$<.getStreamById($1)", streamId)
	return m
}

func (*RTCPeerConnection) RemoveStream(stream *mediastream.MediaStream) {
	js.Rewrite("$<.removeStream($1)", stream)
}

func (*RTCPeerConnection) SetLocalDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback func(), failureCallback func(err *domerror.DOMError)) {
	js.Rewrite("await $<.setLocalDescription($1, $2, $3)", description, successCallback, failureCallback)
}

func (*RTCPeerConnection) SetRemoteDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback func(), failureCallback func(err *domerror.DOMError)) {
	js.Rewrite("await $<.setRemoteDescription($1, $2, $3)", description, successCallback, failureCallback)
}

func (*RTCPeerConnection) GetCanTrickleIceCandidates() (canTrickleIceCandidates bool) {
	js.Rewrite("$<.canTrickleIceCandidates")
	return canTrickleIceCandidates
}

func (*RTCPeerConnection) GetIceConnectionState() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState) {
	js.Rewrite("$<.iceConnectionState")
	return iceConnectionState
}

func (*RTCPeerConnection) GetIceGatheringState() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState) {
	js.Rewrite("$<.iceGatheringState")
	return iceGatheringState
}

func (*RTCPeerConnection) GetLocalDescription() (localDescription *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("$<.localDescription")
	return localDescription
}

func (*RTCPeerConnection) GetOnaddstream() (addstream *mediastreamevent.MediaStreamEvent) {
	js.Rewrite("$<.onaddstream")
	return addstream
}

func (*RTCPeerConnection) SetOnaddstream(addstream *mediastreamevent.MediaStreamEvent) {
	js.Rewrite("$<.onaddstream = $1", addstream)
}

func (*RTCPeerConnection) GetOnicecandidate() (icecandidate *rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent) {
	js.Rewrite("$<.onicecandidate")
	return icecandidate
}

func (*RTCPeerConnection) SetOnicecandidate(icecandidate *rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent) {
	js.Rewrite("$<.onicecandidate = $1", icecandidate)
}

func (*RTCPeerConnection) GetOniceconnectionstatechange() (iceconnectionstatechange *event.Event) {
	js.Rewrite("$<.oniceconnectionstatechange")
	return iceconnectionstatechange
}

func (*RTCPeerConnection) SetOniceconnectionstatechange(iceconnectionstatechange *event.Event) {
	js.Rewrite("$<.oniceconnectionstatechange = $1", iceconnectionstatechange)
}

func (*RTCPeerConnection) GetOnicegatheringstatechange() (icegatheringstatechange *event.Event) {
	js.Rewrite("$<.onicegatheringstatechange")
	return icegatheringstatechange
}

func (*RTCPeerConnection) SetOnicegatheringstatechange(icegatheringstatechange *event.Event) {
	js.Rewrite("$<.onicegatheringstatechange = $1", icegatheringstatechange)
}

func (*RTCPeerConnection) GetOnnegotiationneeded() (negotiationneeded *event.Event) {
	js.Rewrite("$<.onnegotiationneeded")
	return negotiationneeded
}

func (*RTCPeerConnection) SetOnnegotiationneeded(negotiationneeded *event.Event) {
	js.Rewrite("$<.onnegotiationneeded = $1", negotiationneeded)
}

func (*RTCPeerConnection) GetOnremovestream() (removestream *mediastreamevent.MediaStreamEvent) {
	js.Rewrite("$<.onremovestream")
	return removestream
}

func (*RTCPeerConnection) SetOnremovestream(removestream *mediastreamevent.MediaStreamEvent) {
	js.Rewrite("$<.onremovestream = $1", removestream)
}

func (*RTCPeerConnection) GetOnsignalingstatechange() (signalingstatechange *event.Event) {
	js.Rewrite("$<.onsignalingstatechange")
	return signalingstatechange
}

func (*RTCPeerConnection) SetOnsignalingstatechange(signalingstatechange *event.Event) {
	js.Rewrite("$<.onsignalingstatechange = $1", signalingstatechange)
}

func (*RTCPeerConnection) GetRemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("$<.remoteDescription")
	return remoteDescription
}

func (*RTCPeerConnection) GetSignalingState() (signalingState *rtcsignalingstate.RTCSignalingState) {
	js.Rewrite("$<.signalingState")
	return signalingState
}
