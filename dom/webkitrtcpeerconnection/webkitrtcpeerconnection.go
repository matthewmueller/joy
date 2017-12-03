package webkitrtcpeerconnection

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/mediastreamevent"
	"github.com/matthewmueller/joy/dom/rtcconfiguration"
	"github.com/matthewmueller/joy/dom/rtcicecandidate"
	"github.com/matthewmueller/joy/dom/rtciceconnectionstate"
	"github.com/matthewmueller/joy/dom/rtcicegatheringstate"
	"github.com/matthewmueller/joy/dom/rtcofferoptions"
	"github.com/matthewmueller/joy/dom/rtcpeerconnection"
	"github.com/matthewmueller/joy/dom/rtcpeerconnectioniceevent"
	"github.com/matthewmueller/joy/dom/rtcsessiondescription"
	"github.com/matthewmueller/joy/dom/rtcsignalingstate"
	"github.com/matthewmueller/joy/dom/rtcstatsreport"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ rtcpeerconnection.RTCPeerConnection = (*WebkitRTCPeerConnection)(nil)
var _ window.EventTarget = (*WebkitRTCPeerConnection)(nil)

// New fn
func New(configuration *rtcconfiguration.RTCConfiguration) *WebkitRTCPeerConnection {
	macro.Rewrite("webkitRTCPeerConnection")
	return &WebkitRTCPeerConnection{}
}

// WebkitRTCPeerConnection struct
// js:"WebkitRTCPeerConnection,omit"
type WebkitRTCPeerConnection struct {
}

// AddIceCandidate fn
// js:"addIceCandidate"
func (*WebkitRTCPeerConnection) AddIceCandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	macro.Rewrite("await $_.addIceCandidate($1, $2, $3)", candidate, successCallback, failureCallback)
}

// AddStream fn
// js:"addStream"
func (*WebkitRTCPeerConnection) AddStream(stream *window.MediaStream) {
	macro.Rewrite("$_.addStream($1)", stream)
}

// Close fn
// js:"close"
func (*WebkitRTCPeerConnection) Close() {
	macro.Rewrite("$_.close()")
}

// CreateAnswer fn
// js:"createAnswer"
func (*WebkitRTCPeerConnection) CreateAnswer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription) {
	macro.Rewrite("await $_.createAnswer($1, $2)", successCallback, failureCallback)
	return r
}

// CreateOffer fn
// js:"createOffer"
func (*WebkitRTCPeerConnection) CreateOffer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription) {
	macro.Rewrite("await $_.createOffer($1, $2, $3)", successCallback, failureCallback, options)
	return r
}

// GetConfiguration fn
// js:"getConfiguration"
func (*WebkitRTCPeerConnection) GetConfiguration() (r *rtcconfiguration.RTCConfiguration) {
	macro.Rewrite("$_.getConfiguration()")
	return r
}

// GetLocalStreams fn
// js:"getLocalStreams"
func (*WebkitRTCPeerConnection) GetLocalStreams() (w []*window.MediaStream) {
	macro.Rewrite("$_.getLocalStreams()")
	return w
}

// GetRemoteStreams fn
// js:"getRemoteStreams"
func (*WebkitRTCPeerConnection) GetRemoteStreams() (w []*window.MediaStream) {
	macro.Rewrite("$_.getRemoteStreams()")
	return w
}

// GetStats fn
// js:"getStats"
func (*WebkitRTCPeerConnection) GetStats(selector *window.MediaStreamTrack, successCallback *func(report *rtcstatsreport.RTCStatsReport), failureCallback *func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport) {
	macro.Rewrite("await $_.getStats($1, $2, $3)", selector, successCallback, failureCallback)
	return r
}

// GetStreamByID fn
// js:"getStreamById"
func (*WebkitRTCPeerConnection) GetStreamByID(streamId string) (w *window.MediaStream) {
	macro.Rewrite("$_.getStreamById($1)", streamId)
	return w
}

// RemoveStream fn
// js:"removeStream"
func (*WebkitRTCPeerConnection) RemoveStream(stream *window.MediaStream) {
	macro.Rewrite("$_.removeStream($1)", stream)
}

// SetLocalDescription fn
// js:"setLocalDescription"
func (*WebkitRTCPeerConnection) SetLocalDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	macro.Rewrite("await $_.setLocalDescription($1, $2, $3)", description, successCallback, failureCallback)
}

// SetRemoteDescription fn
// js:"setRemoteDescription"
func (*WebkitRTCPeerConnection) SetRemoteDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError)) {
	macro.Rewrite("await $_.setRemoteDescription($1, $2, $3)", description, successCallback, failureCallback)
}

// AddEventListener fn
// js:"addEventListener"
func (*WebkitRTCPeerConnection) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*WebkitRTCPeerConnection) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*WebkitRTCPeerConnection) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// CanTrickleIceCandidates prop
// js:"canTrickleIceCandidates"
func (*WebkitRTCPeerConnection) CanTrickleIceCandidates() (canTrickleIceCandidates bool) {
	macro.Rewrite("$_.canTrickleIceCandidates")
	return canTrickleIceCandidates
}

// IceConnectionState prop
// js:"iceConnectionState"
func (*WebkitRTCPeerConnection) IceConnectionState() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState) {
	macro.Rewrite("$_.iceConnectionState")
	return iceConnectionState
}

// IceGatheringState prop
// js:"iceGatheringState"
func (*WebkitRTCPeerConnection) IceGatheringState() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState) {
	macro.Rewrite("$_.iceGatheringState")
	return iceGatheringState
}

// LocalDescription prop
// js:"localDescription"
func (*WebkitRTCPeerConnection) LocalDescription() (localDescription *rtcsessiondescription.RTCSessionDescription) {
	macro.Rewrite("$_.localDescription")
	return localDescription
}

// Onaddstream prop
// js:"onaddstream"
func (*WebkitRTCPeerConnection) Onaddstream() (onaddstream func(*mediastreamevent.MediaStreamEvent)) {
	macro.Rewrite("$_.onaddstream")
	return onaddstream
}

// SetOnaddstream prop
// js:"onaddstream"
func (*WebkitRTCPeerConnection) SetOnaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent)) {
	macro.Rewrite("$_.onaddstream = $1", onaddstream)
}

// Onicecandidate prop
// js:"onicecandidate"
func (*WebkitRTCPeerConnection) Onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent)) {
	macro.Rewrite("$_.onicecandidate")
	return onicecandidate
}

// SetOnicecandidate prop
// js:"onicecandidate"
func (*WebkitRTCPeerConnection) SetOnicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent)) {
	macro.Rewrite("$_.onicecandidate = $1", onicecandidate)
}

// Oniceconnectionstatechange prop
// js:"oniceconnectionstatechange"
func (*WebkitRTCPeerConnection) Oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event)) {
	macro.Rewrite("$_.oniceconnectionstatechange")
	return oniceconnectionstatechange
}

// SetOniceconnectionstatechange prop
// js:"oniceconnectionstatechange"
func (*WebkitRTCPeerConnection) SetOniceconnectionstatechange(oniceconnectionstatechange func(window.Event)) {
	macro.Rewrite("$_.oniceconnectionstatechange = $1", oniceconnectionstatechange)
}

// Onicegatheringstatechange prop
// js:"onicegatheringstatechange"
func (*WebkitRTCPeerConnection) Onicegatheringstatechange() (onicegatheringstatechange func(window.Event)) {
	macro.Rewrite("$_.onicegatheringstatechange")
	return onicegatheringstatechange
}

// SetOnicegatheringstatechange prop
// js:"onicegatheringstatechange"
func (*WebkitRTCPeerConnection) SetOnicegatheringstatechange(onicegatheringstatechange func(window.Event)) {
	macro.Rewrite("$_.onicegatheringstatechange = $1", onicegatheringstatechange)
}

// Onnegotiationneeded prop
// js:"onnegotiationneeded"
func (*WebkitRTCPeerConnection) Onnegotiationneeded() (onnegotiationneeded func(window.Event)) {
	macro.Rewrite("$_.onnegotiationneeded")
	return onnegotiationneeded
}

// SetOnnegotiationneeded prop
// js:"onnegotiationneeded"
func (*WebkitRTCPeerConnection) SetOnnegotiationneeded(onnegotiationneeded func(window.Event)) {
	macro.Rewrite("$_.onnegotiationneeded = $1", onnegotiationneeded)
}

// Onremovestream prop
// js:"onremovestream"
func (*WebkitRTCPeerConnection) Onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent)) {
	macro.Rewrite("$_.onremovestream")
	return onremovestream
}

// SetOnremovestream prop
// js:"onremovestream"
func (*WebkitRTCPeerConnection) SetOnremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent)) {
	macro.Rewrite("$_.onremovestream = $1", onremovestream)
}

// Onsignalingstatechange prop
// js:"onsignalingstatechange"
func (*WebkitRTCPeerConnection) Onsignalingstatechange() (onsignalingstatechange func(window.Event)) {
	macro.Rewrite("$_.onsignalingstatechange")
	return onsignalingstatechange
}

// SetOnsignalingstatechange prop
// js:"onsignalingstatechange"
func (*WebkitRTCPeerConnection) SetOnsignalingstatechange(onsignalingstatechange func(window.Event)) {
	macro.Rewrite("$_.onsignalingstatechange = $1", onsignalingstatechange)
}

// RemoteDescription prop
// js:"remoteDescription"
func (*WebkitRTCPeerConnection) RemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription) {
	macro.Rewrite("$_.remoteDescription")
	return remoteDescription
}

// SignalingState prop
// js:"signalingState"
func (*WebkitRTCPeerConnection) SignalingState() (signalingState *rtcsignalingstate.RTCSignalingState) {
	macro.Rewrite("$_.signalingState")
	return signalingState
}
