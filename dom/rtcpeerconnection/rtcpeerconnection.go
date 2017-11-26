package rtcpeerconnection

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/mediastreamevent"
	"github.com/matthewmueller/golly/dom/rtcconfiguration"
	"github.com/matthewmueller/golly/dom/rtcicecandidate"
	"github.com/matthewmueller/golly/dom/rtciceconnectionstate"
	"github.com/matthewmueller/golly/dom/rtcicegatheringstate"
	"github.com/matthewmueller/golly/dom/rtcofferoptions"
	"github.com/matthewmueller/golly/dom/rtcpeerconnectioniceevent"
	"github.com/matthewmueller/golly/dom/rtcsessiondescription"
	"github.com/matthewmueller/golly/dom/rtcsignalingstate"
	"github.com/matthewmueller/golly/dom/rtcstatsreport"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCPeerConnection,omit"
type RTCPeerConnection interface {
	window.EventTarget

	// AddIceCandidate
	// js:"addIceCandidate,rewrite=addicecandidate"
	AddIceCandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// AddStream
	// js:"addStream,rewrite=addstream"
	AddStream(stream *window.MediaStream)

	// Close
	// js:"close,rewrite=cls"
	Close()

	// CreateAnswer
	// js:"createAnswer,rewrite=createanswer"
	CreateAnswer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription)

	// CreateOffer
	// js:"createOffer,rewrite=createoffer"
	CreateOffer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription)

	// GetConfiguration
	// js:"getConfiguration,rewrite=getconfiguration"
	GetConfiguration() (r *rtcconfiguration.RTCConfiguration)

	// GetLocalStreams
	// js:"getLocalStreams,rewrite=getlocalstreams"
	GetLocalStreams() (w []*window.MediaStream)

	// GetRemoteStreams
	// js:"getRemoteStreams,rewrite=getremotestreams"
	GetRemoteStreams() (w []*window.MediaStream)

	// GetStats
	// js:"getStats,rewrite=getstats"
	GetStats(selector *window.MediaStreamTrack, successCallback *func(report *rtcstatsreport.RTCStatsReport), failureCallback *func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport)

	// GetStreamByID
	// js:"getStreamById,rewrite=getstreambyid"
	GetStreamByID(streamId string) (w *window.MediaStream)

	// RemoveStream
	// js:"removeStream,rewrite=removestream"
	RemoveStream(stream *window.MediaStream)

	// SetLocalDescription
	// js:"setLocalDescription,rewrite=setlocaldescription"
	SetLocalDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// SetRemoteDescription
	// js:"setRemoteDescription,rewrite=setremotedescription"
	SetRemoteDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// CanTrickleIceCandidates prop
	// js:"canTrickleIceCandidates,rewrite=cantrickleicecandidates"
	CanTrickleIceCandidates() (canTrickleIceCandidates bool)

	// IceConnectionState prop
	// js:"iceConnectionState,rewrite=iceconnectionstate"
	IceConnectionState() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState)

	// IceGatheringState prop
	// js:"iceGatheringState,rewrite=icegatheringstate"
	IceGatheringState() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState)

	// LocalDescription prop
	// js:"localDescription,rewrite=localdescription"
	LocalDescription() (localDescription *rtcsessiondescription.RTCSessionDescription)

	// Onaddstream prop
	// js:"onaddstream,rewrite=onaddstream"
	Onaddstream() (onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onaddstream prop
	// js:"setonaddstream,rewrite=setonaddstream"
	SetOnaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onicecandidate prop
	// js:"onicecandidate,rewrite=onicecandidate"
	Onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Onicecandidate prop
	// js:"setonicecandidate,rewrite=setonicecandidate"
	SetOnicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Oniceconnectionstatechange prop
	// js:"oniceconnectionstatechange,rewrite=oniceconnectionstatechange"
	Oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event))

	// Oniceconnectionstatechange prop
	// js:"setoniceconnectionstatechange,rewrite=setoniceconnectionstatechange"
	SetOniceconnectionstatechange(oniceconnectionstatechange func(window.Event))

	// Onicegatheringstatechange prop
	// js:"onicegatheringstatechange,rewrite=onicegatheringstatechange"
	Onicegatheringstatechange() (onicegatheringstatechange func(window.Event))

	// Onicegatheringstatechange prop
	// js:"setonicegatheringstatechange,rewrite=setonicegatheringstatechange"
	SetOnicegatheringstatechange(onicegatheringstatechange func(window.Event))

	// Onnegotiationneeded prop
	// js:"onnegotiationneeded,rewrite=onnegotiationneeded"
	Onnegotiationneeded() (onnegotiationneeded func(window.Event))

	// Onnegotiationneeded prop
	// js:"setonnegotiationneeded,rewrite=setonnegotiationneeded"
	SetOnnegotiationneeded(onnegotiationneeded func(window.Event))

	// Onremovestream prop
	// js:"onremovestream,rewrite=onremovestream"
	Onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onremovestream prop
	// js:"setonremovestream,rewrite=setonremovestream"
	SetOnremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onsignalingstatechange prop
	// js:"onsignalingstatechange,rewrite=onsignalingstatechange"
	Onsignalingstatechange() (onsignalingstatechange func(window.Event))

	// Onsignalingstatechange prop
	// js:"setonsignalingstatechange,rewrite=setonsignalingstatechange"
	SetOnsignalingstatechange(onsignalingstatechange func(window.Event))

	// RemoteDescription prop
	// js:"remoteDescription,rewrite=remotedescription"
	RemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription)

	// SignalingState prop
	// js:"signalingState,rewrite=signalingstate"
	SignalingState() (signalingState *rtcsignalingstate.RTCSignalingState)
}

// addicecandidate fn
func addicecandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	js.Rewrite("await $<.addIceCandidate($1, $2, $3)", candidate, successCallback, failureCallback)
}

// addstream fn
func addstream(stream *window.MediaStream) {
	js.Rewrite("$<.addStream($1)", stream)
}

// cls fn
func cls() {
	js.Rewrite("$<.close()")
}

// createanswer fn
func createanswer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("await $<.createAnswer($1, $2)", successCallback, failureCallback)
	return r
}

// createoffer fn
func createoffer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("await $<.createOffer($1, $2, $3)", successCallback, failureCallback, options)
	return r
}

// getconfiguration fn
func getconfiguration() (r *rtcconfiguration.RTCConfiguration) {
	js.Rewrite("$<.getConfiguration()")
	return r
}

// getlocalstreams fn
func getlocalstreams() (w []*window.MediaStream) {
	js.Rewrite("$<.getLocalStreams()")
	return w
}

// getremotestreams fn
func getremotestreams() (w []*window.MediaStream) {
	js.Rewrite("$<.getRemoteStreams()")
	return w
}

// getstats fn
func getstats(selector *window.MediaStreamTrack, successCallback *func(report *rtcstatsreport.RTCStatsReport), failureCallback *func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport) {
	js.Rewrite("await $<.getStats($1, $2, $3)", selector, successCallback, failureCallback)
	return r
}

// getstreambyid fn
func getstreambyid(streamId string) (w *window.MediaStream) {
	js.Rewrite("$<.getStreamById($1)", streamId)
	return w
}

// removestream fn
func removestream(stream *window.MediaStream) {
	js.Rewrite("$<.removeStream($1)", stream)
}

// setlocaldescription fn
func setlocaldescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	js.Rewrite("await $<.setLocalDescription($1, $2, $3)", description, successCallback, failureCallback)
}

// setremotedescription fn
func setremotedescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	js.Rewrite("await $<.setRemoteDescription($1, $2, $3)", description, successCallback, failureCallback)
}

// cantrickleicecandidates prop
func cantrickleicecandidates() (canTrickleIceCandidates bool) {
	js.Rewrite("$<.canTrickleIceCandidates")
	return canTrickleIceCandidates
}

// iceconnectionstate prop
func iceconnectionstate() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState) {
	js.Rewrite("$<.iceConnectionState")
	return iceConnectionState
}

// icegatheringstate prop
func icegatheringstate() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState) {
	js.Rewrite("$<.iceGatheringState")
	return iceGatheringState
}

// localdescription prop
func localdescription() (localDescription *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("$<.localDescription")
	return localDescription
}

// onaddstream prop
func onaddstream() (onaddstream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$<.onaddstream")
	return onaddstream
}

// setonaddstream prop
func setonaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$<.onaddstream = onaddstream")
}

// onicecandidate prop
func onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent)) {
	js.Rewrite("$<.onicecandidate")
	return onicecandidate
}

// setonicecandidate prop
func setonicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent)) {
	js.Rewrite("$<.onicecandidate = onicecandidate")
}

// oniceconnectionstatechange prop
func oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event)) {
	js.Rewrite("$<.oniceconnectionstatechange")
	return oniceconnectionstatechange
}

// setoniceconnectionstatechange prop
func setoniceconnectionstatechange(oniceconnectionstatechange func(window.Event)) {
	js.Rewrite("$<.oniceconnectionstatechange = oniceconnectionstatechange")
}

// onicegatheringstatechange prop
func onicegatheringstatechange() (onicegatheringstatechange func(window.Event)) {
	js.Rewrite("$<.onicegatheringstatechange")
	return onicegatheringstatechange
}

// setonicegatheringstatechange prop
func setonicegatheringstatechange(onicegatheringstatechange func(window.Event)) {
	js.Rewrite("$<.onicegatheringstatechange = onicegatheringstatechange")
}

// onnegotiationneeded prop
func onnegotiationneeded() (onnegotiationneeded func(window.Event)) {
	js.Rewrite("$<.onnegotiationneeded")
	return onnegotiationneeded
}

// setonnegotiationneeded prop
func setonnegotiationneeded(onnegotiationneeded func(window.Event)) {
	js.Rewrite("$<.onnegotiationneeded = onnegotiationneeded")
}

// onremovestream prop
func onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$<.onremovestream")
	return onremovestream
}

// setonremovestream prop
func setonremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent)) {
	js.Rewrite("$<.onremovestream = onremovestream")
}

// onsignalingstatechange prop
func onsignalingstatechange() (onsignalingstatechange func(window.Event)) {
	js.Rewrite("$<.onsignalingstatechange")
	return onsignalingstatechange
}

// setonsignalingstatechange prop
func setonsignalingstatechange(onsignalingstatechange func(window.Event)) {
	js.Rewrite("$<.onsignalingstatechange = onsignalingstatechange")
}

// remotedescription prop
func remotedescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription) {
	js.Rewrite("$<.remoteDescription")
	return remoteDescription
}

// signalingstate prop
func signalingstate() (signalingState *rtcsignalingstate.RTCSignalingState) {
	js.Rewrite("$<.signalingState")
	return signalingState
}
