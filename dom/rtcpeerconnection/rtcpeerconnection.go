package rtcpeerconnection

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/mediastreamevent"
	"github.com/matthewmueller/golly/dom2/rtcconfiguration"
	"github.com/matthewmueller/golly/dom2/rtcicecandidate"
	"github.com/matthewmueller/golly/dom2/rtciceconnectionstate"
	"github.com/matthewmueller/golly/dom2/rtcicegatheringstate"
	"github.com/matthewmueller/golly/dom2/rtcofferoptions"
	"github.com/matthewmueller/golly/dom2/rtcpeerconnectioniceevent"
	"github.com/matthewmueller/golly/dom2/rtcsessiondescription"
	"github.com/matthewmueller/golly/dom2/rtcsignalingstate"
	"github.com/matthewmueller/golly/dom2/rtcstatsreport"
	"github.com/matthewmueller/golly/dom2/window"
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
	SetOnaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onicecandidate prop
	// js:"onicecandidate"
	Onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Onicecandidate prop
	SetOnicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Oniceconnectionstatechange prop
	// js:"oniceconnectionstatechange"
	Oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event))

	// Oniceconnectionstatechange prop
	SetOniceconnectionstatechange(oniceconnectionstatechange func(window.Event))

	// Onicegatheringstatechange prop
	// js:"onicegatheringstatechange"
	Onicegatheringstatechange() (onicegatheringstatechange func(window.Event))

	// Onicegatheringstatechange prop
	SetOnicegatheringstatechange(onicegatheringstatechange func(window.Event))

	// Onnegotiationneeded prop
	// js:"onnegotiationneeded"
	Onnegotiationneeded() (onnegotiationneeded func(window.Event))

	// Onnegotiationneeded prop
	SetOnnegotiationneeded(onnegotiationneeded func(window.Event))

	// Onremovestream prop
	// js:"onremovestream"
	Onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onremovestream prop
	SetOnremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onsignalingstatechange prop
	// js:"onsignalingstatechange"
	Onsignalingstatechange() (onsignalingstatechange func(window.Event))

	// Onsignalingstatechange prop
	SetOnsignalingstatechange(onsignalingstatechange func(window.Event))

	// RemoteDescription prop
	// js:"remoteDescription"
	RemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription)

	// SignalingState prop
	// js:"signalingState"
	SignalingState() (signalingState *rtcsignalingstate.RTCSignalingState)
}
