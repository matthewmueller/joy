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
)

// js:"RTCPeerConnection,omit"
type RTCPeerConnection interface {
	window.EventTarget

	// AddIceCandidate
	// js:"addIceCandidate"
	AddIceCandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// AddStream
	// js:"addStream"
	AddStream(stream *window.MediaStream)

	// Close
	// js:"close"
	Close()

	// CreateAnswer
	// js:"createAnswer"
	CreateAnswer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription)

	// CreateOffer
	// js:"createOffer"
	CreateOffer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription)

	// GetConfiguration
	// js:"getConfiguration"
	GetConfiguration() (r *rtcconfiguration.RTCConfiguration)

	// GetLocalStreams
	// js:"getLocalStreams"
	GetLocalStreams() (w []*window.MediaStream)

	// GetRemoteStreams
	// js:"getRemoteStreams"
	GetRemoteStreams() (w []*window.MediaStream)

	// GetStats
	// js:"getStats"
	GetStats(selector *window.MediaStreamTrack, successCallback *func(report *rtcstatsreport.RTCStatsReport), failureCallback *func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport)

	// GetStreamByID
	// js:"getStreamById"
	GetStreamByID(streamId string) (w *window.MediaStream)

	// RemoveStream
	// js:"removeStream"
	RemoveStream(stream *window.MediaStream)

	// SetLocalDescription
	// js:"setLocalDescription"
	SetLocalDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// SetRemoteDescription
	// js:"setRemoteDescription"
	SetRemoteDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// CanTrickleIceCandidates prop
	// js:"canTrickleIceCandidates"
	CanTrickleIceCandidates() (canTrickleIceCandidates bool)

	// IceConnectionState prop
	// js:"iceConnectionState"
	IceConnectionState() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState)

	// IceGatheringState prop
	// js:"iceGatheringState"
	IceGatheringState() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState)

	// LocalDescription prop
	// js:"localDescription"
	LocalDescription() (localDescription *rtcsessiondescription.RTCSessionDescription)

	// Onaddstream prop
	// js:"onaddstream"
	Onaddstream() (onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onaddstream prop
	// js:"setonaddstream"
	SetOnaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onicecandidate prop
	// js:"onicecandidate"
	Onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Onicecandidate prop
	// js:"setonicecandidate"
	SetOnicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Oniceconnectionstatechange prop
	// js:"oniceconnectionstatechange"
	Oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event))

	// Oniceconnectionstatechange prop
	// js:"setoniceconnectionstatechange"
	SetOniceconnectionstatechange(oniceconnectionstatechange func(window.Event))

	// Onicegatheringstatechange prop
	// js:"onicegatheringstatechange"
	Onicegatheringstatechange() (onicegatheringstatechange func(window.Event))

	// Onicegatheringstatechange prop
	// js:"setonicegatheringstatechange"
	SetOnicegatheringstatechange(onicegatheringstatechange func(window.Event))

	// Onnegotiationneeded prop
	// js:"onnegotiationneeded"
	Onnegotiationneeded() (onnegotiationneeded func(window.Event))

	// Onnegotiationneeded prop
	// js:"setonnegotiationneeded"
	SetOnnegotiationneeded(onnegotiationneeded func(window.Event))

	// Onremovestream prop
	// js:"onremovestream"
	Onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onremovestream prop
	// js:"setonremovestream"
	SetOnremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onsignalingstatechange prop
	// js:"onsignalingstatechange"
	Onsignalingstatechange() (onsignalingstatechange func(window.Event))

	// Onsignalingstatechange prop
	// js:"setonsignalingstatechange"
	SetOnsignalingstatechange(onsignalingstatechange func(window.Event))

	// RemoteDescription prop
	// js:"remoteDescription"
	RemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription)

	// SignalingState prop
	// js:"signalingState"
	SignalingState() (signalingState *rtcsignalingstate.RTCSignalingState)
}
